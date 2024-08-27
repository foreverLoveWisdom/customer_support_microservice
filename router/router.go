package router

import (
	"github.com/foreverLoveWisdom/customer_support_microservice.git/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Register routes
	r.GET("/ping", handlers.Ping)
	r.POST("/tickets", handlers.CreateTicket)
	// r.GET("/tickets", handlers.GetTickets)

	return r
}
