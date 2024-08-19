package repository

import (
	"database/sql"
	"fmt"
	"todo/internal/models"
	"todo/internal/repository/database"
)

type AuthPostgres struct {
	database *sql.DB
}

func NewAuthPostgres(database *sql.DB) *AuthPostgres {
	return &AuthPostgres{database: database}
}

func (repository *AuthPostgres) CreateUser(user models.UserSignUpInput) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		database.UsersTable,
	)

	row := repository.database.QueryRow(
		query,
		user.Name, user.Email, user.Password,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *AuthPostgres) GetUser(email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE email=$1",
		database.UsersTable,
	)

	row := repository.database.QueryRow(
		query,
		email,
	)

	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}
