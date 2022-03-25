package dhcp4

import "fmt"

type DhcpHandler struct {
}

func (d *DhcpHandler) discover() error {
	fmt.Println("aaa")
	return nil
}

func (d *DhcpHandler) request() error {
	fmt.Println("aaa")
	return nil
}

func (d *DhcpHandler) decline() error {
	fmt.Println("aaa")
	return nil
}

func (d *DhcpHandler) release() error {
	fmt.Println("aaa")
	return nil
}

func NewDhcpHandler() *DhcpHandler {
	return new(DhcpHandler)
}
