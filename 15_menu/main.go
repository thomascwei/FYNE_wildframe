package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	// New app
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	//app title
	w := a.NewWindow("Main Menu")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// Menu Items
	menuItem1 := fyne.NewMenuItem("New", func() { fmt.Println("New pressed") })
	menuItem2 := fyne.NewMenuItem("Save", func() { fmt.Println("Save pressed") })
	menuItem3 := fyne.NewMenuItem("edit", nil)
	newMenu := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)   // main menu
	newMenu2 := fyne.NewMenu("Other", menuItem1, menuItem2, menuItem3) // main menu
	newMenu3 := fyne.NewMenu("Help", menuItem1, menuItem2, menuItem3)

	// creating new main menu
	menu := fyne.NewMainMenu(newMenu, newMenu2, newMenu3)
	w.SetMainMenu(menu) /// we are done 🙂
	//
	w.ShowAndRun()
}
