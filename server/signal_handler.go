package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//receiver sigterm -> cancel mainctx
func SigTermHandler(cancelMain context.CancelFunc) {
	endChannel := make(chan os.Signal, 1)
	defer close(endChannel)
	signal.Notify(endChannel, syscall.SIGINT, syscall.SIGTERM)

	//stop proccess by stop/ctrl-c
	<-endChannel // block func
	fmt.Println("mainctx was canceled")
	cancelMain() //send true to ctx.Done channel
}
