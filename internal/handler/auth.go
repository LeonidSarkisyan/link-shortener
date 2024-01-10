package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"url-shotener-api/internal/models"
	"url-shotener-api/pkg/utils"
)

func (h *Handler) SignUp(c *gin.Context) {
	var userInput models.UserInput

	if err := c.BindJSON(&userInput); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	userId, err := h.services.Auth.Register(userInput)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id": userId,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var userInput models.UserInput

	if err := c.BindJSON(&userInput); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := h.services.Auth.Login(userInput)
	if err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (h *Handler) currentUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, errors.New("пустой auth заголовок"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, errors.New("некорректный auth заголовок"))
		return
	}

	userId, err := utils.ExtractUserIDFromToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, errors.New("некорректный токен"))
		return
	}
	c.Set("userId", userId)
}
