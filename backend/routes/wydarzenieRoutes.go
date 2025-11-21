package routes

import (
	"backend/CAS"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.SetTrustedProxies([]string{"172.20.0.2"})

	casMiddleware, _ := CAS.SetupCASMiddleware("https://cas.al.edu.pl/cas")
	router.GET("/api/login", CAS.LoginHandler)
	router.GET("/api/getusername", casMiddleware, CAS.GetUserID)

	router.GET("/api/all", services.GetWydarzenieList)
	router.GET("/api/:wydarzenie", services.GetWydarzenie)
	router.POST("/api/zapisz", services.SaveWydarzenie)
	router.DELETE("/api/:wydarzenie", services.DeleteWydarzenie)
	router.PUT("/api/:alias", services.UpdateWydarzenie)
	router.PUT("/api/visibility-switcher/:alias", services.VisibilitySwitcher)
	router.GET("/api/isAdmin", services.IsAdmin)
	router.GET("/api/allC", services.GetCalendarData)
	router.GET("/api/allT", services.FindByTitle)
}
