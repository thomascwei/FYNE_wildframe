package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"image/color"
	"os"
	"time"
)

var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause = false

func makeCell() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{128, 128, 128, 255})
	rect.SetMinSize(fyne.NewSize(30, 30))
	return rect
}

func main() {
	go func(msg string) {
		fmt.Println(msg)
		if streamer == nil {
		} else {
			//slider.Value = float64(streamer.Position())
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")
	time.Sleep(time.Second)
	a := app.New()
	w := a.NewWindow("audio player...")
	w.Resize(fyne.NewSize(400, 400))
	logo := canvas.NewImageFromFile("111.png")
	logo.FillMode = canvas.ImageFillOriginal
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			// f, _ = os.Open("hen.mp3")
			fmt.Println("play")
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		//widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			fmt.Println("pause")

			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		//widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			fmt.Println("stop")
			speaker.Clear()
			speaker.Close()
		}),
		widget.NewToolbarSpacer(),
	)
	label := widget.NewLabel("Audio MP3..")
	label.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel("Play MP3..")
	label2.Alignment = fyne.TextAlignCenter
	browse_files := widget.NewButton("Browse...", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})
	// slider := widget.NewSlider(0, 100)
	c := container.NewVBox(label, browse_files, label2, toolbar)
	w.SetContent(
		container.NewBorder(logo, makeCell(), makeCell(), makeCell(), c),
	)
	w.ShowAndRun()
}
