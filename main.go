package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/repassyl/tawa/timer"
	"github.com/repassyl/tawa/settings"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Looking at Tawa")
	myWindow.SetIcon(theme.HistoryIcon())
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Welcome", theme.HomeIcon(), widget.NewLabel("TODO: links to the teachings")),
		container.NewTabItemWithIcon("Morning", theme.MediaPlayIcon(), timer.Init(settings.MinutesMorning)),
		container.NewTabItemWithIcon("Looking at Tawa", theme.MediaFastForwardIcon(), timer.Init(settings.MinutesShort)),
		container.NewTabItemWithIcon("Evening", theme.MediaPlayIcon(), timer.Init(settings.MinutesEvening)),
		container.NewTabItemWithIcon("Statistics", theme.HistoryIcon(), widget.NewLabel("TODO")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settings.Screen()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.SetCloseIntercept(func() {
		myApp.Quit()
	})
	myWindow.SetContent(tabs)
	// TODO: Android
	// myWindow.SetFullScreen(true)
	myWindow.Resize(fyne.Size{Height: 600, Width: 800})
	myWindow.ShowAndRun()
}
