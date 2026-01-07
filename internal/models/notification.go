package models

type NotificationRequest struct {
	Recipient Recipient            `json:"recipient" binding:"required"`
	Channels  []string             `json:"channels" binding:"required,min=1"`
	Message   NotificationBody     `json:"message" binding:"required"`
	Options   *NotificationOptions `json:"options,omitempty"`
}

type Recipient struct {
	UserID    string `json:"user_id,omitempty"`
	Email     string `json:"email,omitempty" binding:"omitempty,email"`
	Phone     string `json:"phone,omitempty"`
	PushToken string `json:"push_token,omitempty"`
}

type NotificationBody struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type NotificationOptions struct {
	Priority       string `json:"priority,omitempty"`
	ScheduleAt     string `json:"schedule_at,omitempty"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

type NotificationResponse struct {
	NotificationID string `json:"notification_id"`
	Status         string `json:"status"`
}
