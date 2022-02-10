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
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("Tool bar - music player ui")
	// resize main window
	w.Resize(fyne.NewSize(400, 400))
	// new widget---- Tool bar
	toolbar := widget.NewToolbar(
		//toolbar items
		widget.NewToolbarAction(
			// first argument is icon
			// 2nd is action/func(){}
			theme.MediaPlayIcon(), func() {
				fmt.Println("Play music...")
			},
		),
		// copy paste for other items
		// pause
		widget.NewToolbarAction(
			// first argument is icon
			// 2nd is action/func(){}
			theme.MediaPauseIcon(), func() {
				fmt.Println("Pause music...")
			},
		),
		widget.NewToolbarAction(
			// first argument is icon
			// 2nd is action/func(){}
			theme.MediaStopIcon(), func() {
				fmt.Println("Stop music...")
			},
		),
		// spacer widget
		widget.NewToolbarSpacer(),
		// new item for support
		/// copy paste
		widget.NewToolbarAction(
			// first argument is icon
			// 2nd is action/func(){}
			theme.HelpIcon(), func() {
				fmt.Println("support music...")
			},
		),
	)
	// put everything in a container
	//c := container.NewVBox(toolbar)
	// let's change container
	c := container.NewBorder(
		// 5 elements / arguments
		// top, right,left,bottom, center
		toolbar, nil, nil, nil, widget.NewLabel("Content here"),
	)
	w.SetContent(c)
	w.ShowAndRun()
}
