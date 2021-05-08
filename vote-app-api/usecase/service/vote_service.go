package service

import (
	"context"

	"vote-app-api/domain/request"
	"vote-app-api/domain/response"
	"vote-app-api/usecase/presenter"
	"vote-app-api/usecase/repository"
)

type voteService struct {
	VoteRepository repository.VoteRepository
	VotePresenter  presenter.VotePresenter
}

type VoteService interface {
	CreateVote(s *request.CreateVote) error
	GetALLVote() (*response.VoteResultSummary, error)
	GetVoteIsEnd() (bool, error)
}

func NewVoteService(repo repository.VoteRepository, pre presenter.VotePresenter) VoteService {
	return &voteService{repo, pre}
}

func (voteService *voteService) CreateVote(s *request.CreateVote) error {
	err := voteService.VoteRepository.Create(s)
	if err != nil {
		return err
	}
	return nil
}

func (voteService *voteService) GetALLVote() (*response.VoteResultSummary, error) {
	ctx := context.Background()
	vote, err := voteService.VoteRepository.GetALL(ctx)
	if err != nil {
		return &response.VoteResultSummary{}, err
	}
	return voteService.VotePresenter.ResponseVote(vote), nil
}

func (voteService *voteService) GetVoteIsEnd() (bool, error) {
	ctx := context.Background()
	res, err := voteService.VoteRepository.GetVoteIsEnd(ctx)
	if err != nil {
		return false, err
	}
	return res, nil
}
