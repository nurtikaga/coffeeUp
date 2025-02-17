package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nurtikaga/coffeeUp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	auth := routes.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := routes.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.listCoffee)
			lists.POST("/:id", h.createCoffee)
			lists.PUT("/:id", h.updateCoffee)
			lists.DELETE("/:id", h.deleteCoffee)
		}

		inventory := api.Group("/inventory")
		{
			inventory.GET("/", h.listInventory)
			inventory.PUT("/:id", h.updateInventory)
		}
	}
	return routes
}
