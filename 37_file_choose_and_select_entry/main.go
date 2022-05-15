package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jasonlvhit/gocron"
	"sync/atomic"
)

var SendCount int32 = 0

func taskWithParams(Label4 *widget.Label, File, SelectEntry string) {
	atomic.AddInt32(&SendCount, 1)
	if SendCount%10 == 0 {
		fmt.Printf("I am running %s %s. Send count to %d.\n", File, SelectEntry, SendCount)
		Label4.Text = fmt.Sprintf("Send count to : %d", SendCount)
		Label4.Refresh()
	}
}

func main() {
	var DataType string
	var FilePath string

	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("Select entry widget, drop down")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets show our selected entry in label
	label1 := widget.NewLabel("data type : ")
	label2 := widget.NewLabel("file : ")
	label3 := widget.NewLabel("ready???")
	label4 := widget.NewLabel("Count to:")

	// dropdown/ select entry
	//[]string{} all our option goes in slice
	// s is the variable to get the selected value
	dd := widget.NewSelect(
		[]string{"temperature", "battery_usage"},
		func(s string) {
			DataType = s
			label1.Text = fmt.Sprintf("data type : %s", s)
			label1.Refresh()
		})
	// more than one widget. so use container

	btn1 := widget.NewButton("Click to choose a file", func() {
		dialog.ShowFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				fmt.Println(r.URI().Path())
				FilePath = r.URI().Path()
				entry := widget.NewMultiLineEntry()
				entry.SetText(r.URI().Path())
				entry.Wrapping = fyne.TextWrapBreak
				label2.Text = fmt.Sprintf("file : %s", FilePath)
				label2.Refresh()
			}, w)
	})

	btn2 := widget.NewButton("start", func() {
		label3.Text = " GO!!! " + label1.Text + " " + label2.Text
		label3.Refresh()
		// Clear all scheduled jobs
		label4.Text = fmt.Sprintf("Send count to : 0")
		label4.Refresh()
		err := gocron.Every(1).Second().Do(taskWithParams, label4, FilePath, DataType)
		if err != nil {
			panic(err)
		}
		// Start all the pending jobs
		go func() {
			<-gocron.Start()
		}()

	})

	btn3 := widget.NewButton("stop", func() {
		fmt.Println("clear")
		gocron.Clear()
		atomic.StoreInt32(&SendCount, 0)

	})

	c := container.NewVBox(dd, label1, btn1, label2, btn2, label3, label4, btn3)
	w.SetContent(c)
	//show and run
	w.ShowAndRun()
}
