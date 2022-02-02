package domains

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
