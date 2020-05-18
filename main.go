package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/semyon-dev/whissage-desktop/model"
	"strconv"
	"time"
)

var a fyne.App
var username string

func main() {
	a = app.New()
	w := a.NewWindow("Whissage")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))

	usernameEntry := widget.NewEntry()

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Your username:"),
		usernameEntry,
		widget.NewButton("Let's go", func() {
			if usernameEntry.Text != "" {
				username = usernameEntry.Text
				w.Close()
				chatWindow()
			}
		}),
	))

	w.ShowAndRun()
}

func chatWindow() {
	w := a.NewWindow("Chat with " + username)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))

	messageEntry := widget.NewEntry()

	content := widget.NewVBox(widget.NewLabel(""),
		widget.NewHBox(
			widget.NewLabel("New message"),
			messageEntry,
			widget.NewButton("Send", func() {
				send(messageEntry.Text)
			})))

	w.SetContent(content)
	w.Show()

	content.Append(widget.NewButton("Add more items", func() {
		content.Append(widget.NewLabel("message"))
	}))
}

func send(text string) {
	var msg model.Message
	msg.Message = text
	msg.User = username
	msg.Time = strconv.Itoa(int(time.Now().Unix()))
	msgBytes, _ := json.Marshal(msg)
	fmt.Println(msgBytes)
}
