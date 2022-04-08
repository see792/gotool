package server

import (
	"github.com/see792/gotool/config"
	"github.com/see792/gotool/util/os"
)

var appConfig *config.Config

type BaseApp interface {
	Init()
}
func Init(ConfigPath string,app BaseApp) {
	app.Init()
	os.WaitQuit()
}
