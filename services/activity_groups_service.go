package services

import (
	"api_activity/entity"
	"api_activity/repository"
	"api_activity/requestdata"
)

type ActivityGroupsService interface {
	FindByID(id int64) (entity.ActivityGroups, error)
	FindByAll() ([]entity.ActivityGroups, error)
	Create(input requestdata.CreateActicityGroups) (entity.ActivityGroups, error)
	UpdateByID(input requestdata.UpdateActivityGroups, id int64) (entity.ActivityGroups, error)
	DeleteByID(id int64) error
}

type activityGroupService struct {
	repo repository.ActivityGroupRepo
}

func New(repo repository.ActivityGroupRepo) ActivityGroupsService {
	return &activityGroupService{repo}
}

func (s *activityGroupService) FindByID(id int64) (entity.ActivityGroups, error) {
	res, err := s.repo.FindByID(id)
	return res, err
}
func (s *activityGroupService) FindByAll() ([]entity.ActivityGroups, error) {
	res, err := s.repo.FindAll()
	return res, err
}
func (s *activityGroupService) Create(input requestdata.CreateActicityGroups) (entity.ActivityGroups, error) {
	data := entity.ActivityGroups{
		Email: input.Email,
	}
	res, err := s.repo.Create(data)
	return res, err
}
func (s *activityGroupService) UpdateByID(input requestdata.UpdateActivityGroups, id int64) (entity.ActivityGroups, error) {
	data := map[string]interface{}{
		"tittle": input.Tittle,
	}
	res, err := s.repo.UpdateByID(data, id)
	return res, err
}
func (s *activityGroupService) DeleteByID(id int64) error {
	err := s.repo.DeleteByID(id)
	return err
}
