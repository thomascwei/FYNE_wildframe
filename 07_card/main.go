package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("cards")
	w.Resize(fyne.NewSize(400, 400))

	card1 := widget.NewCard("Book a table", "Which time suits?",
		widget.NewRadioGroup([]string{"6:30pm", "7:00pm", "7:45pm"}, func(string) {}))
	card2 := widget.NewCard("With media", "No content, with image", nil)
	card2.Image = canvas.NewImageFromResource(theme.FyneLogo())
	card3 := widget.NewCard("Title 3", "Another card", widget.NewLabel("Content"))

	w.SetContent(container.NewGridWithColumns(2, container.NewVBox(card1, card3),
		container.NewVBox(card2)))
	w.ShowAndRun()
}
