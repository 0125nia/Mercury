package client

import (
	"fmt"
	"log"
	"os"

	"github.com/0125nia/Mercury/common/sdk"
	"github.com/gookit/color"
	"github.com/rocket049/gocui"
)

var (
	chat *sdk.Chat
	buf  string
)

// RunMain client main logic function
func RunMain() {
	// init chat
	chat = sdk.NewChat("127.0.0.1:8080", "mercury", "123123", "123123")

	// init gocui
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln("gocui init err: ", err)
	}
	// enable cursor, disable mouse and ascii
	g.Cursor = true
	g.Mouse = false
	g.ASCII = false
	// set layout
	g.SetManagerFunc(layout)

	// register callback events
	keysbindings(g)

	// async receiving message
	go receiveMsg(g)

	// main loop
	if err := g.MainLoop(); err != nil {
		log.Println(err)
	}

	// write log to file
	os.WriteFile("chat.log", []byte(buf), 0644)
}

// receiveMsg goroutine for receiving message
func receiveMsg(g *gocui.Gui) {
	recv := chat.Recv()
	for message := range recv {
		switch message.Type {
		case sdk.MsgTypeText:
			outPrint(g, message.Name, message.Content)
		}
	}
	g.Close()
}

// MsgOutput is a struct that represents a message that is sent between clients.
type MsgOutput struct {
	Name, Msg string
}

// Show shows the message in the out view
func (out MsgOutput) Show(g *gocui.Gui) error {
	v, err := g.View(output)
	if err != nil {
		return nil
	}
	fmt.Fprintf(v, "%v: %v\n", color.FgGreen.Text(out.Name), color.FgYellow.Text(out.Msg))
	return nil
}

// outPrint prints the message in the out view
func outPrint(g *gocui.Gui, name, msg string) {
	var out MsgOutput
	out.Name = name
	out.Msg = msg
	g.Update(out.Show)
}
