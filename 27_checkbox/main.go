package main

// import fyne
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("check Box tutorial")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// our first widget
	label := canvas.NewText("Your choice is ...", color.Black)
	// check box
	// first argument is name(string),
	// 2nd is function

	chk_male := widget.NewCheck("male", func(b bool) {
		// change name of label
		if b == true {
			label.Text = "male"
			// refresh
			label.Refresh()
		} else {
			label.Text = "deseleted"
			// refresh
			label.Refresh()
		}
	})

	// create new one for female
	chk_female := widget.NewCheck("female", func(b bool) {
		// change name of label
		if b == true {
			label.Text = "female"
			label.Hide()
			// refresh
			label.Refresh()
		} else {
			label.Text = "deseleted"
			label.Show()
			// refresh
			label.Refresh()
		}
	})

	w.SetContent(
		// here we put all our widgets
		// we are done...
		// lets show content on screen
		container.NewVBox(
			label,
			chk_male,
			chk_female,
		),
	)
	w.ShowAndRun()
}