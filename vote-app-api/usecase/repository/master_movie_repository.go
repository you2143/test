package repository

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
)

type MasterPersonRepository interface {
	Get(req *request.GetMasterPerson) (*model.MasterPerson, error)
	GetAll() ([]*model.MasterPerson, error)
	Create(req *request.CreateMasterPerson) error
	Update(ctx context.Context, req *request.UpdateMasterPerson) error
	Delete(ctx context.Context, req *request.DeleteMasterPerson) error
}
