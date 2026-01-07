package handlers

import (
	"net/http"

	"github.com/Sudhanshu-NITR/notification-service/internal/models"
	"github.com/Sudhanshu-NITR/notification-service/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationHandler struct{}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

func (h *NotificationHandler) Create(c *gin.Context) {
	var req models.NotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: persist Notification
	// TODO: enqueue jobs

	notificationID := uuid.New().String()

	response.Success(c, http.StatusAccepted, models.NotificationResponse{
		NotificationID: notificationID,
		Status:         "accepted",
	})

}
