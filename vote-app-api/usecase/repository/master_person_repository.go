package repository

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
)

type MasterMovieRepository interface {
	Get(req *request.GetMasterMovie) (*model.MasterMovie, error)
	GetAll() ([]*model.MasterMovie, error)
	Create(req *request.CreateMasterMovie) error
	Update(ctx context.Context, req *request.UpdateMasterMovie) error
	Delete(ctx context.Context, req *request.DeleteMasterMovie) error
}
