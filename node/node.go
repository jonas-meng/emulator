package node

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
)

// node state
const (
	_ = iota
	// network status
	CONNECTED = 1 << iota
	DELAYED
	DISCONNECTED

	// node aliveness
	ALIVE
	DEAD

	// node status
	NORMAL
	CHAOTIC
	MALICIOUS
)

type Node struct {
	Key    *ecdsa.PrivateKey
	Weight float64
	State  int
}

func (n Node) Run() {}

func (n *Node) ApplyEvent(State int) {
	switch State {
	case CONNECTED, DELAYED, DISCONNECTED:
		n.State = (n.State & ^(CONNECTED | DELAYED | DISCONNECTED)) | State
	case ALIVE, DEAD:
		n.State = (n.State & ^(ALIVE | DEAD)) | State
	case NORMAL, CHAOTIC, MALICIOUS:
		n.State = (n.State & ^(NORMAL | CHAOTIC | MALICIOUS)) | State
	}
}

func (n Node) LogInfo() {
	logrus.Printf("Private Key: %v", hexutil.EncodeBig(n.Key.D))
	logrus.Printf("Public Key: %v", hexutil.Encode(crypto.FromECDSAPub(&n.Key.PublicKey)))
	logrus.Printf("Address: %v", crypto.PubkeyToAddress(n.Key.PublicKey).Hex())
	logrus.Printf("Weight: %v", n.Weight)
	n.LogState()
}

func (n Node) LogState() {
	logrus.Printf("is connected: %v", n.IsConnected())
	logrus.Printf("is delayed: %v", n.IsDelayed())
	logrus.Printf("is disconnected: %v", n.IsDisconnected())
	logrus.Printf("is alive: %v", n.IsAlive())
	logrus.Printf("is dead: %v", n.IsDead())
	logrus.Printf("is normal: %v", n.IsNormal())
	logrus.Printf("is chaotic: %v", n.IsNormal())
	logrus.Printf("is malicious: %v", n.IsMalicious())
}

func (n Node) IsConnected() bool {
	return n.State&CONNECTED != 0
}

func (n Node) IsDisconnected() bool {
	return n.State&DISCONNECTED != 0
}

func (n Node) IsDelayed() bool {
	return n.State&DELAYED != 0
}

func (n Node) IsAlive() bool {
	return n.State&ALIVE != 0
}

func (n Node) IsDead() bool {
	return n.State&DEAD != 0
}

func (n Node) IsNormal() bool {
	return n.State&NORMAL != 0
}

func (n Node) IsChaotic() bool {
	return n.State&CHAOTIC != 0
}

func (n Node) IsMalicious() bool {
	return n.State&MALICIOUS != 0
}

func NewNode() *Node {
	key, _ := crypto.GenerateKey()
	vote := rand.Float64()
	return &Node{key, vote, CONNECTED | ALIVE | NORMAL}
}

func GenerateNodes(size uint64) []*Node {
	count := 0.0
	nodes := make([]*Node, size)
	for i := uint64(0); i < size; i++ {
		nodes[i] = NewNode()
		count += nodes[i].Weight
	}
	for _, n := range nodes {
		n.Weight = n.Weight / count
	}
	return nodes
}
