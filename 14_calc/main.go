package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"strconv"
)

func main() {
	// Convert logic to Beautiful UI
	// New app
	a := app.New()
	//New title
	w := a.NewWindow("BMI Calc")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// label
	label := canvas.NewText("BMI Calc", color.Black)
	label.Alignment = fyne.TextAlignCenter
	label.TextSize = 20
	// for resutl
	result := canvas.NewText("", color.Black)
	result.Alignment = fyne.TextAlignCenter
	result.TextSize = 20
	// input height
	inputH := widget.NewEntry()
	inputH.SetPlaceHolder("Enter height in cm..")
	inputW := widget.NewEntry()
	inputW.SetPlaceHolder("Enter Weight in KG..")
	btn1 := widget.NewButton("Calc BMI", func() {
		h, _ := strconv.ParseFloat(inputH.Text, 64)
		w, _ := strconv.ParseFloat(inputW.Text, 64)
		result.Text = calculateBMI(h/100, w)
		result.Refresh()
	})
	// setup content
	w.SetContent(
		container.NewVBox(
			label,
			result,
			inputH,
			inputW,
			btn1,
		))
	w.ShowAndRun()
}

// converting into function
func calculateBMI(height, weight float64) string {
	// copy the above code and paste here
	// BMI formula BMI = w/ h ^2
	var BMI float64 = weight / math.Pow(height, 2) // math.Pow(base,power)
	// conditions
	// BMI <= 18.4 "You are underweight.")
	if BMI <= 18.4 {
		fmt.Println("You are underweight.")
		return "You are underweight."
	} else if BMI <= 24.9 { // BMI <= 24.9  "You are healthy.")
		fmt.Println("You are healthy.")
		return "You are healthy."
	} else if BMI <= 29.9 { // BMI <= 29.9  "You are over weight.")
		fmt.Println("You are over weight.")
		return "You are over weight."
	} else if BMI <= 34.9 { // BMI <= 34.9  "You are severely over weight.")
		fmt.Println("You are severely over weight.")
		return "You are severely over weight."
	} else if BMI <= 39.9 { // BMI <= 39.9  "You are obese.")
		fmt.Println("You are obese.")
		return "You are obese."
	} else { // "You are severely obese.")
		fmt.Println("You are severely obese.")
		return "You are severely obese."
	}
}
