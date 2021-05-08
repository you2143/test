package repository

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
)

type VoteRepository interface {
	Create(req *request.CreateVote) error
	GetALL(ctx context.Context) ([]*model.Vote, error)
	GetVoteIsEnd(ctx context.Context) (bool, error)
}
