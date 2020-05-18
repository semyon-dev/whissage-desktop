package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var a fyne.App

func main() {
	a = app.New()
	w := a.NewWindow("Whissage")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Login:"),
		widget.NewEntry(),
		widget.NewButton("Let's go", func() {
			w.Close()
			chatWindow()
		}),
	))

	w.ShowAndRun()
}

func chatWindow() {
	w := a.NewWindow("Chat")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewVBox(
		widget.NewLabel("New message"),
		widget.NewEntry(),
		widget.NewButton("Send", func() {
			send()
		}),
	))
	w.Show()
}

func send() {

}
