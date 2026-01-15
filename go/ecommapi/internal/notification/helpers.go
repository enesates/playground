package notif

func CreateEventNotif(uid string, title string, msg string) error {
	_, err := CreateNotification(uid, title, msg)
	return err
}
