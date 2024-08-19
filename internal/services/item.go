package services

import (
	"todo/internal/models"
	"todo/internal/repository"
)

type ItemService struct {
	repository *repository.Repository
}

func NewItemService(repository *repository.Repository) *ItemService {
	return &ItemService{repository: repository}
}

func (service *ItemService) Create(userId int, listId int, item models.Item) (int, error) {
	_, err := service.repository.List.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return service.repository.Item.Create(listId, item)
}

func (service *ItemService) GetAll(userId int, listId int) ([]models.Item, error) {
	return service.repository.Item.GetAll(userId, listId)
}

func (service *ItemService) GetById(userId int, itemId int) (models.Item, error) {
	return service.repository.Item.GetById(userId, itemId)
}

func (service *ItemService) UpdateById(userId int, itemId int, item models.UpdateItemInput) error {
	return service.repository.Item.UpdateById(userId, itemId, item)
}

func (service *ItemService) DeleteById(userId int, itemId int) error {
	return service.repository.Item.DeleteById(userId, itemId)
}
