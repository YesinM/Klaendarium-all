package models

import "time"

type Event struct {
	ID          uint       `gorm:"primaryKey;column:id"`
	Nazwa       string     `gorm:"column:nazwa"`
	Alias       string     `gorm:"column:alias"`
	Opis        string     `gorm:"column:opis";type:"text"`
	DataStart   *time.Time `gorm:"column:data_start"`
	DataStop    *time.Time `gorm:"column:data_stop"`
	Organizator string     `gorm:"column:organizator"`
	Lokalizacja string     `gorm:"column:lokalizacja"`
	Aktywne     bool       `gorm:"column:aktywne"`
}

type EventSummary struct {
	ID        uint       `gorm:"primaryKey;column:id"`
	Nazwa     string     `gorm:"column:nazwa"`
	Alias     string     `gorm:"column:alias"`
	DataStart *time.Time `gorm:"column:data_start"`
	Aktywne   bool       `gorm:"column:aktywne"`
}

type EventCalendarType struct {
	ID        uint       `gorm:"primaryKey;column:id"`
	Nazwa     string     `gorm:"column:nazwa"`
	DataStart *time.Time `gorm:"column:data_start"`
	Alias     string     `gorm:"column:alias"`
}
