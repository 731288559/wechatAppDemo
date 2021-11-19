package helper

import (
	"demo/conf"
)

func Init() {
	preInit()
	initDB()
}

func preInit() {
	conf.InitConf()
}
