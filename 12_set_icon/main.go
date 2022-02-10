package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// New app
	a := app.New()
	// New Title and Window
	w := a.NewWindow("Load images from internet")
	// loading images/assets/icons from computer
	r, _ := fyne.LoadResourceFromPath("pictures/dice/dice5.jpg")
	w.SetIcon(r)

	// Images/assets from internet
	r, _ = fyne.LoadResourceFromURLString("https://picsum.photos/200")
	//r, _ := fyne.LoadResourceFromURLString("http://placekitten.com/g/200/300")
	// CAT API Free -- http://placekitten.com/g/200/300
	// New Random image URL https://picsum.photos/200
	img := canvas.NewImageFromResource(r)

	// Resize
	w.Resize(fyne.NewSize(400, 400))

	// Show and Setup content
	w.SetContent(img)
	w.ShowAndRun()
}