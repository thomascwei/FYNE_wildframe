package main

// import fyne

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	// new app
	a := app.New()
	//New title and window
	w := a.NewWindow("Weather API & Fyne")

	//Consuming Api
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=peshawar&APPID=88a3325d8b543b9103c71abe0ebc15ef")
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	//
	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}

	// Now we are almost done
	img := canvas.NewImageFromFile("C:\\Users\\Anna\\Pictures\\dice\\dice1.jpg")
	img.FillMode = canvas.ImageFillOriginal
	label1 := canvas.NewText("Weather Api & fyne", color.Black)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label2 := canvas.NewText(fmt.Sprintf("Country %s", weather.Sys.Country), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("City %s", weather.Name), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("wind speed %.2f", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("TEMP %.2f F", weather.Main.Temp), color.Black)

	w.SetContent(
		container.NewVBox(

			label1,
			img,
			label2,
			label5,
			label3,
			label4,
		),
	)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun() // show and run app
}

// Use api.QuickType.io to Convert Json response

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

// package main

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}