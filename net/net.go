package net

import "github.com/see792/gotool/config"

func Init(config *config.Config) {
	InitWeb(&config.WebServer)
	InitSocket(&config.SocketServer)
	InitWebSocket(&config.WebSocketServer)

}