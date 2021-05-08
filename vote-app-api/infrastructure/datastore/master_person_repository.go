package datastore

import (
	"context"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"

	"github.com/jinzhu/gorm"
)

type masterPersonRepository struct {
	db *gorm.DB
}

type MasterPersonRepository interface {
	Get(req *request.GetMasterPerson) (*model.MasterPerson, error)
	GetAll() ([]*model.MasterPerson, error)
	Create(req *request.CreateMasterPerson) error
	Update(ctx context.Context, req *request.UpdateMasterPerson) error
	Delete(ctx context.Context, req *request.DeleteMasterPerson) error
}

func NewMasterPersonRepository(db *gorm.DB) MasterPersonRepository {
	return &masterPersonRepository{
		db,
	}
}

func (masterPersonRepository *masterPersonRepository) Get(req *request.GetMasterPerson) (*model.MasterPerson, error) {
	masterPerson := model.MasterPerson{}
	err := masterPersonRepository.
		db.
		First(&masterPerson, req.ID).
		Error

	if err != nil {
		return &masterPerson, err
	}
	return &masterPerson, nil
}

func (masterPersonRepository *masterPersonRepository) GetAll() ([]*model.MasterPerson, error) {
	masterPerson := []*model.MasterPerson{}
	err := masterPersonRepository.
		db.
		Find(&masterPerson).
		Error

	if err != nil {
		return nil, err
	}
	return masterPerson, nil
}

func (masterPersonRepository *masterPersonRepository) Create(req *request.CreateMasterPerson) error {
	s := new(model.MasterPerson)
	s.Name = req.Name

	tx := masterPersonRepository.db.Begin()
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

func (repo *masterPersonRepository) Update(ctx context.Context, req *request.UpdateMasterPerson) error {
	values := map[string]interface{}{
		"ID":   req.ID,
		"Name": req.Name,
	}

	masterPerson := model.MasterPerson{}
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&masterPerson).Updates(values).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *masterPersonRepository) Delete(ctx context.Context, req *request.DeleteMasterPerson) error {
	masterPerson := model.MasterPerson{}
	err := repo.
		db.
		Delete(&masterPerson, req.ID).
		Error
	if err != nil {
		return err
	}

	return nil
}
