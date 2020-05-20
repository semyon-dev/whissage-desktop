package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gorilla/websocket"
	"github.com/semyon-dev/whissage-desktop/config"
	"github.com/semyon-dev/whissage-desktop/model"
	"log"
	"net/url"
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

var scroll *widget.Group

func chatWindow() {
	w := a.NewWindow("Chat with " + username)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))

	messageEntry := widget.NewEntry()
	box := widget.NewVBox(
		widget.NewLabel("New message"),
		messageEntry,
		widget.NewButton("Quit & close connection", func() {
			closeConn()
			w.Close()
			a.Quit()
		}),
		widget.NewButton("Send", func() {
			if messageEntry.Text != "" {
				send(messageEntry.Text)
				messageEntry.Text = ""
			}
		}))

	scroll = widget.NewGroupWithScroller("title")
	w.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, box, nil, nil), scroll, box))
	w.Show()
	connect()
}

func appendMessage(text []byte) {
	var data model.Message
	err := json.Unmarshal(text, &data)
	if err != nil {
		fmt.Println(err)
	} else {
		scroll.Prepend(widget.NewLabel("from: " + data.User + " message: " + data.Message))
	}
}

func send(text string) {
	var msg model.Message
	msg.Message = text
	msg.User = username
	msg.Time = strconv.Itoa(int(time.Now().Unix()))
	msgBytes, _ := json.Marshal(msg)
	err := c.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		fmt.Println("write:", err)
	}
}

var c *websocket.Conn

func closeConn() {
	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	defer c.Close()
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		fmt.Println("write close:", err)
	}
}

func connect() {

	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: config.Url, Path: "/ws/"}
	fmt.Printf("connecting to %s", u.String())

	var err error
	c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			fmt.Printf("recv: %s", message)
			appendMessage(message)
		}
	}()
}
