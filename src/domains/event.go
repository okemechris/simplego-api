package domains

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
