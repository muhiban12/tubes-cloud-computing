package repository

import "unit_testku/notification-service/models"

type NotificationRepository interface {
	Send(notif *models.Notification) error
	SaveLog(notif *models.Notification) error
}