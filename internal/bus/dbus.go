package bus

import (
	"fmt"
	"log"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/zeshi09/zeshifyd/internal/model"
)

type NotificationHandler struct {
	OnNotify func(n model.Notification) // callback - what we are doing with notification
}

func (h *NotificationHandler) Start() error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect DBus: %w", err)
	}

	// 1. register the name
	reply, err := conn.RequestName("org.freedesktop.Notifications", dbus.NameFlagDoNotQueue) // if name already taken to smth else, do not wait, return err
	if err != nil {
		return fmt.Errorf("failed to request name: %w", err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner { // you took then name
		return fmt.Errorf("name already taken")
	}

	// 2. export "Notify" method
	err = conn.Export(h, "/org/freedesktop/Notifications", "org.freedesktop.Notifications")
	if err != nil {
		return fmt.Errorf("failed to export interface: %w", err)
	}

	log.Println("DBus is ready, listening for notifications")

	select {}
}

// Notify method what call every notify-send
func (h *NotificationHandler) Notify(
	appName string,
	replacesId uint32,
	appIcon string,
	summary string,
	body string,
	actions []string,
	hints map[string]dbus.Variant,
	expireTimeout int32,
) (uint32, *dbus.Error) {
	n := model.Notification{
		Appname: appName,
		Summary: summary,
		Body:    body,
		Icon:    appIcon,
		Time:    time.Now(),
		Timeout: int(expireTimeout),
	}

	if h.OnNotify != nil {
		h.OnNotify(n)
	}

	return 0, nil
}

// useless methods
func (h *NotificationHandler) GetServerInformation() (string, string, string, string, *dbus.Error) {
	serverName := "zeshifyd"
	maintainer := "blackzeshi"
	version := "0.1"
	good := "boy"

	return serverName, maintainer, version, good, nil
}
