package settings

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	MinutesMorning = "Morning meditation"
	MinutesShort = "Short meditations"
	NumberOfShort = "Number of short meditations"
	MinutesEvening = "Evening meditation"
)

var a fyne.App

func init() {
	a = app.NewWithID("repassyl.tawa.preferences")
}

func settingsSelector(values []string, title string, init string, suffix string) *fyne.Container {
	selector := widget.NewSelect(values, func(selected string) {
		a.Preferences().SetString(title + ":", selected)
	})
	selector.SetSelected(a.Preferences().StringWithFallback(title, init))
	return container.New(layout.NewHBoxLayout(), widget.NewLabel(title), layout.NewSpacer(), selector, widget.NewLabel(suffix))
}

func settingsMinutes(title string, init string) *fyne.Container {
	values := []string{"1", "3", "5", "8", "10", "15", "20", "25", "30", "45", "60", "90", "120"}
	return settingsSelector(values, title, init, " minutes")
}

func settingsNumber(title string, init string) *fyne.Container {
	values := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}
	return settingsSelector(values, title, init, " times")
}

func Screen() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		settingsMinutes(MinutesMorning, "25"),
		settingsMinutes(MinutesShort, "5"),
		settingsNumber(NumberOfShort, "8"),
		settingsMinutes(MinutesEvening, "10"),
	)
}

func Get(label string) string {
	return a.Preferences().StringWithFallback(label, "5")
}

func GetMinutes(label string) time.Duration {
	i, _ := strconv.Atoi(label)
	return time.Duration(i) * time.Minute
}