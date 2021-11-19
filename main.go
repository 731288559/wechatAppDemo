package main

import (
	"demo/helper"
	"demo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	helper.Init()

	app := gin.Default()
	router.Http(app)

	if err := app.Run(":8081"); err != nil {
		panic(err.Error())
	}
}
