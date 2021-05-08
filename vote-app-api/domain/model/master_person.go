package model

import (
	"github.com/jinzhu/gorm"
)

type MasterPerson struct {
	gorm.Model
	Name string `json:"name" sql:"not null"`
}
