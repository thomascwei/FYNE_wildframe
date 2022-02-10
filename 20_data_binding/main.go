package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	// New window and title
	w := a.NewWindow("Slider & data binding")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// first slider widget
	// 3 arguments
	// min value, max value, and data source

	f := 20.0 // any float value
	data := binding.BindFloat(&f)
	slider1 := widget.NewSliderWithData(
		0, 100, data)
	// Slider part is done
	// Now lets attach slider value with label
	label1 := widget.NewLabelWithData(
		binding.FloatToString(data),
	)

	// we are done :)
	// Lets show and setup content
	w.SetContent(
		container.NewVBox(
			label1,
			slider1,
		),
	)
	// Show data
	w.ShowAndRun()
}
