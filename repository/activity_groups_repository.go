package repository

import (
	"api_activity/entity"
	"errors"

	"gorm.io/gorm"
)

type ActivityGroupRepo interface {
	FindByID(id int64) (entity.ActivityGroups, error)
	FindAll() ([]entity.ActivityGroups, error)
	Create(entity.ActivityGroups) (entity.ActivityGroups, error)
	DeleteByID(id int64) error
	UpdateByID(data map[string]interface{}, id int64) (entity.ActivityGroups, error)
}

type activityGroupRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ActivityGroupRepo {
	return &activityGroupRepo{db}
}
func (r *activityGroupRepo) FindByID(id int64) (entity.ActivityGroups, error) {
	var acg entity.ActivityGroups
	var err error
	query := r.db.Find(&acg,id)
	if query.RowsAffected == 0 {
		err = errors.New("Not Found")
	}
	return acg, err
}
func (r *activityGroupRepo) FindAll() ([]entity.ActivityGroups, error) {
	var acg []entity.ActivityGroups
	err := r.db.Find(&acg).Error
	return acg, err
}
func (r *activityGroupRepo) Create(input entity.ActivityGroups) (entity.ActivityGroups, error) {
	err := r.db.Create(&input).Error
	return input, err
}
func (r *activityGroupRepo) DeleteByID(id int64) error {
	var acg entity.ActivityGroups
	err := r.db.Where("id =", id).Delete(&acg).Error
	return err
}
func (r *activityGroupRepo) UpdateByID(data map[string]interface{}, id int64) (entity.ActivityGroups, error) {
	var acg entity.ActivityGroups
	err := r.db.Model(&acg).Where("id =", id).Updates(&data).Error
	return acg, err
}
