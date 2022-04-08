package net

import (
	"github.com/see792/gotool/config"
	"fmt"
)

func InitSocket(cf *config.SocketServer){
	if !cf.Enable {
		return
	}
	fmt.Println("Socket Enable Port :",cf.PORT)


}
