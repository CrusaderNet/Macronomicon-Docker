package model

type User struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Email string `json:"email"`
}
