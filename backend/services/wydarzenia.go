package services

import (
	"fmt"
	"log"

	"strconv"

	"net/http"

	"backend/models"

	"backend/config"

	"github.com/gin-gonic/gin"

	"time"
)

var now = time.Now()
var today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

func GetWydarzenieList(c *gin.Context) {
	var wydarzenia []models.Wydarzenia
	years := c.QueryArray("year")
	intYears := make([]int, 0, len(years))

	query := config.DB.Model(&models.Wydarzenia{})
	for _, y := range years {
		if val, err := strconv.Atoi(y); err == nil {
			intYears = append(intYears, val)
		}
	}

	if len(intYears) > 0 {
		query = query.Order("data_start DESC").Where("YEAR(data_start) IN ?", intYears)
	} else {
		query = query.Order("data_start DESC").Where("data_start > ?", today)
	}

	if err := query.Find(&wydarzenia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wydarzenia)
}

func GetWydarzenie(c *gin.Context) {
	var wydarzenie models.Wydarzenia
	aliasStr := c.Param("wydarzenie")
	fmt.Println("alias:", aliasStr)
	if err := config.DB.Where("alias = ?", aliasStr).First(&wydarzenie).Error; err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, wydarzenie)
}

func DeleteWydarzenie(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Delete(&models.Wydarzenia{}, id).Error; err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func SaveWydarzenie(c *gin.Context) {
	wydarzenie := models.Wydarzenia{}
	config.Connect()

	if err := c.ShouldBindJSON(&wydarzenie); err != nil {
		c.JSON(400, gin.H{"Error:": err.Error()})
	}

	if err := config.DB.Create(&wydarzenie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zapisać"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Zapisane"})
}

func UpdateWydarzenie(c *gin.Context) {

}
