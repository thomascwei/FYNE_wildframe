package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Colored Button")
	//resize window
	w.Resize(fyne.NewSize(400, 400))
	// Let show out button on screen
	disabledButton := widget.NewButton("Disabled", func() {})
	disabledButton.Disable()
	w.SetContent(
		container.NewVBox(
			disabledButton,
			red_button(),
			red_button(),
			red_button(),
			red_button(),
			// let show / display buttons
			blue_button(),
			green_button(),
			img_button(),
		),
	)
	w.ShowAndRun() // show and run app
}

// first colored button
func red_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewCircle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// Blue button
// copy previous button code and change it
func blue_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 05, G: 0, B: 255, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// Green button
// copy and paste previous button
func green_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewLine(
		color.NRGBA{R: 0, G: 255, B: 0, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// button with image as background
func img_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// img button color
	img := canvas.NewImageFromFile("pictures/dice/dice5.jpg")
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		img,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}
