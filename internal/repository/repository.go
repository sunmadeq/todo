package repository

import (
	"database/sql"
	models2 "todo/internal/models"
	"todo/internal/repository/postgres"
)

type Auth interface {
	CreateUser(user models2.UserSignUpInput) (int, error)
	GetUser(email string) (models2.User, error)
}

type List interface {
	Create(userId int, list models2.List) (int, error)
	GetAll(userId int) ([]models2.List, error)
	GetById(userId int, listId int) (models2.List, error)
	UpdateById(userId int, listId int, input models2.UpdateListInput) error
	DeleteById(userId int, listId int) error
}

type Item interface {
	Create(listId int, item models2.Item) (int, error)
	GetAll(userId int, listId int) ([]models2.Item, error)
	GetById(userId int, itemId int) (models2.Item, error)
	UpdateById(userId int, itemId int, item models2.UpdateItemInput) error
	DeleteById(userId int, itemId int) error
}

type Repository struct {
	Auth
	List
	Item
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		Auth: repository.NewAuthPostgres(database),
		List: repository.NewListPostgres(database),
		Item: repository.NewItemPostgres(database),
	}
}
