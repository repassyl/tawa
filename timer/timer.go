package timer

import (
	"time"
	"fyne.io/fyne/v2/widget"
)

func updateText(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}

func Start() *widget.Label {
	clock := widget.NewLabel("")
	updateText(clock)
	go func() {
		t := time.NewTicker(500 * time.Millisecond) 
		for range t.C {
			updateText(clock)
		}
	}()
	return clock
}
