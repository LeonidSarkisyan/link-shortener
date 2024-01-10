package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"url-shotener-api/internal/models"
)

func (h *Handler) SaveURL(c *gin.Context) {
	var input models.ShortenInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := h.services.Shortening.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("url", input.RawURL).Str("alias", id).Msgf("Сохранён URL")
	c.JSON(200, map[string]any{
		"id": id,
	})
}

func (h *Handler) Redirect(c *gin.Context) {
	alias := c.Param("alias")

	url, err := h.services.Shortening.Get(alias)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("alias", alias).Str("url", url).Msgf("Совершён редирект")
	c.Redirect(http.StatusFound, url)
}
