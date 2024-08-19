package handlers

import (
	"github.com/gin-gonic/gin"
	"todo/internal/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api", handler.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", handler.createList)
			lists.GET("/", handler.getAllLists)
			lists.GET("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", handler.createItem)
				items.GET("/", handler.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", handler.getItemById)
			items.PUT("/:id", handler.updateItem)
			items.DELETE("/:id", handler.deleteItem)
		}
	}

	return router
}
