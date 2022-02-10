package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

func main() {
	// New app
	a := app.New()
	// New title
	w := a.NewWindow("file Handling - Open Images")
	//resize
	w.Resize(fyne.NewSize(400, 400))
	btn := widget.NewButton("Open .jpg & .Png", func() {
		// dialog for opening files
		// 2 arguments
		fileDialog := dialog.NewFileOpen(
			// _ to ignore error
			func(uc fyne.URIReadCloser, _ error) {
				// reader to read data
				data, _ := ioutil.ReadAll(uc)
				// static resource
				// 2 arguments
				// first is file name (string)
				// second is data from reader
				res := fyne.NewStaticResource(uc.URI().Name(), data)
				// Now image widget to display our image
				img := canvas.NewImageFromResource(res)
				// setup new window for image and set content
				w1 := fyne.CurrentApp().NewWindow(uc.URI().Name())
				w1.SetContent(img)
				// resize window
				w1.Resize(fyne.NewSize(400, 400))
				w1.Show() // display our image
			}, w)
		// filtering files
		fileDialog.SetFilter(
			// filter jpg and png
			// ignore rest of the files
			storage.NewExtensionFileFilter([]string{".png", ".jpg"}))
		fileDialog.Show()
		// we are done 🙂
	})
	// display button in parent window
	w.SetContent(btn)
	w.ShowAndRun()
}
