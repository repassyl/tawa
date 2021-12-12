package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func settingsSelector(values []string, title string, init string, suffix string) *fyne.Container {
	a := app.NewWithID("repassyl.tawa.preferences")
	selector := widget.NewSelect(values, func(selected string) {
		a.Preferences().SetString(title, selected)
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

func settingsScreen() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		settingsMinutes("Morning meditation:", "25"),
		settingsMinutes("Short meditations:", "5"),
		settingsNumber("Number of short meditations:", "8"),
		settingsMinutes("Evening meditation:", "10"),
	)
}
