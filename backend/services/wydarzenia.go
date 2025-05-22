package services

import (
	"log"

	"strconv"

	"net/http"

	"backend/models"

	"backend/config"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetWydarzenieList(c *gin.Context) {
	var wydarzenia []models.Wydarzenia
	if err := config.DB.Find(&wydarzenia).Error; err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, wydarzenia)
}

func GetWydarzenie(c *gin.Context) {
	var wydarzenie models.Wydarzenia
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err = config.DB.First(&wydarzenie, id).Error; err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, wydarzenie)
}

func DeleteWydarzenie(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err = config.DB.Delete(&models.Wydarzenia{}, id).Error; err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func InsertWydarzenie(db *gorm.DB) {
	wydarzenie := models.Wydarzenia{}
	db.Create(&wydarzenie)
}
