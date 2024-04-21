package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func CreateError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, &ResponseMessage{Message: err.Error()})
}

func CreateSuccess(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, &ResponseMessage{Message: message})
}
