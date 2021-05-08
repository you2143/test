package presenters

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type masterMoviePresenter struct{}

func NewMasterMoviePresenter() MasterMoviePresenter {
	return &masterMoviePresenter{}
}

type MasterMoviePresenter interface {
	ResponseMasterMovie(st *model.MasterMovie) *response.MasterMovie
}

func (masterMoviePresenter *masterMoviePresenter) ResponseMasterMovie(st *model.MasterMovie) *response.MasterMovie {
	res := response.MasterMovie{
		ID:    uint64(st.Model.ID),
		Title: st.Title,
	}

	return &res
}
