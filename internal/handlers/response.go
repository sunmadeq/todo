package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Message string `json:"message"`
}

func NewErrorResponse(context *gin.Context, status int, message string) {
	logrus.Error(message)
	context.AbortWithStatusJSON(status, Response{
		Message: message,
	})
}
