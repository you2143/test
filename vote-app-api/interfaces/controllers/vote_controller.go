package controllers

import (
	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/service"
)

type voteController struct {
	service service.VoteService
}

type VoteController interface {
	CreateVote(s *request.CreateVote) error
	GetALLVote() (*response.VoteResultSummary, error)
	GetVoteIsEnd() (bool, error)
}

func NewVoteController(st service.VoteService) VoteController {
	return &voteController{st}
}

// CreateVote godoc
// @Summary Create Vote
// @Description create vote by request
// @Accept mpfd
// @Produce json
// @Param movie_id formData int true "movie id is required"
// @Param person_id formData int true "person id is required"
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/vote [post]
func (con *voteController) CreateVote(s *request.CreateVote) error {
	err := con.service.CreateVote(s)
	if err != nil {
		return err
	}
	return nil
}

// GetAllVote godoc
// @Summary Get All Vote
// @Description Get All Vote
// @Accept mpfd
// @Produce json
// @Success 200 {Array} response.VoteResultSummary
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/vote [get]
func (con *voteController) GetALLVote() (*response.VoteResultSummary, error) {
	return con.service.GetALLVote()
}

// VoteIsEnd godoc
// @Summary Check vote is End
// @Description Check vote is End
// @Accept mpfd
// @Produce json
// @Success 200
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/voteIsEnd [get]
func (con *voteController) GetVoteIsEnd() (bool, error) {
	return con.service.GetVoteIsEnd()
}
