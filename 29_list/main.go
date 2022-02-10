package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// window and title
	w := a.NewWindow("List View")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// list View
	// 3 arguments
	// item count, 3 itmes in list
	// widget, I want to use label widget
	// update data of widget

	data := make([]string, 10)
	for i := range data {
		data[i] = fmt.Sprintf("Test Item %d", i)
	}
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List000")
	hbox := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.DocumentIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}

	w.SetContent(container.NewHSplit(list, container.NewCenter(hbox)))
	w.ShowAndRun()
}
