/*
* 7. Text and Change Color
* 8. image
* 9. Circle
* 10. rectangle
* 11. Canvas Line
* 12. ICON
* 13. Linear, Radial, Horizontal and Vertical Gradient
 */
package main

// import Fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	// Create new app
	a := app.New()

	// Create New Window and Title
	w := a.NewWindow("Canvas Tutorial")
	// Resize window
	w.Resize(fyne.NewSize(400, 400))
	// Our Text widget in Canvas
	//first value for label/cation
	// 2nd for color value
	// Now creating Color
	colorx := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	// R is Red  and Range is zero to 255
	// B is blue and G is Green
	// A is for alpha for opacity
	textX := canvas.NewText("Here is my blue text", colorx)

	rect := canvas.NewRectangle(color.Black)
	//fill color
	rect.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	rect.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	rect.StrokeWidth = 5.0

	circle1 := canvas.NewCircle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	circle1.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	circle1.StrokeWidth = 10 // Stroke width

	// creating line
	lineX := canvas.NewLine(color.Black)
	lineX.StrokeColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	lineX.StrokeWidth = 12.0

	// Creating Icon Widget
	IconX := widget.NewIcon(theme.CancelIcon())

	// Radial Gradient
	//gradient1 := canvas.NewRadialGradient(color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}) // 45 degree angle
	// Linear Gradient
	//gradient2 := canvas.NewLinearGradient(color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}, 270) // 45 degree angle
	//Horizontal Gradient
	//gradient3 := canvas.NewHorizontalGradient(color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	// Vertical Gradient
	gradient4 := canvas.NewVerticalGradient(color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}) // 45 degree angle
	gradient4.Resize(fyne.NewSize(400, 200))
	// setup content
	//w.SetContent(gradient4)

	// Now Image widget
	img := canvas.NewImageFromFile("pictures/gopher.png")
	//w.SetContent(img)
	// setup content to show
	w.SetContent(
		container.NewVSplit(
			container.NewHSplit(
				//textX,
				container.NewVSplit(textX, lineX),
				container.NewVSplit(rect, circle1)),

			//lineX,
			container.NewHSplit(
				container.NewHSplit(IconX, gradient4),
				img),
		),
	)
	w.ShowAndRun() // Show and run App
}
