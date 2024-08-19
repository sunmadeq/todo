package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/internal/models"
)

func (handler *Handler) signUp(context *gin.Context) {
	var input models.UserSignUpInput

	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.CreateUser(input)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) signIn(context *gin.Context) {
	var input models.UserSignInInput

	if err := context.BindJSON(&input); err != nil {
		NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	token, err := handler.service.GenerateToken(input)
	if err != nil {
		NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
