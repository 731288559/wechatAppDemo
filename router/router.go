package router

import (
	"demo/controller"

	"github.com/gin-gonic/gin"
)

func Http(app *gin.Engine) {
	// app.Use(middleware.Recovery)

	app.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "home"})
	})
	g1 := app.Group("/a")
	{
		g1_1 := g1.Group("/aa")
		{
			g1_1.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "pong1"})
			})
		}
	}
	g2 := app.Group("/b")
	{
		g2.GET("/ping2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong2",
			})
		})
	}

	user := app.Group("/user")
	user.GET("/list", controller.MysqlTest)
}
