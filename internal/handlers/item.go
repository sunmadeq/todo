package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/models"
)

func (handler *Handler) createItem(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, "Invalid list id.")
		return
	}

	var input models.Item

	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Item.Create(userId, listId, input)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getAllItems(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, "Invalid list id.")
		return
	}

	items, err := handler.service.Item.GetAll(userId, listId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, items)
}

func (handler *Handler) getItemById(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, "Invalid item id.")
		return
	}

	item, err := handler.service.Item.GetById(userId, itemId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, item)
}

func (handler *Handler) updateItem(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, "Invalid item id.")
		return
	}

	var input models.UpdateItemInput
	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.service.Item.UpdateById(userId, id, input); err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (handler *Handler) deleteItem(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, "Invalid item id.")
		return
	}

	err = handler.service.Item.DeleteById(userId, itemId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
