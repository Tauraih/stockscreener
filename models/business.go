package models 

import ( 
	"time"
)

// Business ...
type Business struct {
	ID			int				`gorm:"primary_key;auto_increment" json:"id"`
	Name           string       `json:"name"`
	Symbol		   string	    `json:"symbol"`
	High		float32			`json:"high"`
	Low			float32			`json:"low"`
	Opening		float32			`json:"opening"`
	Closing		float32			`json:"closing"`
	LastTraded	float32			`json:"lastTraded"`
	Change		float32			`json:"change"`
	CreatedAt   time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
