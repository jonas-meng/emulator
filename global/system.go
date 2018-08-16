package global

import (
	"github.com/jonas-meng/emulator/node"
	"math/rand"
	"github.com/sirupsen/logrus"
	"github.com/jonas-meng/emulator/event"
)

type System struct {
	NodeSize    uint64
	TotalWeight uint64
	Nodes       []*node.Node
	Oracle      *event.Oracle
}

func (s System) LeaderElection() uint64 {
	threshold := rand.Float64()
	c := 0.0
	leader := 0
	for i, n := range s.Nodes {
		c = c + n.Weight
		if c >= threshold {
			leader = i
			break
		}
	}
	return uint64(leader)
}

func Init(nodeSize uint64) *System {
	sys := System{nodeSize, 1, node.GenerateNodes(nodeSize), event.NewOracle()}
	logrus.Printf("Node size: %v", sys.NodeSize)
	logrus.Printf("Total Weight : %v", sys.TotalWeight)
	for _, n := range sys.Nodes {
		n.LogInfo()
	}
	return &sys
}

func (s System) EventApplication() {
	e := s.Oracle.FetchEvent()
	if e != nil {
		logrus.Printf(e.ToString())
		s.Nodes[e.NodeID].ApplyEvent(e.State)
	}
}

func (s System) Run() {
	// initiate all nodes
	for _, n := range s.Nodes {
		go n.Run()
	}

	go s.Oracle.EventGenerator(len(s.Nodes), 7)

	for round := 0; ; round++ {
		// elected leader
		leaderID := s.LeaderElection()
		logrus.Printf("Round %v, Elected Leader Node ID: %v", round, leaderID)
	}
}
