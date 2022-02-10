package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	w := a.NewWindow("Save files...")
	// file handling tutorial
	// Resize
	w.Resize(fyne.NewSize(400, 400))
	// entry filename
	fileName := widget.NewEntry()
	fileName.PlaceHolder = "enter filename"
	// Entry to enter text
	entry := widget.NewMultiLineEntry()
	entry.PlaceHolder="write some content"
	// btn to save text
	btn := widget.NewButton("Save file as .txt", func() {
		// dialog
		// 2 arguments
		// one function
		// 2nd parent window
		fileDialog := dialog.NewFileSave(
			// data of entry
			// []byte() function is used to convert
			// string to bytes slice
			func(uc fyne.URIWriteCloser, _ error) {
				data := []byte(entry.Text)
				//_ to ignore error
				// Lets write data
				uc.Write(data)
			}, w) // w is parent window
		// File name(temporary)
		fileDialog.SetFileName(fileName.Text + ".txt")
		// Show and setup
		fileDialog.Show()
	})
	// show our two widgets on screen
	w.SetContent(
		container.NewVBox(
			fileName,
			entry,
			btn,
		),
	)
	w.ShowAndRun()
}
