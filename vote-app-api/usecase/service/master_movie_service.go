package service

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/presenter"
	"vote-app-api/usecase/repository"
)

type masterMovieService struct {
	MasterMovieRepository repository.MasterMovieRepository
	MasterMoviePresenter  presenter.MasterMoviePresenter
}

type MasterMovieService interface {
	GetMovie(s *request.GetMasterMovie) (*response.MasterMovie, error)
	GetMovies() ([]*model.MasterMovie, error)
	Create(s *request.CreateMasterMovie) error
	Update(s *request.UpdateMasterMovie) error
	Delete(s *request.DeleteMasterMovie) error
}

func NewMasterMovieService(repo repository.MasterMovieRepository, pre presenter.MasterMoviePresenter) MasterMovieService {
	return &masterMovieService{repo, pre}
}

func (masterMovieService *masterMovieService) GetMovie(s *request.GetMasterMovie) (*response.MasterMovie, error) {
	movie, err := masterMovieService.MasterMovieRepository.Get(s)
	if err != nil {
		return &response.MasterMovie{}, err
	}
	return masterMovieService.MasterMoviePresenter.ResponseMasterMovie(movie), nil
}

func (masterMovieService *masterMovieService) GetMovies() ([]*model.MasterMovie, error) {
	movie, err := masterMovieService.MasterMovieRepository.GetAll()
	if err != nil {
		return []*model.MasterMovie{}, err
	}
	return movie, nil
}

func (masterMovieService *masterMovieService) Create(s *request.CreateMasterMovie) error {
	err := masterMovieService.MasterMovieRepository.Create(s)
	if err != nil {
		return err
	}
	return nil
}

func (masterMovieService *masterMovieService) Update(s *request.UpdateMasterMovie) error {
	ctx := context.Background()
	return masterMovieService.MasterMovieRepository.Update(ctx, s)
}

func (masterMovieService *masterMovieService) Delete(s *request.DeleteMasterMovie) error {
	ctx := context.Background()
	return masterMovieService.MasterMovieRepository.Delete(ctx, s)
}
