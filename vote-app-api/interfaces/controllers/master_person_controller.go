package controllers

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/service"
)

type masterPersonController struct {
	service service.MasterPersonService
}

type MasterPersonController interface {
	GetPerson(s *request.GetMasterPerson) (*response.MasterPerson, error)
	GetPersons() ([]*model.MasterPerson, error)
	CreateMasterPerson(s *request.CreateMasterPerson) error
	UpdateMasterPerson(s *request.UpdateMasterPerson) error
	DeleteMasterPerson(s *request.DeleteMasterPerson) error
}

func NewMasterPersonController(st service.MasterPersonService) MasterPersonController {
	return &masterPersonController{st}
}

// GetMasterPerson godoc
// @Summary Get MasterPerson
// @Description Get masterPerson by request
// @Accept mpfd
// @Produce json
// @Param id path int true "person id is required"
// @Success 200 {object} response.MasterPerson
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterPerson/{id} [get]
func (con *masterPersonController) GetPerson(s *request.GetMasterPerson) (*response.MasterPerson, error) {
	masterPersonw, err := con.service.GetPerson(s)
	if err != nil {
		return nil, err
	}
	return masterPersonw, nil
}

// GetAllMasterPerson godoc
// @Summary Get All MasterPerson
// @Description Get all masterPerson by request
// @Accept mpfd
// @Produce json
// @Success 200 {array} model.MasterPerson
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterPerson [get]
func (con *masterPersonController) GetPersons() ([]*model.MasterPerson, error) {
	masterPersonw, err := con.service.GetPersons()
	if err != nil {
		return nil, err
	}
	return masterPersonw, nil
}

// CreateMasterPerson godoc
// @Summary Create MasterPerson
// @Description create masterPerson by request
// @Accept mpfd
// @Produce json
// @Param name formData string true "person name is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterPerson [post]
func (con *masterPersonController) CreateMasterPerson(s *request.CreateMasterPerson) error {
	err := con.service.Create(s)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMasterPerson godoc
// @Summary Update MasterPerson
// @Description Update store by request
// @Accept mpfd
// @Produce json
// @Param id path int true "person id is required"
// @Param name formData string true "name id is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterPerson/{id} [put]
func (con *masterPersonController) UpdateMasterPerson(s *request.UpdateMasterPerson) error {
	return con.service.Update(s)
}

// DeleteMasterPerson godoc
// @Summary Delete MasterPerson
// @Description delete store by request
// @Accept mpfd
// @Produce json
// @Param id path int true "person id is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/masterPerson/{id} [delete]
func (con *masterPersonController) DeleteMasterPerson(s *request.DeleteMasterPerson) error {
	return con.service.Delete(s)
}
