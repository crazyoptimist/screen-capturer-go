package common

import (
	"github.com/gin-gonic/gin"
)

func RaiseHttpError(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, HttpError{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}

// Exporting this in order to use it in API docs
type HttpError struct {
	StatusCode int    `json:"statusCode" example:"400"`
	Message    string `json:"message" example:"Bad Request"`
}

func (e *HttpError) Error() string {
	return e.Message
}
