package controllers

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/service"
)

type masterMovieController struct {
	service service.MasterMovieService
}

type MasterMovieController interface {
	GetMovie(s *request.GetMasterMovie) (*response.MasterMovie, error)
	GetMovies() ([]*model.MasterMovie, error)
	CreateMovie(s *request.CreateMasterMovie) error
	UpdateMovie(s *request.UpdateMasterMovie) error
	DeleteMovie(s *request.DeleteMasterMovie) error
}

func NewMasterMovieController(st service.MasterMovieService) MasterMovieController {
	return &masterMovieController{st}
}

// GetMasterMovie godoc
// @Summary Get MasterMovie
// @Description Get masterMovie by request
// @Accept mpfd
// @Produce json
// @Param id path int true "movie id is required"
// @Success 200 {object} response.MasterMovie
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterMovie/{id} [get]
func (con *masterMovieController) GetMovie(s *request.GetMasterMovie) (*response.MasterMovie, error) {
	masterMoview, err := con.service.GetMovie(s)
	if err != nil {
		return nil, err
	}
	return masterMoview, nil
}

// GetAllMasterMovie godoc
// @Summary Get All MasterMovie
// @Description Get all masterMovie by request
// @Accept mpfd
// @Produce json
// @Success 200 {array} model.MasterMovie
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterMovie [get]
func (con *masterMovieController) GetMovies() ([]*model.MasterMovie, error) {
	masterMoview, err := con.service.GetMovies()
	if err != nil {
		return nil, err
	}
	return masterMoview, nil
}

// CreateMasterMovie godoc
// @Summary Create MasterMovie
// @Description create masterMovie by request
// @Accept mpfd
// @Produce json
// @Param title formData string true "movie title is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterMovie [post]
func (con *masterMovieController) CreateMovie(s *request.CreateMasterMovie) error {
	err := con.service.Create(s)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMasterMovie godoc
// @Summary Update MasterMovie
// @Description Update store by request
// @Accept mpfd
// @Produce json
// @Param id path int true "master movie id is required"
// @Param title formData string true "title id is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterMovie/{id} [put]
func (con *masterMovieController) UpdateMovie(s *request.UpdateMasterMovie) error {
	return con.service.Update(s)
}

// DeleteMasterMovie godoc
// @Summary Delete MasterMovie
// @Description delete store by request
// @Accept mpfd
// @Produce json
// @Param id path int true "master movie id is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterMovie/{id} [delete]
func (con *masterMovieController) DeleteMovie(s *request.DeleteMasterMovie) error {
	return con.service.Delete(s)
}
