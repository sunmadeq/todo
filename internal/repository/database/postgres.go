package database

import (
	"database/sql"
	"fmt"
)

const (
	UsersTable      = "users"
	ListsTable      = "lists"
	ItemsTable      = "items"
	UsersListsTable = "users_lists"
	ListsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Mode     string
}

func NewPostgresDatabase(config Config) (*sql.DB, error) {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Database, config.Mode,
	)

	database, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
