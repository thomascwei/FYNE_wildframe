package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var DataType string
	var FilePath string
	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("Select entry widget, drop down")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets show our selected entry in label
	label1 := widget.NewLabel("data type : ")
	label2 := widget.NewLabel("file : ")
	label3 := widget.NewLabel("ready???")

	// dropdown/ select entry
	//[]string{} all our option goes in slice
	// s is the variable to get the selected value
	dd := widget.NewSelect(
		[]string{"temperature", "battery_usage"},
		func(s string) {
			DataType = s
			label1.Text = fmt.Sprintf("data type : %s", s)
			label1.Refresh()
		})
	// more than one widget. so use container

	btn1 := widget.NewButton("Open a file", func() {
		dialog.ShowFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				fmt.Println(r.URI().Path())
				FilePath = r.URI().Path()
				entry := widget.NewMultiLineEntry()
				entry.SetText(r.URI().Path())
				entry.Wrapping = fyne.TextWrapBreak
				label2.Text = fmt.Sprintf("file : %s", FilePath)
				label2.Refresh()
			}, w)
	})

	btn2 := widget.NewButton("execute", func() {
		label3.Text = " GO!!! " + label1.Text + " " + label2.Text
		label3.Refresh()
	})

	c := container.NewVBox(dd, label1, btn1, label2, btn2, label3)
	w.SetContent(c)
	//show and run
	w.ShowAndRun()
}
