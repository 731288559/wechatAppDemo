package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer()
	// httpServer()
}

func ginServer() {
	router := gin.Default()
	v1 := router.Group("/a")
	{
		v1_1 := v1.Group("/aa")
		{
			v1_1.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "pong1"})
			})
		}
	}
	v2 := router.Group("/b")
	{
		v2.GET("/ping2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong2",
			})
		})
	}

	router.Run(":8080")
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
