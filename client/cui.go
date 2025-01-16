package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/0125nia/Mercury/client/sdk"
	"github.com/gookit/color"
	"github.com/rocket049/gocui"
)

const (
	head   = "head"
	output = "out"
	input  = "main"
)

// 布局函数
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if err := headView(g, 1, 1, maxX-1, 3); err != nil {
		return err
	}
	if err := outputView(g, 1, 4, maxX-1, maxY-4); err != nil {
		return err
	}
	if err := inputView(g, 1, maxY-3, maxX-1, maxY-1); err != nil {
		return err
	}
	return nil
}

func keysbindings(g *gocui.Gui) {
	//Ctrl + c quit
	if err := g.SetKeybinding(input, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	//enter send message
	if err := g.SetKeybinding(input, gocui.KeyEnter, gocui.ModNone, sendAndUpdate); err != nil {
		log.Panicln(err)
	}
	//pgUp Look up historical information
	if err := g.SetKeybinding(input, gocui.KeyPgup, gocui.ModNone, scrollUpMsg); err != nil {
		log.Panicln(err)
	}
	//pgDown Look down historical information
	if err := g.SetKeybinding(input, gocui.KeyPgdn, gocui.ModNone, scrollDownMsg); err != nil {
		log.Panicln(err)
	}
	//↑ show the previous message in the input box
	if err := g.SetKeybinding(input, gocui.KeyArrowUp, gocui.ModNone, parseUpMsg); err != nil {
		log.Panicln(err)
	}
	//↓ show the next message in the input box
	if err := g.SetKeybinding(input, gocui.KeyArrowDown, gocui.ModNone, parseDownMsg); err != nil {
		log.Panicln(err)
	}
}

// headView set the head view
func headView(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView(head, x0, y0, x1, y1); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		// disable wrap
		v.Wrap = false
		// enable overwrite
		v.Overwrite = true
		setHeadText(g, "start chat!")
	}
	return nil
}

// setHeadText set the text information of the Head view
func setHeadText(g *gocui.Gui, msg string) {
	v, err := g.View(head)
	if err == nil {
		// clear the view
		v.Clear()
		fmt.Fprint(v, color.FgGreen.Text(msg))
	}
}

// outputView set the output view
func outputView(g *gocui.Gui, x0, y0, x1, y1 int) error {
	v, err := g.SetView(output, x0, y0, x1, y1)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		// enable wrap
		v.Wrap = true
		// disable overwrite
		v.Overwrite = false
		// enable autoscroll
		v.Autoscroll = true
		// set the bg color
		v.SelBgColor = gocui.ColorRed
		// set the title
		v.Title = "Messages"
	}
	return nil
}

// inputView set the input view
func inputView(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView(input, x0, y0, x1, y1); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		// when the view is not found, create a new view
		// enable editable
		v.Editable = true
		// enable wrap
		v.Wrap = true
		// disable overwrite
		v.Overwrite = false
		// set the current view to input
		if _, err := g.SetCurrentView(input); err != nil {
			return err
		}
	}
	return nil
}

// quit quit the gocui
func quit(g *gocui.Gui, _ *gocui.View) error {
	// release resource
	chat.Close()

	// quit the gocui
	ov, _ := g.View(output)
	buf = ov.Buffer()
	g.Close()
	// ErrQuit is used to decide if the MainLoop finished successfully
	return gocui.ErrQuit
}

// sendAndUpdate send the message and update the view
func sendAndUpdate(g *gocui.Gui, v *gocui.View) error {
	// send message
	sendMsg(g, v)
	// get the len of the input view buf
	l := len(v.Buffer())
	// move the cursor to the start of the input view
	v.MoveCursor(0-l, 0, true)
	// clear the input view
	v.Clear()
	return nil
}

// sendMsg send the message by reading the content of the input view
func sendMsg(g *gocui.Gui, cv *gocui.View) {
	v, err := g.View(output)
	if cv != nil && err == nil {
		// read the content of the input view
		editor := cv.ReadEditor()
		if editor != nil {
			msg := &sdk.Message{
				Type:    sdk.MsgTypeText,
				Name:    "nia",
				FromId:  "1",
				ToId:    "2",
				Session: "123123",
				Content: string(editor)}
			outPrint(g, "me", msg.Content)
			chat.Send(msg)
		}
		v.Autoscroll = true
	}
}

// scrollUpMsg Scroll up the output view
func scrollUpMsg(g *gocui.Gui, _ *gocui.View) error {
	v, err := g.View(output)
	// disable autoscroll
	v.Autoscroll = false
	// get the origin point of the view
	x, y := v.Origin()

	if err == nil {
		// Move the view scroll origin down one unit
		_ = v.SetOrigin(x, y-1)
	}
	return nil
}

// scrollDownMsg Scroll down the output view
func scrollDownMsg(g *gocui.Gui, _ *gocui.View) error {
	v, err := g.View(output)

	// get the size of the view
	_, y := v.Size()
	ox, oy := v.Origin()

	// calculate the total number of lines in the view
	l := len(v.BufferLines())

	if err == nil {
		// where is the bottom: The scrolling origin of the view is the total number of message lines minus the bottom of the view and then minus 1 to represent the bottom
		// If scroll to the bottom, set it to auto scroll, so that new content can be automatically scrolled to display the latest message when added to the view
		if oy > l-y-1 {
			v.Autoscroll = true
		} else {
			// if not at the bottom, scroll down one line
			_ = v.SetOrigin(ox, oy+1)
		}
	}
	return nil
}

// message index
var pos int

// parseUpMsg show the previous message in the input view
func parseUpMsg(g *gocui.Gui, iv *gocui.View) error {
	v, err := g.View(output)
	if err != nil {
		fmt.Fprintf(iv, "error:%s", err)
		return nil
	}
	// get the total number of lines in the view
	lines := v.BufferLines()
	l := len(lines)

	// If there are more historical messages to view, point to the next message
	if pos < l-1 {
		pos++
	}
	// clear the input view
	iv.Clear()
	// fill with the history message content at the pos index
	fmt.Fprintf(iv, "%s", lines[l-pos-1])
	return nil
}

// parseDownMsg show the next message in the input view
func parseDownMsg(g *gocui.Gui, iv *gocui.View) error {
	// get the output view
	v, err := g.View(output)

	if err != nil {
		fmt.Fprintf(iv, "error:%s", err)
		return nil
	}
	if pos > 0 {
		pos--
	}
	// get the information in the buffer
	lines := v.BufferLines()
	l := len(lines)

	// clear the input view
	iv.Clear()
	// fill with the history message content at the pos index
	fmt.Fprintf(iv, "%s", lines[l-pos-1])
	return nil
}
