package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jasonlvhit/gocron"
	"sync/atomic"
)

var SendCount int32 = 0

func taskWithParams(Label4 *widget.Label, File, SelectEntry string) {
	atomic.AddInt32(&SendCount, 1)
	fmt.Printf("I am running %s %s. Send count : %d.\n", File, SelectEntry, SendCount)
	if SendCount%10 == 0 {
		Label4.Text = fmt.Sprintf("Send count : %d", SendCount)
		Label4.Refresh()
	}
}

func main() {
	var DataType string
	var FilePath string

	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("VAST stress test LAS60X")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets show our selected entry in label
	label1 := widget.NewLabel("data type : ")
	label2 := widget.NewLabel("file : ")
	label3 := widget.NewLabel("")
	label4 := widget.NewLabel("Send count : 0")

	dd := widget.NewSelect(
		[]string{"temperature", "battery_usage"},
		func(s string) {
			DataType = s
			label1.Text = fmt.Sprintf("data type : %s", s)
			label1.Refresh()
		})
	// more than one widget. so use container
	//butOuter:=widget.NewToolbarSeparator()

	btn1 := widget.NewButtonWithIcon("click to choose a file", theme.FileTextIcon(), func() {
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
	btn1.Resize(fyne.NewSize(200, 30))
	btn1.Move(fyne.NewPos(40, 5))

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			if DataType == "" {
				label3.Text = "DataType Error"
				label3.Refresh()
			} else if FilePath == "" {
				label3.Text = "File Error"
				label3.Refresh()
			} else {
				label3.Text = "GO!!! " + label1.Text + " " + label2.Text
				label3.Refresh()
				// Clear all scheduled jobs
				label4.Text = fmt.Sprintf("Send count : 0")
				label4.Refresh()
				gocron.Clear()
				atomic.StoreInt32(&SendCount, 0)
				err := gocron.Every(1).Second().Do(taskWithParams, label4, FilePath, DataType)
				if err != nil {
					panic(err)
				}
				// Start all the pending jobs
				go func() {
					<-gocron.Start()
				}()
			}
		}),
		//widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			label3.Text = "STOP!!!"
			label3.Refresh()
			fmt.Println(label3.Text)
			gocron.Clear()
			atomic.StoreInt32(&SendCount, 0)
		}),
		widget.NewToolbarSpacer(),
	)

	btn11 := container.NewWithoutLayout(btn1)
	c := container.NewVBox(dd, label1, btn11, label2, label3, label4, toolbar)
	w.SetContent(c)

	//show and run
	w.ShowAndRun()
}
