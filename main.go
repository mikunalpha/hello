package main

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func main() {
	Init()

	Router.Run(":5678")
}

func Init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()

	Router = gin.New()

	Router.GET("/hello/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello " + c.Param("name"),
		})
	})
}
