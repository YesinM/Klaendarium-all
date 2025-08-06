package services

import (
	"errors"
	"fmt"
	"log"
	"strings"

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
var adminStatus = true

func IsAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"admin": adminStatus})
}

func GenerateAlias(wydarzenie *models.Event) {
	prefix := ""
	wydarzenie.Alias = strings.ReplaceAll(wydarzenie.Nazwa, " ", "-")
	wydarzenie.Alias = strings.ToLower(wydarzenie.Alias)
	//var count int64
	// config.DB.Where("alias LIKE ?", wydarzenie.Alias+"%").Count(&count)
	// if count != 0 {
	// 	wydarzenie.Alias = wydarzenie.Alias + strconv.FormatInt(count + 1)
	// }
	//dodanie do aliasu prefiksa, jeśli nie jest unikalny
	for i := 0; ; i++ {
		tempEvent := models.Event{}
		if err := config.DB.Where("alias = ?", wydarzenie.Alias+prefix).First(&tempEvent).Error; err == gorm.ErrRecordNotFound {
			wydarzenie.Alias = wydarzenie.Alias + prefix
			break
		} else {
			prefix = "-" + strconv.Itoa(i+1)
		}
	}
}

func GetWydarzenieList(c *gin.Context) {
	var wydarzenia []models.EventSummary
	years := c.QueryArray("year")
	intYears := make([]int, 0, len(years))

	query := config.DB.Model(&models.Event{})
	query = query.Select("ID", "Nazwa", "Alias", "DataStart", "Aktywne")
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
	var wydarzenie models.Event
	aliasStr := c.Param("wydarzenie")
	if err := config.DB.Where("alias = ?", aliasStr).First(&wydarzenie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, wydarzenie)
}

func DeleteWydarzenie(c *gin.Context) {
	if !adminStatus {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	aliasStr := c.Param("wydarzenie")
	if err := config.DB.Where("alias = ?", aliasStr).Delete(&models.Event{}).Error; err != nil {
		c.JSON(400, gin.H{"Error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func SaveWydarzenie(c *gin.Context) {

	if !adminStatus {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	wydarzenie := models.Event{}

	config.Connect()

	if err := c.ShouldBindJSON(&wydarzenie); err != nil {
		c.JSON(400, gin.H{"Error:": err.Error()})
		return
	}

	GenerateAlias(&wydarzenie)

	if err := config.DB.Create(&wydarzenie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zapisać"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Zapisane"})
}

func UpdateWydarzenie(c *gin.Context) {

	if !adminStatus {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	alias := c.Param("alias")
	wydarzenie := models.Event{}
	config.Connect()

	if err := c.ShouldBindJSON(&wydarzenie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	previousWydarzenie := models.Event{}
	if err := config.DB.Where("Alias = ?", alias).First(&previousWydarzenie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	GenerateAlias(&wydarzenie)

	previousWydarzenie.Nazwa = wydarzenie.Nazwa
	previousWydarzenie.Alias = wydarzenie.Alias
	previousWydarzenie.Opis = wydarzenie.Opis
	previousWydarzenie.DataStart = wydarzenie.DataStart
	previousWydarzenie.DataStop = wydarzenie.DataStop
	previousWydarzenie.Organizator = wydarzenie.Organizator
	previousWydarzenie.Lokalizacja = wydarzenie.Lokalizacja

	if err := config.DB.Save(&previousWydarzenie).Error; err != nil {
		log.Println("Błąd podczas zapisu do bazy:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func VisibilitySwitcher(c *gin.Context) {

	if !adminStatus {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	alias := c.Param("alias")
	wydarzenie := models.Event{}
	config.Connect()
	if err := config.DB.Where("Alias = ?", alias).First(&wydarzenie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nie znaleziono wydarzenia"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd bazy danych"})
	}
	wydarzenie.Aktywne = !wydarzenie.Aktywne
	if err := config.DB.Save(&wydarzenie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd serwera"})
	}
	if wydarzenie.Aktywne {
		c.JSON(http.StatusOK, gin.H{"message": "Wydarzenie widoczne"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wydarzenie nie widoczne"})
}

func realIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	fmt.Println("ClientIP:", c.ClientIP())
	fmt.Println("All headers:", c.Request.Header)

	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}
	return c.ClientIP()
}
