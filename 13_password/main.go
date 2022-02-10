package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"
	// import fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New App
	a := app.New()
	// New title and window
	w := a.NewWindow("Password Generator")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	title := canvas.NewText("Password Generator", color.Black)
	// input Box
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter password length")

	//label to show password
	text := canvas.NewText("", color.Black)
	text.TextSize = 20
	// button to generate password
	btn1 := widget.NewButton("Generate", func() {
		// input
		passlength, _ := strconv.Atoi(input.Text) // convert string to integer
		text.Text = PasswordGenerator(passlength)
		text.Refresh()

	})
	// show content
	w.SetContent(
		container.NewVBox(
			// put all widgets here
			title,
			input,
			text,
			btn1,
		),
	)
	w.ShowAndRun()
	// Fyne appp
}

// Converting the code to a function
func PasswordGenerator(passwordLength int) string {
	// Password Generator
	// Lower case
	lowCase := "abcdefghijklmnopqrstuvxyz"
	// Upper Case
	upCase := "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	// Numbers
	Numbers := "0123456789"
	// Special characters
	SpecialChar := "£$&()*+[]@#^-_!?"
	// Password Length
	// passwordLength := 8
	// variable for storing password
	password := ""
	// loop
	for n := 0; n < passwordLength; n++ {
		// Now random characters
		time.Sleep(time.Nanosecond)
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(4)
		fmt.Println(randNum)
		// Switch statment
		switch randNum {
		case 0:
			//rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(lowCase))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(lowCase[randNum])
			// it is byte... we need string
		// first case completed
		case 1:
			//rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(upCase))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(upCase[randNum])
			// it is byte... we need string
			// second done 🙂
		case 2:
			//rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(Numbers))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(Numbers[randNum])
			// it is byte... we need string
			// 3rd done 🙂
		case 3:
			//rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(SpecialChar))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(SpecialChar[randNum])
			// it is byte... we need string
		} // end of switch
	} // end of for loop
	fmt.Println(password)
	// return password
	return password
}
