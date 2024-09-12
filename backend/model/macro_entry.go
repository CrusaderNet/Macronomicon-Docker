package model

import "time"

type MacroEntry struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	UserID     int64     `json:"user_id"`
	SubmitTime time.Time `gorm:"autoCreateTime" json:"submitTime"`
	Proteins   float64   `json:"proteins"`
	Carbs      float64   `json:"carbs"`
	Fats       float64   `json:"fats"`
}

type InsertMacroEntryRequest struct {
	UserID   int64   `json:"user_id"`
	Proteins float64 `json:"proteins"`
	Carbs    float64 `json:"carbs"`
	Fats     float64 `json:"fats"`
}
