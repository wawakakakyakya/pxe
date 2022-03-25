package dhcp4

import (
	"fmt"
	"net"
	"pxe/config"
	"pxe/logger"
)

type AddrRecord struct {
}

type AddrRecords []AddrRecord

func dhcp() {
	fmt.Println("dhcp.dhcp")
}

type DhcpSevrer struct {
	config    *config.Config
	handler   *DhcpHandler
	listener  *net.UDPConn
	logger    *logger.Logger
	closeChan *chan bool
}

func (d *DhcpSevrer) Run() {
	d.logger.Info("[dhcp] dhcp server running")
	listenAddr := &net.UDPAddr{
		IP:   net.ParseIP(d.config.Dhcp.Listen),
		Port: d.config.Dhcp.Port,
	}
	listener, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		d.logger.Error("[dhcp] Run Server failed")
		d.logger.Error(err.Error())
	} else {
		d.logger.Info(fmt.Sprintf("DHCP running on %s:%d", d.config.Dhcp.Listen, d.config.Dhcp.Port))
	}

	// defer d.Close()

	d.listener = listener
	packetChan := make(chan int)
	go d.Read(packetChan)

	for {
		select {
		case <-*d.closeChan:
			d.logger.Info("close channed received, end read process")
			return
		case <-packetChan:
			// d.logger.Debug(fmt.Sprintf("Received from %v, Data : %s\n", addr, string(buf[:n])))
			d.logger.Info("pacetChan block!")
		}
	}
}

func (d *DhcpSevrer) Read(packetChan chan int) {
	var buf [1500]byte
	for {
		d.logger.Info("start readfromudp, waiting...")
		n, addr, err := d.listener.ReadFromUDP(buf[0:])
		if err != nil {
			d.logger.Info("Read packet fail")
			d.logger.Info(err.Error())
			continue
		}
		packetChan <- n
		d.logger.Info(fmt.Sprintf("Received from %v, Data : %s\n", addr, string(buf[:n])))
		// d.logger.Debug("end readfromudp")
		// https://kobatako.hatenablog.com/entry/2017/11/24/210943
		// conn.WirteToUDP([]byte(daytime), addr)
	}
}

//called from server.Close
func (d *DhcpSevrer) Close() {
	if err := d.listener.Close(); err != nil {
		d.logger.Info("DHCP Server closed with error")
		d.logger.Info(err.Error())
	} else {
		d.logger.Info("DHCP Server closed ")
	}

}

func NewDhcpServer(c *config.Config, closeChan *chan bool) *DhcpSevrer {
	d := new(DhcpSevrer)
	d.config = c
	d.logger = logger.NewLogger("dhcp")
	d.closeChan = closeChan

	return d
}

// func RunServer(handler Handler) error {
// 	l, err := net.ListenPacket("udp4", ":67")
// 	if err != nil {
// 		return err
// 	}
// 	defer l.Close()
// 	return Serve(l, handler)
// }

// func (d *DhcpService) Run(cancel context.CancelFunc) {
// 	listen := "0.0.0.0:67"
// 	conn, err := dhcp4.NewConn(listen)
// 	if err != nil {
// 		log.Fatalf("[FATAL] Unable to listen on %s: %v", listen, err)
// 	}
// 	defer conn.Close()

// 	log.Printf("[INFO] Starting DHCP server...")
// 	for {
// 		req, intf, err := conn.RecvDHCP() //    (1)
// 		if err != nil {
// 			log.Fatalf("[ERROR] Failed to receive DHCP package: %v", err)
// 		}

// 		log.Printf("[INFO] Received %s from %s", req.Type, req.HardwareAddr)
// 		resp := &dhcp4.Packet{ // (2)
// 			TransactionID: req.TransactionID,
// 			HardwareAddr:  req.HardwareAddr,
// 			ClientAddr:    req.ClientAddr,
// 			YourAddr:      net.IPv4(172, 24, 32, 1),
// 			Options:       make(dhcp4.Options),
// 		}
// 		resp.Options[dhcp4.OptSubnetMask] = net.IPv4Mask(255, 255, 0, 0)

// 		switch req.Type { // (3)
// 		case dhcp4.MsgDiscover:
// 			resp.Type = dhcp4.MsgOffer

// 		case dhcp4.MsgRequest:
// 			resp.Type = dhcp4.MsgAck

// 		default:
// 			log.Printf("[WARN] message type %s not supported", req.Type)
// 			continue
// 		}

// 		log.Printf("[INFO] Sending %s to %s", resp.Type, resp.HardwareAddr)
// 		err = conn.SendDHCP(resp, intf) // (4)
// 		if err != nil {
// 			log.Printf("[ERROR] unable to send DHCP packet: %v", err)
// 		}
// 	}
// }
