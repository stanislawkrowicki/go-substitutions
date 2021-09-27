package notifications

import (
	"github.com/go-toast/toast"
	"path/filepath"
)

const (
	Image   = "assets/images/logo.png"
	AppName = "go-substitutions"
)

func Show(title, changes string) error {
	image, _ := filepath.Abs(Image)
	notification := toast.Notification{
		AppID:   AppName,
		Title:   title,
		Message: changes,
		Icon:    image,
	}

	err := notification.Push()
	if err != nil {
		return err
	}

	return nil
}
