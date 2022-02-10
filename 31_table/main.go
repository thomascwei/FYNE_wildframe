package main
// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("Table widget")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// create table widget
	// table is like list , just 2 values, instead of one
	//table := widget.NewTable(
	//	// row and col . I want to create 3 row , 3 col
	//	func() (int, int) { return 3, 3 },
	//	// Now I want to specify widget. like label, checkbox
	//	func() fyne.CanvasObject { return widget.NewLabel("....") },
	//	// update/details data in widget
	//	func(i widget.TableCellID, obj fyne.CanvasObject) {
	//		// remember it is label and not newlabel
	//		// i is for index
	//		// i.col will set col value
	//		// i.row will present row value
	//		obj.(*widget.Label).SetText(fmt.Sprintf("%d %d", i.Col, i.Row))
	//	},
	//)

	t := widget.NewTable(
		func() (int, int) { return 5, 10 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText("A longer cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
			}
		})
	//t.SetColumnWidth(0, 34)
	t.SetColumnWidth(1, 102)
	w.SetContent(t)
	w.ShowAndRun()
}