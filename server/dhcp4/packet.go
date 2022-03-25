package dhcp4

import (
	"net"
)

const (
	bflagUni      = 0
	bflagBrd      = 1
	opCodeRequest = 1
	opCodeReply   = 2
)

type DhcpPacket struct {
	Op      int
	HType   int
	HLen    int
	Hops    int
	Xid     string
	Sec     int
	Flags   int
	YiAddr  net.IP
	SiAddr  net.IP
	GiAddr  net.IP
	ChAddr  string
	SName   string
	File    string
	Options []interface{}
}

func NewDhcpPacket() *DhcpPacket {
	// yaddr := net.ParseIP("8.8.8.8").To4()
	return new(DhcpPacket)
}
