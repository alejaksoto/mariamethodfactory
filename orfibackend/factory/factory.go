package factory

import (
	"errors"

	"notification/notification"
)

func GetNotificationChannel(channelType string) (notification.Notification, error) {
	switch channelType {
	case "email":
		return &notification.EmailNotification{}, nil
	case "whatsapp":
		return &notification.WhatsAppNotification{}, nil
	default:
		return nil, errors.New("canal de notificación no soportado")
	}
}
