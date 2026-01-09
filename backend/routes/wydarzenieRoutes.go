package routes

import (
	"backend/CAS"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.SetTrustedProxies([]string{"172.20.0.2"})

	router.GET("/api/login", CAS.LoginHandler)
	router.GET("/api/callback", CAS.CallbackHandler)

	router.GET("/api/all", services.GetWydarzenieList)
	router.GET("/api/:wydarzenie", services.GetWydarzenie)
	router.POST("/api/zapisz", services.SaveWydarzenie)
	router.DELETE("/api/:wydarzenie", services.DeleteWydarzenie)
	router.PUT("/api/:alias", services.UpdateWydarzenie)
	router.PUT("/api/visibility-switcher/:alias", services.VisibilitySwitcher)
	router.GET("/api/isAdmin", services.IsAdmin)
	router.GET("/api/allC", services.GetCalendarData)
}
