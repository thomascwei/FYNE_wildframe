package main

// import fyne
import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// New app
	a := app.New()
	// New window and Title
	w := a.NewWindow("Dicee Game")

	// resize Window
	w.Resize(fyne.NewSize(400, 400))
	// image of Dice
	img := canvas.NewImageFromFile("pictures/dice/dice6.jpg")
	img.FillMode = canvas.ImageFillOriginal

	// button
	btn1 := widget.NewButton("Play", func() {
		// UI is finished.. Now Logic part
		rand := rand.Intn(6) + 1 // ignore zero 0+1=1
		img.File = fmt.Sprintf("pictures/dice/dice%d.jpg", rand)
		img.Refresh()
	})
	btn1.Resize(fyne.NewSize(20,20))

	// Setup content and finish UI
	w.SetContent(
		// NewVBox.. More than on Widgets
		container.NewVBox(
			img,
			btn1,
		),
	)
	// show and run app
	w.ShowAndRun()
}