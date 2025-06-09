package services

import (
	"log"

	"strconv"

	"net/http"

	"backend/models"

	"backend/config"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"time"
)

var now = time.Now()
var today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

func GetWydarzenieList(c *gin.Context) {
	var wydarzenia []models.Wydarzenia
	if err := config.DB.Order("data_start DESC").Where("data_stop >= ?", today).Find(&wydarzenia).Error; err != nil {
		log.Println(err)
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
