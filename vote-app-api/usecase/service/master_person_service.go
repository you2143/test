package service

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/presenter"
	"vote-app-api/usecase/repository"
)

type masterPersonService struct {
	MasterPersonRepository repository.MasterPersonRepository
	MasterPersonPresenter  presenter.MasterPersonPresenter
}

type MasterPersonService interface {
	GetPerson(s *request.GetMasterPerson) (*response.MasterPerson, error)
	GetPersons() ([]*model.MasterPerson, error)
	Create(s *request.CreateMasterPerson) error
	Update(s *request.UpdateMasterPerson) error
	Delete(s *request.DeleteMasterPerson) error
}

func NewMasterPersonService(repo repository.MasterPersonRepository, pre presenter.MasterPersonPresenter) MasterPersonService {
	return &masterPersonService{repo, pre}
}

func (masterPersonService *masterPersonService) GetPerson(s *request.GetMasterPerson) (*response.MasterPerson, error) {
	person, err := masterPersonService.MasterPersonRepository.Get(s)
	if err != nil {
		return &response.MasterPerson{}, err
	}
	return masterPersonService.MasterPersonPresenter.ResponseMasterPerson(person), nil
}

func (masterPersonService *masterPersonService) GetPersons() ([]*model.MasterPerson, error) {
	person, err := masterPersonService.MasterPersonRepository.GetAll()
	if err != nil {
		return []*model.MasterPerson{}, err
	}
	return person, nil
}

func (masterPersonService *masterPersonService) Create(s *request.CreateMasterPerson) error {
	err := masterPersonService.MasterPersonRepository.Create(s)
	if err != nil {
		return err
	}
	return nil
}

func (masterPersonService *masterPersonService) Update(s *request.UpdateMasterPerson) error {
	ctx := context.Background()
	return masterPersonService.MasterPersonRepository.Update(ctx, s)
}

func (masterPersonService *masterPersonService) Delete(s *request.DeleteMasterPerson) error {
	ctx := context.Background()
	err := masterPersonService.MasterPersonRepository.Delete(ctx, s)
	if err != nil {
		return err
	}
	return nil
}
