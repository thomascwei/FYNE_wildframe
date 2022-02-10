package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// new app
	a := app.New()
	// title of app
	w := a.NewWindow("Child Menu")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// new menu items
	//first parameter is label, 2nd is function
	item1 := fyne.NewMenuItem("edit", nil)
	item2 := fyne.NewMenuItem("details", nil)
	item3 := fyne.NewMenuItem("home", nil)
	item4 := fyne.NewMenuItem("run", nil)
	// child menu
	item1.ChildMenu = fyne.NewMenu(
		"",                            // leave label blank
		fyne.NewMenuItem("copy", nil), // child menu items
		fyne.NewMenuItem("cut", nil),
		fyne.NewMenuItem("paste", nil),
	)
	// create child menu for 2nd item
	item2.ChildMenu = fyne.NewMenu(
		"",                             // leave label blank
		fyne.NewMenuItem("books", nil), // child menu items
		fyne.NewMenuItem("magzine", nil),
		fyne.NewMenuItem("notebook", nil),
	)
	// create child menu for third item
	item3.ChildMenu = fyne.NewMenu(
		"",                              // leave label blank
		fyne.NewMenuItem("school", nil), // child menu items
		fyne.NewMenuItem("college", nil),
		fyne.NewMenuItem("university", nil),
	)
	NewMenu1 := fyne.NewMenu("File", item1, item2, item3, item4)
	NewMenu2 := fyne.NewMenu("Help", item1, item2, item3, item4)
	// main menu
	menu := fyne.NewMainMenu(NewMenu1, NewMenu2)
	// setup menu
	w.SetMainMenu(menu) /// we are done 🙂
	// show and run
	w.ShowAndRun()
}
