package notif

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
)

func FetchNotification(nid string) (*db.Notification, error) {
	notification := db.Notification{
		ID: nid,
	}

	if err := db.GormDB.Preload("User").First(&notification).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func FetchNotifications(uid string) ([]db.Notification, error) {
	notifications := []db.Notification{}

	if err := db.GormDB.
		Preload("User").
		Where("user_id = ?", uid).
		Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

func CreateNotification(uid string, title string, msg string) (*db.Notification, error) {
	notification := db.Notification{
		ID:      utils.GetUUID(),
		UserID:  uid,
		Title:   title,
		Message: msg,
		IsRead:  false,
	}

	if err := db.GormDB.Create(&notification).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func UpdateNotificationStatus(id string) error {
	if err := db.GormDB.
		Model(&db.Notification{}).
		Where("id = ?", id).
		UpdateColumn("is_read", true).
		Error; err != nil {
		return err
	}

	return nil
}
