package main

import (
	"fmt"
	"github.com/EnsicoinDevs/ensicoincoin/blockchain"
	"github.com/EnsicoinDevs/ensicoincoin/consensus"
	"github.com/EnsicoinDevs/ensicoincoin/mempool"
	"github.com/EnsicoinDevs/ensicoincoin/network"
	"github.com/EnsicoinDevs/ensicoincoin/peer"
	log "github.com/sirupsen/logrus"
	"net"
)

type invWithPeer struct {
	hashes []string
	peer   *ServerPeer
}

type ServerPeer struct {
	*peer.Peer

	server *Server
}

func (server *Server) newServerPeer() *ServerPeer {
	return &ServerPeer{
		server: server,
	}
}

func (sp *ServerPeer) newConfig() *peer.Config {
	return &peer.Config{
		Callbacks: peer.PeerCallbacks{
			OnReady:     sp.onReady,
			OnWhoami:    sp.onWhoami,
			OnInv:       sp.onInv,
			OnGetBlocks: sp.onGetBlocks,
			OnTx:        sp.onTx,
		},
	}
}

func (server *Server) newIngoingServerPeer(conn net.Conn) *ServerPeer {
	sp := server.newServerPeer()

	sp.Peer = peer.NewIngoingPeer(sp.newConfig())

	sp.AttachConn(conn)

	return sp
}

func (server *Server) newOutgoingServerPeer(conn net.Conn) *ServerPeer {
	sp := server.newServerPeer()

	sp.Peer = peer.NewOutgoingPeer(sp.newConfig())

	sp.AttachConn(conn)

	return sp
}

type Server struct {
	blockchain *blockchain.Blockchain
	mempool    *mempool.Mempool

	blocksInvs       chan *invWithPeer
	transactionsInvs chan *invWithPeer

	peers    []*ServerPeer
	listener net.Listener

	synced bool
}

func NewServer(blockchain *blockchain.Blockchain, mempool *mempool.Mempool) *Server {
	server := &Server{
		blockchain: blockchain,
		mempool:    mempool,

		blocksInvs:       make(chan *invWithPeer),
		transactionsInvs: make(chan *invWithPeer),
	}

	return server
}

func (server *Server) registerPeer(peer *ServerPeer) {
	server.peers = append(server.peers, peer)
}

func (server *Server) syncWith(peer *ServerPeer) {
	longestChain, err := server.blockchain.FindLongestChain()
	if err != nil {
		log.WithError(err).Error("error finding the longest chain")
		return
	}

	peer.Send("getblocks", &network.GetBlocksMessage{
		Hashes: []string{longestChain.Hash},
	})

	server.synced = true
}

func (server *Server) Start() {
	var err error
	server.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", consensus.INGOING_PORT))
	if err != nil {
		log.WithError(err).Panic("error launching the tcp server")
	}

	go server.handleBlocksInvs()
	go server.handleTransactionsInvs()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.WithError(err).Error("error accepting a new connection")
		} else {
			log.Info("we have a new ingoing peer")
			server.registerPeer(server.newIngoingServerPeer(conn))
		}
	}
}

func (server *Server) handleBlocksInvs() {
	for inv := range server.blocksInvs {
		invToSend := network.InvMessage{
			Type: "b",
		}

		for _, hash := range inv.hashes {
			block, err := server.blockchain.FindBlockByHash(hash)
			if err != nil {
				log.WithError(err).WithField("hash", hash).Error("error finding a block")
			}

			if block == nil {
				invToSend.Hashes = append(invToSend.Hashes, hash)
			}
		}

		if len(invToSend.Hashes) > 0 {
			inv.peer.Send("getdata", &network.GetDataMessage{
				Inv: invToSend,
			})
		}
	}
}

func (server *Server) handleTransactionsInvs() {
	for inv := range server.transactionsInvs {
		invToSend := network.InvMessage{
			Type: "t",
		}

		for _, hash := range inv.hashes {
			tx := server.mempool.FindTxByHash(hash)

			if tx == nil {
				invToSend.Hashes = append(invToSend.Hashes, hash)
			}
		}

		if len(invToSend.Hashes) > 0 {
			inv.peer.Send("getdata", &network.GetDataMessage{
				Inv: invToSend,
			})
		}
	}
}

func (server *Server) Stop() {
	server.listener.Close()
}

func (server *Server) RegisterOutgoingPeer(conn net.Conn) {
	server.registerPeer(server.newOutgoingServerPeer(conn))
}

func (sp *ServerPeer) onReady() {
	if !sp.Ingoing() {
		sp.Send("whoami", &network.WhoamiMessage{
			Version: 0,
		})
	}
}

func (sp *ServerPeer) onWhoami(message *network.WhoamiMessage) {
	log.WithField("peer", sp).Info("whoami received")

	if sp.Ingoing() {
		sp.Send("whoami", &network.WhoamiMessage{
			Version: 0,
		})

		// sp.Connected = true
	} else {
		// sp.Connected = true
	}

	log.WithField("peer", sp).Info("connection established")

	if !sp.server.synced {
		sp.server.syncWith(sp)
	}
}

func (peer *ServerPeer) onInv(message *network.InvMessage) {
	if message.Type == "b" {
		peer.server.blocksInvs <- &invWithPeer{
			hashes: message.Hashes,
			peer:   peer,
		}
	} else if message.Type == "t" {
		peer.server.transactionsInvs <- &invWithPeer{
			hashes: message.Hashes,
			peer:   peer,
		}
	} else {
		log.WithFields(log.Fields{
			"peer":    peer,
			"message": message,
		}).Error("unknown inv type")
	}
}

func (peer *ServerPeer) onBlock(message *network.BlockMessage) {

}

func (server *Server) broadcastTx(tx *blockchain.Tx, sourcePeer *ServerPeer) {
	for _, peer := range server.peers {
		if peer != sourcePeer {
			peer.Send("inv", &network.InvMessage{
				Type:   "t",
				Hashes: []string{tx.Hash},
			})
		}
	}
}

func (peer *ServerPeer) onTx(message *network.TxMessage) {
	tx := blockchain.NewTxFromTxMessage(message)
	tx.ComputeHash()
	valid := peer.server.mempool.ProcessTx(tx)

	if valid {
		peer.server.broadcastTx(tx, peer)
	}
}

func (peer *ServerPeer) onGetBlocks(message *network.GetBlocksMessage) {
	var startAt string

	for _, hash := range message.Hashes {
		block, err := peer.server.blockchain.FindBlockByHash(hash)
		if err != nil {
			log.WithFields(log.Fields{
				"peer": peer,
			}).WithError(err).Error("error finding a block by hash")
			continue
		}

		if block == nil {
			continue
		}

		startAt = block.Hash
		break
	}

	if startAt == "" {
		startAt = peer.server.blockchain.GenesisBlock.Hash
	}

	invToSend := &network.InvMessage{
		Type: "b",
	}

	hashes, err := peer.server.blockchain.FindBlockHashesStartingAt(startAt)
	if err != nil {
		log.WithField("startAt", startAt).WithError(err).Error("error finding the block hashes")
		return
	}

	for _, hash := range hashes {
		invToSend.Hashes = append(invToSend.Hashes, hash)
	}

	peer.Send("inv", invToSend)
}
