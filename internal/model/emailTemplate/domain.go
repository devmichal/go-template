package emailTemplate

import (
	"time"
)

type Template struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time
	Template  string `json:"template" gorm:'text(5000)'`
	Type      int    `json:"type"`
}
