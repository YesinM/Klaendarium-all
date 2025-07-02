package routes

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/all", services.GetWydarzenieList)
	router.GET("/api/:wydarzenie", services.GetWydarzenie)
	router.POST("/api/zapisz", services.SaveWydarzenie)
}
