package main

import (
	"context"
	"fmt"
	"os"
	"pxe/config"
	"pxe/logger"
	"pxe/server"
	"pxe/server/dhcp4"
	"pxe/server/tftp"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// UDP でクライアント毎にコネクションを作成する
	// https://kechako.dev/posts/2018/12/03/create-udp-conn-each-client/
	var servers []server.Server
	tftp.Tftp()
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println("[main] Load config failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	mainLogger := logger.NewLogger("main")
	mainLogger.Info("Test")
	closeChan := make(chan bool)
	defer close(closeChan)
	if config.Dhcp.Enabled {
		servers = append(servers, dhcp4.NewDhcpServer(config, &closeChan))
		fmt.Println("adding wg")
		wg.Add(1)
	}
	mainCtx, cancel := context.WithCancel(context.Background())
	// if config.Tftp.Enabled {
	// 	servers = append(servers, dhcp4.NewDhcpServer(config))
	// }
	// if config.Http.Enabled {
	// 	servers = append(servers, dhcp4.NewDhcpServer(config))
	// }
	go server.SigTermHandler(cancel) //receive sigterm -> cancel mainct
	server.Run(servers)
	// wg.Wait() //cancel mainctx -> wg.Done() -> here

	mainLogger.Info("waiting ctr-c...")
	<-mainCtx.Done()
	mainLogger.Info("Close process started")
	closeChan <- true
	go server.Close(servers, &wg)
	mainLogger.Info("waiting all process ends...")
	wg.Wait()

	// server.Close(servers)
	mainLogger.Info("finish process")
}
