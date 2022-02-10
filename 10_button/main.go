package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// set theme
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("theme")
	w.Resize(fyne.Size{
		Width:  400,
		Height: 400,
	})

	label := widget.NewLabel("fyne theme")
	btn1 := widget.NewButton("light", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	btn2 := widget.NewButton("dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	btn3 := widget.NewButton("exit", func() {
		a.Quit()
	})
	w.SetContent(
		container.NewVBox(
			label,
			btn1,
			btn2,
			btn3,
		))
	w.ShowAndRun()
}
