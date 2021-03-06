package main

import (
	"encoding/hex"
	"fmt"
	"github.com/EnsicoinDevs/eccd/blockchain"
	"github.com/EnsicoinDevs/eccd/mempool"
	"github.com/EnsicoinDevs/eccd/network"
	"github.com/EnsicoinDevs/eccd/peer"
	"github.com/EnsicoinDevs/eccd/sssync"
	"github.com/EnsicoinDevs/eccd/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"sync"
)

// Server provides an ensicoin server for handling communications to and from ensicoin peers.
type Server struct {
	blockchain   *blockchain.Blockchain
	mempool      *mempool.Mempool
	synchronizer *sssync.Synchronizer
	rpcServer    *rpcServer

	listener net.Listener

	mutex *sync.RWMutex
	peers map[string]*peer.Peer

	quit chan struct{}
}

// NewServer returns a new Server instance.
func NewServer() *Server {
	server := &Server{
		mutex: &sync.RWMutex{},
		peers: make(map[string]*peer.Peer),

		quit: make(chan struct{}),
	}

	server.blockchain = blockchain.NewBlockchain(&blockchain.Config{
		DataDir:       viper.GetString("datadir"),
		OnPushedBlock: server.onPushedBlock,
		OnPoppedBlock: server.onPoppedBlock,
	})

	server.rpcServer = newRPCServer(server)

	server.mempool = mempool.NewMempool(&mempool.Config{
		FetchUtxos:   server.blockchain.FetchUtxos,
		OnAcceptedTx: server.rpcServer.HandleAcceptedTx,
	})

	server.synchronizer = sssync.NewSynchronizer(&sssync.Config{
		Broadcast: server.Broadcast,
	}, server.blockchain, server.mempool)

	return server
}

// Start starts the server.
func (server *Server) Start() {
	go server.blockchain.Load()
	go server.mempool.Start()

	go server.synchronizer.Start()
	go server.rpcServer.Start()

	var err error
	server.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil {
		log.WithError(err).Panic("error launching the tcp server")
	}

	log.Infof("server listening on :%d", viper.GetInt("port"))

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			select {
			case <-server.quit:
				return
			default:
				log.WithError(err).Error("error accepting a new connection")
			}

			return
		}

		peer := peer.NewPeer(conn, &peer.Config{
			Callbacks: peer.PeerCallbacks{
				OnReady:        server.onReady,
				OnDisconnected: server.onDisconnected,
				OnGetData:      server.onGetData,
				OnGetBlocks:    server.onGetBlocks,
				OnInv:          server.onInv,
				OnBlock:        server.onBlock,
			},
		}, true)

		log.WithField("peer", peer).Info("new incoming connection")

		go peer.Start()
	}
}

// Stop stops the server.
func (server *Server) Stop() error {
	log.Debug("server shutting down")
	defer log.Debug("server shutdown complete")

	err := server.rpcServer.Stop()
	if err != nil {
		return err
	}

	close(server.quit)

	server.listener.Close()

	err = server.synchronizer.Stop()
	if err != nil {
		return err
	}

	err = server.mempool.Stop()
	if err != nil {
		return err
	}

	return server.blockchain.Stop()
}

func (server *Server) onPushedBlock(block *blockchain.Block) error {
	server.synchronizer.OnPushedBlock(block)
	server.rpcServer.OnPushedBlock(block)

	return nil
}

func (server *Server) onPoppedBlock(block *blockchain.Block) error {
	go server.synchronizer.OnPoppedBlock(block)
	go server.rpcServer.OnPoppedBlock(block)

	return nil
}

// ConnectTo creates a new peer and connect the server to it.
func (server *Server) ConnectTo(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	go peer.NewPeer(conn, &peer.Config{
		Callbacks: peer.PeerCallbacks{
			OnReady:        server.onReady,
			OnDisconnected: server.onDisconnected,
			OnGetData:      server.onGetData,
			OnGetBlocks:    server.onGetBlocks,
			OnInv:          server.onInv,
			OnBlock:        server.onBlock,
			OnTx:           server.onTx,
		},
	}, false).Start()

	return nil
}

// DisconnectFrom disconnects the server from a peer.
func (server *Server) DisconnectFrom(peer *peer.Peer) error {
	return peer.Stop()
}

// FindPeerByAddress searches for a peer using its external address.
func (server *Server) FindPeerByAddress(address string) (*peer.Peer, error) {
	log.WithFields(log.Fields{
		"func":  "FindPeerByAddress",
		"mutex": "server",
	}).Trace("rlocking")
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	defer log.WithFields(log.Fields{
		"func":  "FindPeerByAddress",
		"mutex": "server",
	}).Trace("runlocking")

	return server.peers[address], nil
}

// Broadcast broadcasts a message to all connected peers.
func (server *Server) Broadcast(message network.Message) error {
	for _, peer := range server.peers {
		peer.Send(message)
	}

	return nil
}

func (server *Server) addPeer(peer *peer.Peer) {
	log.WithFields(log.Fields{
		"func":  "addPeer",
		"mutex": "server",
	}).Trace("locking")
	server.mutex.Lock()

	server.peers[peer.RemoteAddr().String()] = peer

	log.WithFields(log.Fields{
		"func":  "addPeer",
		"mutex": "server",
	}).Trace("unlocking")
	server.mutex.Unlock()
}

func (server *Server) removePeer(peer *peer.Peer) {
	log.WithFields(log.Fields{
		"func":  "removePeer",
		"mutex": "server",
	}).Trace("locking")
	server.mutex.Lock()
	server.peers[peer.RemoteAddr().String()] = peer

	log.WithFields(log.Fields{
		"func":  "removePeer",
		"mutex": "server",
	}).Trace("unlocking")
	server.mutex.Unlock()
}

func (server *Server) onReady(peer *peer.Peer) {
	log.WithField("peer", peer).Debug("connection established")
	server.addPeer(peer)
	server.synchronizer.HandleReadyPeer(peer)
}

func (server *Server) onDisconnected(peer *peer.Peer) {
	server.removePeer(peer)

	server.synchronizer.HandleDisconnectedPeer(peer)

	log.WithField("peer", peer).Info("disconnected")
}

func (server *Server) onInv(peer *peer.Peer, message *network.InvMessage) {
	for _, inv := range message.Inventory {
		if inv.InvType == network.INV_VECT_BLOCK {
			server.synchronizer.HandleBlockInvVect(peer, inv)
		} else if inv.InvType == network.INV_VECT_TX {
			if server.mempool.FindTxByHash(inv.Hash) == nil {
				peer.Send(&network.GetDataMessage{
					Inventory: []*network.InvVect{inv},
				})
			}
		} else {
			// TODO: handle this error
		}
	}
}

func (server *Server) onBlock(peer *peer.Peer, message *network.BlockMessage) {
	server.synchronizer.HandleBlock(peer, message)
}

// ProcessBlock validates a block.
func (server *Server) ProcessBlock(message *network.BlockMessage) {
	server.synchronizer.ProcessBlock(message)
}

// ProcessTx validates a tx.
func (server *Server) ProcessTx(message *network.TxMessage) {
	server.mempool.ProcessTx(blockchain.NewTxFromTxMessage(message))
}

func (server *Server) onTx(peer *peer.Peer, message *network.TxMessage) {
	log.WithFields(log.Fields{
		"hash": hex.EncodeToString(message.Hash()[:]),
	}).Info("new tx received")

	server.ProcessTx(message)
}

func (server *Server) onGetData(peer *peer.Peer, message *network.GetDataMessage) {
	for _, invVect := range message.Inventory {
		switch invVect.InvType {
		case network.INV_VECT_BLOCK:
			block, err := server.blockchain.FindBlockByHash(invVect.Hash)
			if err != nil {
				log.WithFields(log.Fields{
					"peer": peer,
				}).WithError(err).Error("error finding a block by hash")
			}

			if block == nil {
				continue
			}

			peer.Send(block.Msg)
		case network.INV_VECT_TX:

		default:
			log.Error("unknown inv vect type")
		}
	}
}

func (server *Server) onGetBlocks(peer *peer.Peer, message *network.GetBlocksMessage) {
	var startAt *utils.Hash

	for _, hash := range message.BlockLocator {
		haveBlock, err := server.blockchain.HaveBlock(hash)
		if err != nil {
			log.WithFields(log.Fields{
				"peer": peer,
			}).WithError(err).Error("error checking if a block exists")
			continue
		}

		if !haveBlock {
			continue
		}

		startAt = hash
		break
	}

	if startAt == nil {
		return
	}

	var inventory []*network.InvVect

	hashes, err := server.blockchain.FindBlockHashesStartingAt(startAt, 500)
	if err != nil {
		log.WithField("startAt", startAt).WithError(err).Error("error finding the block hashes")
		return
	}

	for _, hash := range hashes {
		inventory = append(inventory, &network.InvVect{
			InvType: network.INV_VECT_BLOCK,
			Hash:    hash,
		})

	}

	peer.Send(&network.InvMessage{
		Inventory: inventory,
	})
}
