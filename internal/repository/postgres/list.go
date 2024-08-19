package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"todo/internal/models"
	"todo/internal/repository/database"
)

type ListPostgres struct {
	database *sql.DB
}

func NewListPostgres(database *sql.DB) *ListPostgres {
	return &ListPostgres{database: database}
}

func (repository *ListPostgres) Create(userId int, list models.List) (int, error) {
	tx, err := repository.database.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	var query string

	query = fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		database.ListsTable,
	)

	row := tx.QueryRow(query, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id) VALUES ($1, $2)",
		database.UsersListsTable,
	)

	_, err = tx.Exec(query, userId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (repository *ListPostgres) GetAll(userId int) ([]models.List, error) {
	var lists []models.List
	var query string

	query = fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		database.ListsTable,
		database.UsersListsTable,
	)

	rows, err := repository.database.Query(query, userId)
	if err != nil {
		return lists, err
	}

	for rows.Next() {
		var list models.List

		err := rows.Scan(&list.Id, &list.Title, &list.Description)
		if err != nil {
			return lists, err
		}

		lists = append(lists, list)
	}

	if err = rows.Err(); err != nil {
		return lists, err
	}

	return lists, nil
}

func (repository *ListPostgres) GetById(userId int, listId int) (models.List, error) {
	var list models.List
	var query string

	query = fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		database.ListsTable,
		database.UsersListsTable,
	)

	row := repository.database.QueryRow(query, userId, listId)
	err := row.Scan(&list.Id, &list.Title, &list.Description)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (repository *ListPostgres) UpdateById(userId int, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		database.ListsTable,
		setQuery,
		database.UsersListsTable,
		argId,
		argId+1,
	)

	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := repository.database.Exec(query, args...)

	return err
}

func (repository *ListPostgres) DeleteById(userId int, listId int) error {
	var query string

	query = fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		database.ListsTable,
		database.UsersListsTable,
	)

	_, err := repository.database.Exec(query, userId, listId)

	return err
}
