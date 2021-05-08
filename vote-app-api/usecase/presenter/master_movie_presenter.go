package presenter

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type MasterMoviePresenter interface {
	ResponseMasterMovie(st *model.MasterMovie) *response.MasterMovie
}
