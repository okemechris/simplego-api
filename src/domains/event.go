package domains

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
