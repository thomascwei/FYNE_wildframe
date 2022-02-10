package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	// New title
	w := a.NewWindow("Accordian - fyne ")

	// resize windows size
	w.Resize(fyne.NewSize(400, 400))
	// Let create items for accoridan
	// first argument is title, 2nd is description/details
	item1 := widget.NewAccordionItem("A",
		widget.NewLabel("A for Apple"))

	item2 := widget.NewAccordionItem("B",
		widget.NewLabel("B for Ball"))

	item3 := widget.NewAccordionItem("C",
		widget.NewLabel("C for Cat"))
	item4 := widget.NewAccordionItem("D",
		widget.NewButton("wtf", func() {

		}))
	ac := widget.NewAccordion(item1, item2, item3, item4)

	w.SetContent(ac)

	w.ShowAndRun()
}
