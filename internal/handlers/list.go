package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/models"
)

func (handler *Handler) createList(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	var input models.List
	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.List.Create(userId, input)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAllListsResponse struct {
	Data []models.List `json:"data"`
}

func (handler *Handler) getAllLists(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	lists, err := handler.service.List.GetAll(userId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, GetAllListsResponse{
		Data: lists,
	})
}

func (handler *Handler) getListById(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	list, err := handler.service.List.GetById(userId, listId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, list)
}

type UpdateListResponse struct {
	Status string `json:"status"`
}

func (handler *Handler) updateList(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateListInput
	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.List.UpdateById(userId, listId, input)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, UpdateListResponse{
		Status: "ok",
	})
}

type DeleteListResponse struct {
	Status string `json:"status"`
}

func (handler *Handler) deleteList(context *gin.Context) {
	userId, err := handler.userId(context)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.List.DeleteById(userId, listId)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, DeleteListResponse{
		Status: "ok",
	})
}
