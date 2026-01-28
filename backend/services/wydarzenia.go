package services

import (
	"errors"
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

func GenerateAlias(wydarzenie *models.Event) {
	prefix := ""
	wydarzenie.Alias = strings.ReplaceAll(wydarzenie.Nazwa, " ", "-")
	wydarzenie.Alias = strings.ToLower(wydarzenie.Alias)
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

func GetCalendarData(c *gin.Context) {
	var wydarzenia []models.EventCalendarType
	query := config.DB.Model(&models.Event{})
	query = query.Select("ID", "Nazwa", "DataStart", "Alias")
	if err := query.Find(&wydarzenia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wydarzenia)

}

func FindByTitle(c *gin.Context) {
	var wydarzenia []models.EventSummary
	title := "%" + c.Query("title") + "%"
	query := config.DB.Model(&models.Event{})
	query = query.Select("ID", "Nazwa", "Alias", "DataStart", "Aktywne")

	if len(title) > 0 {
		query = query.Order("data_start DESC").Where("nazwa LIKE ?", title)
	} else {
		query = query.Order("data_stop DESC").Where("data_stop > ?", today)
	}

	if err := query.Find(&wydarzenia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wydarzenia)
}

func GetWydarzenieList(c *gin.Context) {
	var wydarzenia []models.EventSummary
	years := c.QueryArray("year")
	title := c.Query("title")
	intYears := make([]int, 0, len(years))
	query := config.DB.Model(&models.Event{})
	query = query.Select("ID", "Nazwa", "Alias", "DataStart", "Aktywne")

	for _, y := range years {
		if val, err := strconv.Atoi(y); err == nil {
			intYears = append(intYears, val)
		}
	}

	query = query.Order("data_start DESC")
	if len(intYears) > 0 {
		query = query.Where("YEAR(data_start) IN ?", intYears)
	}

	if len(title) > 0 {
		query = query.Where("nazwa LIKE ?", "%"+title+"%")
	}

	if len(intYears) == 0 && len(title) == 0 {
		query = query.Where("data_stop > ?", today)
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
	aliasStr := c.Param("wydarzenie")
	if err := config.DB.Where("alias = ?", aliasStr).Delete(&models.Event{}).Error; err != nil {
		c.JSON(400, gin.H{"Error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func SaveWydarzenie(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Zapisano"})
}

func UpdateWydarzenie(c *gin.Context) {
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
