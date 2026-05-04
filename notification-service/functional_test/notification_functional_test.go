package functional_test

import (
	"testing"
	"unit_testku/notification-service/models"
	"unit_testku/notification-service/repository"
	"unit_testku/notification-service/service"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNotification_Functional(t *testing.T) {
	// 1. Setup Database SQLite khusus notifikasi
	db, _ := gorm.Open(sqlite.Open("notification_test.db"), &gorm.Config{})
	db.AutoMigrate(&models.Notification{})
	db.Exec("DELETE FROM notifications") // Bersihkan data lama

	// 2. Setup Repository dan Service
	repo := repository.NewNotificationRepository(db)
	notifService := service.NewNotificationService(repo)

	// 3. Jalankan Aksi: Kirim Notifikasi
	err := notifService.CreateNotification("user-007", "Paketmu sedang dikirim kurir!", "SMS")

	// 4. Verifikasi: Cek apakah data tersimpan di DB
	var savedNotif models.Notification
	db.First(&savedNotif, "user_id = ?", "user-007")

	// 5. Assert
	assert.NoError(t, err)
	assert.Equal(t, "SMS", savedNotif.Type)
	assert.Equal(t, "Paketmu sedang dikirim kurir!", savedNotif.Message)
}