package models

import "time"

type Wydarzenia struct {
	ID          uint       `gorm:"primaryKey;column:id"`
	Nazwa       string     `gorm:"column:nazwa"`
	Alias       string     `gorm:"column:alias"`
	Opis        string     `gorm:"column:opis"`
	DataStart   *time.Time `gorm:"column:data_start"`
	DataStop    *time.Time `gorm:"column:data_stop"`
	Organizator string     `gorm:"column:organizator"`
	Lokalizacja string     `gorm:"column:lokalizacja"`
	Aktywne     bool       `gorm:"column:aktywne"`
	UtwID       string     `gorm:"column:utw_id"`
	//UtwData     time.Time `gorm:"column:utw_data"`
	//ModID            string    `gorm:"column:mod_id"`
	//ModData          time.Time `gorm:"column:mod_data"`
	ArtykulPowiazany string `gorm:"column:artykul_powiazany"`
	Usuniety         int    `gorm:"column:usuniety"`
}

type WydarzenieSummary struct {
	ID        uint       `gorm:"primaryKey;column:id"`
	Nazwa     string     `gorm:"column:nazwa"`
	Alias     string     `gorm:"column:alias"`
	DataStart *time.Time `gorm:"column:data_start"`
	Aktywne   bool       `gorm:"column:aktywne"`
}
