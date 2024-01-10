package handler

import (
	"github.com/gin-gonic/gin"
	"url-shotener-api/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.currentUser)
	{
		short := api.Group("/short")
		{
			short.POST("/", h.SaveURL)
			short.GET("/:alias", h.Redirect)
		}
	}

	return router
}
