package Notification

import (
	"testing"
)

func TestNotify(t *testing.T) {
	Notify("NotifyTest", "summary", "description", "", 30)
}
