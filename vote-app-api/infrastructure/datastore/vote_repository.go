package datastore

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"

	"github.com/jinzhu/gorm"
)

type voteRepository struct {
	db *gorm.DB
}

type VoteRepository interface {
	Create(req *request.CreateVote) error
	GetALL(ctx context.Context) ([]*model.Vote, error)
	GetVoteIsEnd(ctx context.Context) (bool, error)
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepository{
		db,
	}
}

func (voteRepository *voteRepository) Create(req *request.CreateVote) error {
	vote := new(model.Vote)
	movie := model.MasterMovie{}
	person := model.MasterPerson{}

	err := voteRepository.
		db.
		Find(movie, req.MovieID).
		Error

	if err != nil {
		return err
	}

	err = voteRepository.
		db.
		Find(person, req.PersonID).
		Error

	if err != nil {
		return err
	}

	vote.Movie = movie
	vote.Person = person

	tx := voteRepository.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&vote).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *voteRepository) GetALL(ctx context.Context) ([]*model.Vote, error) {
	vote := []*model.Vote{}
	err := repo.
		db.
		Find(&vote).
		Error

	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (repo *voteRepository) GetVoteIsEnd(ctx context.Context) (bool, error) {
	vote := []*model.Vote{}
	err := repo.
		db.
		Find(&vote).
		Error

	if err != nil {
		return false, err
	}
	return true, nil
}
