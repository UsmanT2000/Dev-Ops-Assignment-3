package main

import (
	"github.com/UsmanT2000/ginAPIs/Users"
	"github.com/UsmanT2000/ginAPIs/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/user", Users.GetUsers)
	router.GET("/user/:id", Users.GetUserById)
	router.POST("/user", Users.CreateUser)
	router.PATCH("/user/:id", Users.UpdateUser)
	router.DELETE("/user/:id", Users.DeleteUser)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/parse", handleParseCSV)
	router.GET("/hello", SayHello)
	router.POST("/signin", middleware.LoginEndpoint)

	authorized := router.Group("/auth")
	authorized.Use(middleware.AuthMiddleware())
	{
		testing := authorized.Group("/testing")
		testing.GET("/analytics", middleware.AnalyticsEndpoint)
	}

	router.Run("localhost:10000")
}
