package main

import (
	"bufio"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/EnsicoinDevs/ensicoincoin/blockchain"
	"github.com/EnsicoinDevs/ensicoincoin/consensus"
	"github.com/EnsicoinDevs/ensicoincoin/mempool"
	"github.com/EnsicoinDevs/ensicoincoin/miner"
)

var (
	peerPort        int
	discordToken    string
	interactiveMode bool
	mining          bool
	epprof          bool
)

func init() {
	flag.IntVar(&peerPort, "port", consensus.INGOING_PORT, "The port of the node.")
	flag.StringVar(&discordToken, "token", "", "A discord token.")
	flag.BoolVar(&interactiveMode, "i", false, "Interactive mode.")
	flag.BoolVar(&mining, "mining", false, "Miner mode.")
	flag.BoolVar(&epprof, "pprof", false, "Enable pprof")
}

func main() {
	log.SetLevel(log.DebugLevel)

	flag.Parse()

	if epprof {
		log.Debug("?")

		go func() {
			log.Debug("pprof")
			log.Debug(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	log.Info("ENSICOINCOIN is starting")

	blockchain := blockchain.NewBlockchain()
	blockchain.Load()

	mempool := mempool.NewMempool(&mempool.Config{
		FetchUtxos: blockchain.FetchUtxos,
	})

	bestBlock, err := blockchain.FindLongestChain()
	if err != nil {
		log.WithError(err).Error("error finding the best block")
	}

	miner := &miner.Miner{
		Config:     &miner.Config{},
		BestBlock:  bestBlock,
		Blockchain: blockchain,
		Mempool:    mempool,
	}

	server := NewServer(blockchain, mempool, miner)

	miner.Config.ProcessBlock = server.ProcessMinerBlock

	go server.Start()

	if discordToken != "" {
		startDiscordBootstraping(server)
	}

	if mining {
		miner.Start()
	}

	rpcServer := newRpcServer(blockchain, server)

	go rpcServer.Start()

	log.Info("ENSICOINCOIN is now running")

	if !interactiveMode {
		ch := make(chan bool)
		<-ch
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		rawCommand, _ := reader.ReadString('\n')

		command := strings.Split(rawCommand[:len(rawCommand)-1], " ")

		if command[0] == "q" || command[0] == "quit" {
			break
		}

		switch command[0] {
		case "help":
			log.Info("quit, help")
		case "connect":
			address := command[1]

			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.WithError(err).WithField("address", address).Error("error dialing this address")
				break
			}

			server.RegisterOutgoingPeer(conn)
		}
	}

	miner.Stop()
	server.Stop()

	log.Info("Good bye.")
}
