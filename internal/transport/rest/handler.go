package rest

import (
	"github.com/co1seam/taskify/internal/services"
	"github.com/gin-gonic/gin"
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
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}

	api := router.Group("/api") 
	{
		subscription := api.Group("/subscription")
		{
			subscription.POST("/", h.Subscribe)
			subscription.GET("/", h.GetAllSubscriptions)
			subscription.GET("/:id", h.GetSubscriptionById)
			subscription.PUT("/:id", h.UpdateSubscription)
			subscription.DELETE("/:id", h.Unsubscribe)
			
			notification := subscription.Group("/:id/notification")
			{
				notification.POST("/", h.SendNotification)
			}
		}
	}

	return router
}