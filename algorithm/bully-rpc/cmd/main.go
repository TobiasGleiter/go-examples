package main

import (
	"fmt"
	"net/rpc"
	"sync"
)


// nodeAddressByID: It includes nodes currently in cluster
var nodeAddressByID = map[string]string{
	"node-01": "localhost:6001",
	"node-02": "localhost:6002",
	"node-03": "localhost:6003",
	"node-04": "localhost:6004",
   }

type Node struct {
	ID       string
	Addr     string
	Peers    *Peers
	eventBus Bus
}

type Peer struct {
	ID        string
	RPCClient *rpc.Client
}

type Peers struct {
	*sync.RWMutex
	peerByID map[string]*Peer
}


func main() {
	fmt.Println("Hello!")
}

func (node *Node) ConnectToPeers() {
	for peerID, peerAddr := range nodeAddressByID {
		if node.IsItself(peerID) {
		continue
		}

		rpcClient := node.connect(peerAddr)
		pingMessage := Message{FromPeerID: node.ID, Type: PING}
		reply, _ := node.CommunicateWithPeer(rpcClient, pingMessage)

		if reply.IsPongMessage() {
		node.Peers.Add(peerID, rpcClient)
		}
	}
}

func (node *Node) Elect() {
	isHighestRankedNodeAvailable := false

	peers := node.Peers.ToList()
	for i := range peers {
		peer := peers[i]

		if node.IsRankHigherThan(peer.ID) {
		continue
		}

		electionMessage := Message{FromPeerID: node.ID, Type: ELECTION}
		reply, _ := node.CommunicateWithPeer(peer.RPCClient, electionMessage)

		if reply.IsAliveMessage() {
		isHighestRankedNodeAvailable = true
		}
	}

	if !isHighestRankedNodeAvailable {
		leaderID := node.ID
		electedMessage := Message{FromPeerID: leaderID, Type: ELECTED}
		node.BroadcastMessage(electedMessage)
	}
}

func (node *Node) PingLeaderContinuously(_ string, payload any) {
	leaderID := payload.(string)
   
   ping:
	pingMessage := Message{FromPeerID: node.ID, Type: PING}
	reply, err := node.CommunicateWithPeer(leader.RPCClient, pingMessage)
	if err != nil {
	 log.Info().Msgf("Leader is down, new election about to start!")
	 node.Peers.Delete(leaderID)
	 node.Elect()
	 return
	}
   
	if reply.IsPongMessage() {
	 time.Sleep(3 * time.Second)
	 goto ping
	}
   }

type HandlerFunc func(eventName string, payload any)

type Bus interface {
	Emit(eventName string, payload any)
	Subscribe(eventName string, handlers ...HandlerFunc)
}

type eventBus struct {
	h map[string][]HandlerFunc
}

func NewBus() Bus {
	return &eventBus{
		h: make(map[string][]HandlerFunc),
	}
}

func (b *eventBus) Subscribe(eventName string, handlers ...HandlerFunc) {
	b.h[eventName] = append(b.h[eventName], handlers...)
}

func (b *eventBus) Emit(eventName string, payload any) {
	for _, handler := range b.h[eventName] {
		go handler(eventName, payload)
	}
}

const (
	LeaderElected = "leaderElected"
)