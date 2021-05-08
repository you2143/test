package model

import (
	"github.com/jinzhu/gorm"
)

type MasterMovie struct {
	gorm.Model
	Title string `json:"title" sql:"not null"`
}
