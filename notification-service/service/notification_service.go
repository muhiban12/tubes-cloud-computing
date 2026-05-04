package service

import (
	"errors"
	"unit_testku/notification-service/models"
	"unit_testku/notification-service/repository"
)

type NotificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(r repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: r}
}

func (s *NotificationService) CreateNotification(userID, msg, msgType string) error {
	if msg == "" {
		return errors.New("pesan tidak boleh kosong")
	}

	notif := &models.Notification{
		UserID:  userID,
		Message: msg,
		Type:    msgType,
		Status:  "SENT",
	}

	// Memanggil provider (lewat repository)
	err := s.repo.Send(notif)
	if err != nil {
		return err
	}

	return s.repo.SaveLog(notif)
}