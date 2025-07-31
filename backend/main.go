package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0/0"})

	routes.RegisterRoutes(r)

	r.Run("0.0.0.0:5172")
}
