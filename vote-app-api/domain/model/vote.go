package model

import (
	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	Movie  MasterMovie  `json:"movie" gorm:"foreignkey:MovieID"`
	Person MasterPerson `json:"person" gorm:"foreignkey:PersonID"`
}
