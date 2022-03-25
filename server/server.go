package server

import (
	"fmt"
	"sync"
)

type Server interface {
	Run()
	Close()
}

// サーバをgorutineで起動
// sigterm受信でサーバを停止
func Run(servers []Server) {

	for _, server := range servers {
		go server.Run()
		fmt.Println("[server] server run")
	}

	// <-mainCtx.Done() // wait ctl-c or kill until cancel func call
	// Close(servers, wg)
}

func Close(servers []Server, wg *sync.WaitGroup) {
	for _, server := range servers {
		fmt.Println("[server] call server.Close")
		server.Close()
		wg.Done()
	}
}
