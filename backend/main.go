package main

import (
	"backend/config"
	"backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":5172")
	fmt.Println("Ala ma kota")
}
