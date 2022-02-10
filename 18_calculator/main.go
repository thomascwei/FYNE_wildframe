package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {

	a := app.New()
	w := a.NewWindow("calc")
	w.Resize(fyne.NewSize(400, 400))

	ClearRequest := false
	var entryText = ""

	entry1 := widget.NewEntry()
	entry1.TextStyle = fyne.TextStyle{Bold: true}
	entry1.SetPlaceHolder("Calculator..")

	btn1 := widget.NewButton("1", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false
		}
		entryText = entryText + "1"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn2 := widget.NewButton("2", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "2"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btn3 := widget.NewButton("3", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "3"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btn4 := widget.NewButton("4", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "4"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn5 := widget.NewButton("5", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "5"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn6 := widget.NewButton("6", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "6"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn7 := widget.NewButton("7", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "7"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn8 := widget.NewButton("8", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "8"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn9 := widget.NewButton("9", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "9"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btn0 := widget.NewButton("0", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "0"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnDot := widget.NewButton(".", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "."
		entry1.SetText(fmt.Sprint(entryText))
	})

	result := canvas.NewText("Result", color.Black)
	result.TextSize = 30
	result.Alignment = fyne.TextAlignCenter
	btnClear := widget.NewButton("Clear", func() {
		entry1.SetText("")
		result.Text = ""
		entryText = ""
	})
	// btnClear.Text = "NoClear"
	btnPlus := widget.NewButton("+", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "+"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnMinus := widget.NewButton("-", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "-"
		entry1.SetText(fmt.Sprint(entryText))

	})
	btnMultiply := widget.NewButton("x", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "*"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnDivide := widget.NewButton("/", func() {
		if ClearRequest {
			entryText = ""
			ClearRequest = false

		}
		entryText = entryText + "/"
		entry1.SetText(fmt.Sprint(entryText))
	})
	btnEqual := widget.NewButton("=", func() {
		// 初始化eval功能
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		// 套件程式碼, 照抄
		fmt.Println(entry1.Text)
		if string(entry1.Text[0]) =="/" || string(entry1.Text[0]) =="*"{
			entry1.SetText("")
			result.Text = ""
			entryText = ""
			return
		}
		// 避免整數相除的商只show整數
		if strings.Contains(entry1.Text, "/") {
			entry1.Text = "1.0*" + entry1.Text
		}
		res, err := i.Eval(entry1.Text)
		if err != nil {
			log.Fatal(err)
		}
		resp, ok := res.Interface().(int)
		if !ok {
			resp, _ := res.Interface().(float64)
			result.Text = strconv.FormatFloat(resp, 'f', -1, 64)
			result.Refresh()
		} else {
			result.Text = strconv.Itoa(resp)
			result.Refresh()
		}
		ClearRequest = true
	})

	w.SetContent(
		container.NewVBox(

			entry1,
			result,
			container.NewGridWithColumns(
				5,
				btn1,
				btn2,
				btn3,
				btn4,
				btn5,
				btn6,
				btn7,
				btn8,
				btn9,
				btn0,
				btnDot,
				btnEqual,
				btnPlus,
				btnMinus,
				btnMultiply,
				btnDivide,
				btnClear,
			),
		),
	)
	w.ShowAndRun()
}
