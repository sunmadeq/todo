package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"todo/internal/models"
	"todo/internal/repository/database"
)

type ItemPostgres struct {
	database *sql.DB
}

func NewItemPostgres(database *sql.DB) *ItemPostgres {
	return &ItemPostgres{database: database}
}

func (repository *ItemPostgres) Create(listId int, item models.Item) (int, error) {
	tx, err := repository.database.Begin()
	if err != nil {
		return 0, err
	}

	var row *sql.Row
	var query string
	var id int

	query = fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		database.ItemsTable,
	)

	row = tx.QueryRow(query, item.Title, item.Description)
	err = row.Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (list_id, item_id) VALUES ($1, $2)",
		database.ListsItemsTable,
	)

	_, err = tx.Exec(query, listId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (repository *ItemPostgres) GetAll(userId int, listId int) ([]models.Item, error) {
	var items []models.Item

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
		  INNER JOIN %s li on li.item_id = ti.id 
		  INNER JOIN %s ul on ul.list_id = li.list_id
 		  WHERE ul.user_id = $1 AND li.list_id = $2`,
		database.ItemsTable,
		database.ListsItemsTable,
		database.UsersListsTable,
	)

	rows, err := repository.database.Query(query, userId, listId)
	if err != nil {
		return items, err
	}

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (repository *ItemPostgres) GetById(userId int, itemId int) (models.Item, error) {
	var item models.Item

	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
		INNER JOIN %s li on li.item_id = ti.id
		INNER JOIN %s ul on ul.list_id = li.list_id 
		WHERE ul.user_id = $1 AND ti.id = $2`,
		database.ItemsTable, database.ListsItemsTable, database.UsersListsTable,
	)

	row := repository.database.QueryRow(query, userId, itemId)
	err := row.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (repository *ItemPostgres) UpdateById(userId int, itemId int, item models.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *item.Title)
		argId++
	}

	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *item.Description)
		argId++
	}

	if item.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *item.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s ti SET %s FROM %s li, %s ul
		WHERE ti.id = li.item_id 
		AND li.list_id = ul.list_id 
		AND ul.user_id = $%d 
		AND ti.id = $%d`,
		database.ItemsTable,
		setQuery,
		database.ListsItemsTable,
		database.UsersListsTable,
		argId,
		argId+1,
	)

	args = append(args, userId, itemId)

	_, err := repository.database.Exec(query, args...)

	return err
}

func (repository *ItemPostgres) DeleteById(userId int, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ti USING %s li, %s ul 
         WHERE ti.id = li.item_id 
		 AND li.list_id = ul.list_id 
		 AND ul.user_id = $1 
		 AND ti.id = $2`,
		database.ItemsTable, database.ListsItemsTable, database.UsersListsTable,
	)

	_, err := repository.database.Exec(query, userId, itemId)

	return err
}
