package model

import "time"

type MacroEntry struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	UserID     int64     `json:"user_id"`
	SubmitTime time.Time `json:"submitTime"`
	Proteins   float64   `json:"proteins"`
	Carbs      float64   `json:"carbs"`
	Fats       float64   `json:"fats"`
}
