package services

import (
	"todo/internal/models"
	"todo/internal/repository"
)

type ListService struct {
	repository *repository.Repository
}

func NewListService(repository *repository.Repository) *ListService {
	return &ListService{repository: repository}
}

func (service *ListService) Create(userId int, list models.List) (int, error) {
	return service.repository.List.Create(userId, list)
}

func (service *ListService) GetAll(userId int) ([]models.List, error) {
	return service.repository.List.GetAll(userId)
}

func (service *ListService) GetById(userId int, listId int) (models.List, error) {
	return service.repository.List.GetById(userId, listId)
}

func (service *ListService) DeleteById(userId int, listId int) error {
	return service.repository.List.DeleteById(userId, listId)
}

func (service *ListService) UpdateById(userId int, listId int, list models.UpdateListInput) error {
	return service.repository.List.UpdateById(userId, listId, list)
}
