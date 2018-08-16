package event

import (
	"github.com/jonas-meng/emulator/node"
	"strconv"
)

type Event struct {
	NodeID int
	State  int
}

func (e Event) ToString() string {
	res := "Node ID: " + strconv.FormatInt(int64(e.NodeID), 10)
	switch e.State {
	case node.CONNECTED:
		res = res + ", Network is Connected"
	case node.DELAYED:
		res = res + ", Network is Delayed"
	case node.DISCONNECTED:
		res = res + ", Network is Disconnected"
	case node.ALIVE:
		res = res + ", Node is Alive"
	case node.DEAD:
		res = res + ", Node is Dead"
	case node.NORMAL:
		res = res + ", Node is Normal"
	case node.CHAOTIC:
		res = res + ", Node is Chaotic"
	case node.MALICIOUS:
		res = res + ", Node is Malicious"
	}
	return res
}
