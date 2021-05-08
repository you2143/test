package presenters

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type votePresenter struct{}

func NewVotePresenter() VotePresenter {
	return &votePresenter{}
}

type VotePresenter interface {
	ResponseVote(st []*model.Vote) *response.VoteResultSummary
}

func (votePresenter *votePresenter) ResponseVote(st []*model.Vote) *response.VoteResultSummary {
	res := response.VoteResultSummary{}
	for _, s := range st {
		voteResult := response.VoteResult{}
		voteResult.Title = s.Movie.Title
		voteResult.VoteCount = 1

		res.VoteResults = append(res.VoteResults, voteResult)
	}
	return &res
}
