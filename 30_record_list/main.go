package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
)

func main() {
	// reading data from file
	res, _ := ioutil.ReadFile("myfilename.json")
	// creating struct for my data
	type Student struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	// slice/array to store all data
	var mydata []Student
	// parse json data
	json.Unmarshal(res, &mydata)
	a := app.New()
	w := a.NewWindow("Student Record List")
	w.Resize(fyne.NewSize(400, 400))
	list := widget.NewList(
		func() int { return len(mydata) },
		func() fyne.CanvasObject {
			return widget.NewLabel("My list item widget")
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(mydata[id].Name)
		},
	)
	label1 := widget.NewLabel("...")
	label2 := widget.NewLabel("...")
	list.OnSelected = func(id widget.ListItemID) {
		label1.Text = mydata[id].Name
		label1.TextStyle = fyne.TextStyle{Bold: true}
		label2.Text = mydata[id].Phone
		// label2.TextStyle = fyne.TextStyle{Bold: true}
		label1.Refresh()
		label2.Refresh()
	}
	e_name := widget.NewEntry()
	e_name.SetPlaceHolder("Enter name ...")
	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter phone ...")
	e_submit := widget.NewButton("Submit", func() {
		if e_name.Text != "" && e_phone.Text != "" {
			obj1 := &Student{
				Name:  e_name.Text,
				Phone: e_phone.Text,
			}
			mydata = append(mydata, *obj1)
			b, _ := json.MarshalIndent(mydata, "", " ")
			os.WriteFile("myfilename.json", b, 0644)
			e_name.Text = ""
			e_phone.Text = ""
			e_name.Refresh()
			e_phone.Refresh()
			list.Refresh()
		} else {
			label1.Text = "Incorrect data"
			label2.Text = ""
			label2.Refresh()
			label1.Refresh()
		}
	})
	e_del := widget.NewButton("DELETE", func() {
		var mydata2 []Student
		for _, a := range mydata {
			if a.Name != label1.Text {
				mydata2 = append(mydata2, a)
				fmt.Println(a)
			}
		}
		mydata = mydata2
		b, _ := json.MarshalIndent(mydata2, "", " ")
		os.WriteFile("myfilename.json", b, 0644)
		e_name.Text = ""
		e_phone.Text = ""
		label1.Text = ""
		label2.Text = ""
		e_name.Refresh()
		e_phone.Refresh()
		label1.Refresh()
		label2.Refresh()
		list.Refresh()
	})
	e_update := widget.NewButton("Update", func() {
		if e_name.Text != "" && e_phone.Text != "" {
			var mydata2 []Student
			for _, a := range mydata {
				if a.Name != label1.Text {
					mydata2 = append(mydata2, a)
					fmt.Println(a)
				} else {
					obj1 := &Student{
						Name:  e_name.Text,
						Phone: e_phone.Text,
					}
					mydata2 = append(mydata2, *obj1)
					fmt.Println(a)
				}
			}
			mydata = mydata2
			b, _ := json.MarshalIndent(mydata2, "", " ")
			os.WriteFile("myfilename.json", b, 0644)
			label1.Text=e_name.Text
			label2.Text=e_phone.Text
			label1.Refresh()
			label2.Refresh()
			e_name.Text = ""
			e_phone.Text = ""
			e_name.Refresh()
			e_phone.Refresh()
			list.Refresh()

		} else {
			label1.Text = "Incorrect data"
			label2.Text = ""
			label2.Refresh()
			label1.Refresh()
		}
	})
	c := container.NewVBox(label1, label2,
		e_name, e_phone, e_submit, e_del, e_update)
	w.SetContent(container.NewHSplit(
		list,
		c,
	))
	w.ShowAndRun()
	fmt.Println(mydata)
}
