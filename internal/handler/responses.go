package handler

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Detail string `json:"detail"`
}

func NewErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{err.Error()})
	return
}
