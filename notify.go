package Notification

import (
	"log"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

func Notify(appName, summary, description, icon string, expireTime int) {
	conn, err := dbus.SessionBusPrivate()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	if err = conn.Auth(nil); err != nil {
		log.Println(err)
		return
	}

	if err = conn.Hello(); err != nil {
		log.Println(err)
		return
	}

	n := notify.Notification{
		AppName:       appName,
		ReplacesID:    uint32(0),
		Summary:       summary,
		Body:          description,
		ExpireTimeout: time.Second * time.Duration(expireTime),
		AppIcon:       icon,
	}

	notify.SendNotification(conn, n)
}
