package main

// importing fyne
import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating new app
	a := app.New()
	// New window and title
	w := a.NewWindow("Random Number Generator")
	// Resize
	w.Resize(fyne.NewSize(400, 400))
	// Label
	label1 := canvas.NewText("Rand Num Gen", color.Black)
	label1.TextSize = 40
	// Button
	btn1 := widget.NewButton("Generate", func() {
		// Ui is ready
		// Now logic
		rand1 := rand.Intn(1000)
		label1.Text = fmt.Sprint(rand1) //conver int to string
		label1.Refresh()                // refresh and render UI
	})

	// Showing and setup content
	w.SetContent(
		// more than on widget so use container newVbox
		container.NewVBox(
			label1,
			btn1,
		),
	)
	//show and run
	w.ShowAndRun()
}