/*
*	4. button
*	5. checkbox
*	6. Hyperlink
 */
package main

// importing fyne
import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// creating new app
	a := app.New()
	// New window
	w := a.NewWindow("Here is my title for 4th tutorial")
	// resizing my window
	w.Resize(fyne.NewSize(400, 400))
	// Now Its Time to use BUTTONs
	// First value is button name
	// 2nd value is any function
	btn := widget.NewButton("button Name", func() {
		// Our is ready
		fmt.Println("Here is Go Button")
	})
	// using our button on
	//w.SetContent(btn)
	// check box widget
	checkbox1 := widget.NewCheck("Male", func(b bool) {
		if b == true {
			fmt.Println("Male")
		} else {
			fmt.Println("Not Male")
		}
	})
	// set up content
	//w.SetContent(checkbox1)
	// creating url
	url, _ := url.Parse("https://google.com")
	// hyperLink Widget
	// first value is label
	// 2nd value is URL/ website address
	hyperlink := widget.NewHyperlink("Visit me", url)
	// setup content
	w.SetContent(
		container.NewVBox(btn, checkbox1, hyperlink),
		)
	// Running app
	w.ShowAndRun()
}
