package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/repassyl/tawa/timer"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Looking at Tawa")
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Welcome", theme.HomeIcon(), widget.NewLabel("TODO: links to the teachings")),
		container.NewTabItemWithIcon("Morning", theme.MediaPlayIcon(), timer.Start()),
		container.NewTabItemWithIcon("Looking at Tawa", theme.MediaFastForwardIcon(), timer.Start()),
		container.NewTabItemWithIcon("Evening", theme.MediaPlayIcon(), timer.Start()),
		container.NewTabItemWithIcon("Statistics", theme.HistoryIcon(), widget.NewLabel("TODO")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settingsScreen()),
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