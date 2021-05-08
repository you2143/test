package datastore

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"

	"github.com/jinzhu/gorm"
)

type masterMovieRepository struct {
	db *gorm.DB
}

type MasterMovieRepository interface {
	Get(req *request.GetMasterMovie) (*model.MasterMovie, error)
	GetAll() ([]*model.MasterMovie, error)
	Create(req *request.CreateMasterMovie) error
	Update(ctx context.Context, req *request.UpdateMasterMovie) error
	Delete(ctx context.Context, req *request.DeleteMasterMovie) error
}

func NewMasterMovieRepository(db *gorm.DB) MasterMovieRepository {
	return &masterMovieRepository{
		db,
	}
}

func (masterMovieRepository *masterMovieRepository) Get(req *request.GetMasterMovie) (*model.MasterMovie, error) {
	masterMovie := model.MasterMovie{}
	err := masterMovieRepository.
		db.
		First(&masterMovie, req.ID).
		Error

	if err != nil {
		return &masterMovie, err
	}
	return &masterMovie, nil
}

func (masterMovieRepository *masterMovieRepository) GetAll() ([]*model.MasterMovie, error) {
	masterMovie := []*model.MasterMovie{}
	err := masterMovieRepository.
		db.
		Find(&masterMovie).
		Error

	if err != nil {
		return nil, err
	}
	return masterMovie, nil
}

func (masterMovieRepository *masterMovieRepository) Create(req *request.CreateMasterMovie) error {
	s := new(model.MasterMovie)
	s.Title = req.Title

	tx := masterMovieRepository.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&s).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *masterMovieRepository) Update(ctx context.Context, req *request.UpdateMasterMovie) error {
	values := map[string]interface{}{
		"ID":    req.ID,
		"Title": req.Title,
	}

	masterMoview := model.MasterMovie{}
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&masterMoview).Updates(values).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *masterMovieRepository) Delete(ctx context.Context, req *request.DeleteMasterMovie) error {
	masterMovie := model.MasterMovie{}
	err := repo.
		db.
		Delete(&masterMovie, req.ID).
		Error
	if err != nil {
		return err
	}

	return nil
}
