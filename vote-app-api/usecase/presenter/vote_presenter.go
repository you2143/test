package presenter

import (
	"vote-app-api/domain/model"
	"vote-app-api/domain/response"
)

type VotePresenter interface {
	ResponseVote(st []*model.Vote) *response.VoteResultSummary
}
