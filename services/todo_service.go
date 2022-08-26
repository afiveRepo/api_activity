package services

import (
	"api_activity/entity"
	"api_activity/repository"
	"api_activity/requestdata"
	"errors"
	"fmt"
)

type TodoService interface {
	FindByID(id int64) (*entity.Todo, error)
	FindByGroupID(id int64) (*[]entity.Todo, error)
	FindAll() ([]entity.Todo, error)
	Create(input requestdata.CreateTodo) (*entity.Todo, error)
	UpdateByID(input requestdata.UpdateTodo, id int64) (*entity.Todo, error)
	DeleteByID(id int64) error
}

type todoService struct {
	agRepo   repository.ActivityGroupRepo
	todoRepo repository.TodoRepository
}

func NewTodo(agRepo repository.ActivityGroupRepo, todoRepo repository.TodoRepository) TodoService {
	return &todoService{agRepo, todoRepo}
}

func (s *todoService) FindByID(id int64) (*entity.Todo, error) {
	res, err := s.todoRepo.FindByID(id)
	return &res, err
}
func (s *todoService) FindByGroupID(id int64) (*[]entity.Todo, error) {
	res, err := s.todoRepo.FindByGroupID(id)
	return &res, err
}
func (s *todoService) FindAll() ([]entity.Todo, error) {
	res, err := s.todoRepo.FindAll()
	return res, err
}
func (s *todoService) Create(input requestdata.CreateTodo) (*entity.Todo, error) {
	var err error
	_, err = s.agRepo.FindByID(input.ActivityGroupID)
	if err != nil {
		id := input.ActivityGroupID
		err = fmt.Errorf("activity with activity_group_id %v Not Found",id)
		return nil, err
	}
	data := entity.Todo{
		ActivityGroupID: input.ActivityGroupID,
		Priority:        input.Priority,
		IsActive:        true,
		Title:           input.Title,
	}
	res, err := s.todoRepo.Create(data)
	return &res, err
}
func (s *todoService) UpdateByID(input requestdata.UpdateTodo, id int64) (*entity.Todo, error) {
	var err error
	_, err = s.agRepo.FindByID(input.ActivityGroupID)
	if err != nil {
		err = fmt.Errorf("activity with activity_group_id %v Not Found",id)
		return nil, err
	}
	data := map[string]interface{}{
		"title":             input.Title,
		"is_active":         input.IsActive,
		"priority":          input.Priority,
		"activity_group_id": input.ActivityGroupID,
	}
	res, err := s.todoRepo.UpdateByID(data, id)
	return &res, err
}
func (s *todoService) DeleteByID(id int64) error {
	var err error
	_, err = s.FindByID(id)
	if err != nil {
		err = errors.New("Not Found")
	} else {
		err = s.todoRepo.DeleteByID(id)
	}
	return err
}
