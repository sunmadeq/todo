package services

import (
	models2 "todo/internal/models"
	"todo/internal/repository"
)

type Auth interface {
	CreateUser(user models2.UserSignUpInput) (int, error)
	GenerateToken(user models2.UserSignInInput) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list models2.List) (int, error)
	GetAll(userId int) ([]models2.List, error)
	GetById(userId int, listId int) (models2.List, error)
	UpdateById(userId int, listId int, list models2.UpdateListInput) error
	DeleteById(userId int, listId int) error
}

type Item interface {
	Create(userId int, listId int, item models2.Item) (int, error)
	GetAll(userId int, listId int) ([]models2.Item, error)
	GetById(userId int, itemId int) (models2.Item, error)
	UpdateById(userId int, itemId int, item models2.UpdateItemInput) error
	DeleteById(userId int, itemId int) error
}

type Service struct {
	Auth
	List
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository),
		List: NewListService(repository),
		Item: NewItemService(repository),
	}
}
