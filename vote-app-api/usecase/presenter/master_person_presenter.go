package presenter

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type MasterPersonPresenter interface {
	ResponseMasterPerson(st *model.MasterPerson) *response.MasterPerson
}
