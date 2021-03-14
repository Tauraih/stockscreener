package models

import ( 
	"time"
)

type New struct {
	ID             int          `json:"id"`
	Number, Name, Symbol, High, Low, Opening, Last, Closing, Change string
	CreatedAt   time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}