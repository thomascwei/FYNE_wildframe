package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	// window title
	w := a.NewWindow("Multi-Line Entry")
	// resize window size
	w.Resize(fyne.NewSize(400, 400))
	// New Multi line Entry Widget
	multilineEntry := widget.NewMultiLineEntry()
	const (
		loremIpsum = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
	)
	multilineEntry.SetText(loremIpsum)
	multilineEntry.Wrapping = fyne.TextWrapWord
	//multilineEntry.Wrapping = fyne.TextWrapOff
	//multilineEntry.Wrapping = fyne.TextTruncate
	//multilineEntry.Wrapping = fyne.TextWrapBreak
	w.SetContent(multilineEntry)
	// show and run
	w.ShowAndRun()
}
