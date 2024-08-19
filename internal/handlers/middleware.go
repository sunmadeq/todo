package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (handler *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(context, http.StatusUnauthorized, "No authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(context, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userId, err := handler.service.Auth.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	context.Set("userId", userId)
}

func (handler *Handler) userId(context *gin.Context) (int, error) {
	id, ok := context.Get("userId")
	if !ok {
		NewErrorResponse(context, http.StatusInternalServerError, "User id not found.")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(context, http.StatusInternalServerError, "User id is of invalid type.")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
