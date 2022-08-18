package historySend

import "time"

type HistorySend struct {
	Id          string `gorm:"primaryKey"`
	CreatedAt   time.Time
	TypeMessage string
	CompanyId   string `gorm:"primaryKey"`
}
