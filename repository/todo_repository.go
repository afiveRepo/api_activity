package repository

import (
	"api_activity/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindByID(id int64) (entity.Todo, error)
	FindByGroupID(id int64) ([]entity.Todo, error)
	FindAll() ([]entity.Todo, error)
	Create(entity.Todo) (entity.Todo, error)
	DeleteByID(id int64) error
	UpdateByID(data map[string]interface{}, id int64) (entity.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}
func (r *todoRepository) FindByID(id int64) (entity.Todo, error) {
	var todo entity.Todo
	err := r.db.Find(&todo, id).Error
	return todo, err
}
func (r *todoRepository) FindByGroupID(id int64) ([]entity.Todo, error) {
	var todo []entity.Todo
	err := r.db.Where("activity_group_id = ", id).Find(&todo).Error
	return todo, err
}
func (r *todoRepository) FindAll() ([]entity.Todo, error) {
	var todo []entity.Todo
	err := r.db.Find(&todo).Error
	return todo, err
}
func (r *todoRepository) Create(entity.Todo) (entity.Todo, error) {
	var todo entity.Todo
	err := r.db.Create(&todo).Error
	return todo, err
}
func (r *todoRepository) DeleteByID(id int64) error {
	var todo entity.Todo
	err := r.db.Where("id =", id).Delete(&todo).Error
	return err
}
func (r *todoRepository) UpdateByID(data map[string]interface{}, id int64) (entity.Todo, error) {
	var todo entity.Todo
	err := r.db.Model(&todo).Where("id =", id).Updates(&data).Error
	return todo, err
}
