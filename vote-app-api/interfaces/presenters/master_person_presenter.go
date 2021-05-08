package presenters

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type masterPersonPresenter struct{}

func NewMasterPersonPresenter() MasterPersonPresenter {
	return &masterPersonPresenter{}
}

type MasterPersonPresenter interface {
	ResponseMasterPerson(st *model.MasterPerson) *response.MasterPerson
}

func (masterPersonPresenter *masterPersonPresenter) ResponseMasterPerson(st *model.MasterPerson) *response.MasterPerson {
	res := response.MasterPerson{
		ID:   uint64(st.Model.ID),
		Name: st.Name,
	}

	return &res
}
