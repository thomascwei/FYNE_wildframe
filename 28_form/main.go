package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("New Form")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// label empty
	label := widget.NewLabel("")
	// form widget
	form := widget.NewForm(
		// new use form items
		// 2 arguments
		// label, widget
		widget.NewFormItem("Username", widget.NewEntry()),
		// password
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	// working on cancel and submit functions of form
	form.OnCancel = func() {
		label.Text = "Canceled"
		label.Refresh()
	}

	form.OnSubmit = func() {
		label.Text = "submitted"
		label.Refresh()
	}

	// we are almost done
	w.SetContent(
		container.NewVBox(
			form,
			label,
		),
	)
	w.ShowAndRun()
}