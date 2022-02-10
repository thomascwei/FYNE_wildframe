package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

func main() {
	// New app
	a := app.New()
	//New title and window
	w := a.NewWindow("Open file in FYNE")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// New Buttton
	btn := widget.NewButton("Open .txt files", func() {
		// Using dialogs to open files
		// first argument func(fyne.URIReadCloser, error)
		// 2nd is parent window in our case "w"
		// r for reader
		// _ is ignore error
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				// read files
				data, _ := ioutil.ReadAll(r)
				// reader will read file and store data
				// now result
				result := fyne.NewStaticResource("thomas", data)
				// lets display our data in label or entry
				entry := widget.NewMultiLineEntry()
				// string() function convert byte to string
				entry.SetText(string(result.StaticContent))
				entry.Wrapping = fyne.TextWrapBreak
				// Lets show and setup content
				// tile of our new window
				w1 := fyne.CurrentApp().NewWindow(result.StaticName) // title/thomas
				w1.SetContent(container.NewScroll(entry))
				w1.Resize(fyne.NewSize(400, 400))
				// show/display content
				w1.Show()
				// we are almost done
			}, w)
		// fiter to open .txt files only
		// array/slice of strings/extensions
		file_Dialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()
		// Show file selection dialog.
	})
	// lets show button in parent window
	w.SetContent(container.NewVBox(
		btn,
	))
	w.ShowAndRun()
}
