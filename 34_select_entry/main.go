package main
// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("Select entry widget, drop down")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets show our selected entry in label
	label1 := widget.NewLabel("...")
	// dropdown/ select entry
	//[]string{} all our option goes in slice
	// s is the variable to get the selected value
	dd := widget.NewSelect(
		[]string{"city1-dehli", "London", "islamabad", "kabul"},
		func(s string) {
			fmt.Printf("I selected %s to live forever..", s)
			label1.Text = s
			label1.Refresh()
		})
	// more than one widget. so use container
	c := container.NewVBox(dd, label1)
	w.SetContent(c)
	//show and run
	w.ShowAndRun()
}