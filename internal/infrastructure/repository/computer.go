package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"screencapturer/internal/domain/computer"
	"screencapturer/internal/domain/model"
)

type computerRepository struct {
	DB *gorm.DB
}

var _ computer.ComputerRepository = (*computerRepository)(nil)

func NewComputerRepository(DB *gorm.DB) *computerRepository {
	return &computerRepository{DB: DB}
}

func (u *computerRepository) FindAll() ([]*model.Computer, error) {
	query := u.DB.Model(&model.Computer{})

	query = query.Order("name ASC")

	var computers []*model.Computer

	if err := query.Find(&computers).Error; err != nil {
		return nil, err
	}

	return computers, nil
}

func (u *computerRepository) FindById(id int) (*model.Computer, error) {
	var computer model.Computer

	err := u.DB.Where("id = ?", id).First(&computer).Error

	return &computer, err
}

func (u *computerRepository) Create(computer model.Computer) (*model.Computer, error) {
	err := u.DB.Save(&computer).Error
	if err != nil {
		return nil, err
	}

	return &computer, nil
}

func (u *computerRepository) Update(computer model.Computer) (*model.Computer, error) {
	err := u.DB.Model(&computer).Clauses(clause.Returning{}).Updates(&computer).Error
	if err != nil {
		return nil, err
	}

	return &computer, nil
}

func (u *computerRepository) Delete(computer model.Computer) error {
	return u.DB.Delete(&computer).Error
}
