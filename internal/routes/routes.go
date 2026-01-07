package routes

import (
	"github.com/Sudhanshu-NITR/notification-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := r.Group("/api/v1")
	{
		handler := handlers.NewNotificationHandler()
		v1.POST("/notifications", handler.Create)
	}
}
