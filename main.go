package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	//Define Global Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//Version One Routes
	v1 := r.Group("/v1")
	{
		v1.GET("/user/:user", getUser)
	}

	r.Run()
}
