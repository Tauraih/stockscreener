package models 

import (
	"time"
)

// Company ..
type Company struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Symbol		   string	    `json:"symbol"`
	Description    string  		`json:"description"`
	Established    string   	`json:"established"`
	CreatedAt 	   time.Time 	`json:"created_at"`
	UpdatedAt      time.Time 	`json:"updated_at"`
}
