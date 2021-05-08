package response

import (
	"vote-app-api/domain/model"
)

type MasterPerson struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type MasterPersons struct {
	MasterPersons []model.MasterPerson `json:"MasterPersons"`
}
