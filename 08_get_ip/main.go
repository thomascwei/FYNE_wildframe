package main

// import Fyne
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	// New Title and Window
	w := a.NewWindow("What is my IP?")
	// Resize
	w.Resize(fyne.NewSize(400, 400))
	// Let Creat UI First
	//Title
	labelTitle := widget.NewLabel("What is my IP?")
	labelIP := widget.NewLabel("Your IP is ...")
	label_Value := widget.NewLabel("...")
	label_City := widget.NewLabel("...")
	label_Country := widget.NewLabel("...")
	btn := widget.NewButton("Run", func() {
		// Logic
		myinfo := MyInfo()
		label_Value.Text = myinfo.Query
		label_Value.Refresh()
		label_City.Text = myinfo.City
		label_City.Refresh()
		label_Country.Text = myinfo.Country
		label_Country.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			labelTitle,
			labelIP,
			label_Value,
			label_City,
			label_Country,
			btn,
		),
	)
	// Show and Run
	w.ShowAndRun()
}

func MyInfo() IP {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return IP{}
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return IP{}
	}
	var ip IP
	json.Unmarshal(body, &ip)

	return ip
}

type IP struct {
	Query   string `json:"query"`
	Country string `json:"country"`
	City    string `json:"city"`
}
