package router

import (
	"encoding/gob"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"networklistener/controller"
	"time"
)

func New(secret string) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//router.Use(middlewares.AuthMiddleware(secret))
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	bridgeController := controller.BridgeController{}
	router.POST("/sub/create", bridgeController.CreateBridge())
	return router
}
