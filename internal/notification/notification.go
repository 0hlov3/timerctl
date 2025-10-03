package notification

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func Notify(name string) {
	fmt.Print("\a")
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	var icon []byte
	err = beeep.Notify("timerctl", name, icon)
	if err != nil {
		fmt.Println("\rcould not send Notification")
	}
}
