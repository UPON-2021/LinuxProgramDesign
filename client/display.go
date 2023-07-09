package client

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type InputBuf struct {
	Content string
}

func Display() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var inputBuf InputBuf = InputBuf{
		Content: "",
	}

	messageReceiveChan := make(chan string)
	userInputChan := make(chan string)

	go receiveMessages(messageReceiveChan, userInputChan)
	go userInput(userInputChan, &inputBuf)

	for {
		drawUI(messageReceiveChan, &inputBuf)
	}
}

func receiveMessages(messages chan<- string, userInputChan <-chan string) {
	for message := range userInputChan {
		messages <- message
	}
}

func drawUI(recv <-chan string, inp *InputBuf) {
	_, height := termbox.Size()

	dividerPos := height / 2

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawMessages(recv, 0, 0, dividerPos-1)
	drawInputBox(inp, 0, dividerPos+1, height)

	termbox.Flush()
}

func drawMessages(messages <-chan string, startX, startY, endY int) {
	x, y := startX, startY

	for message := range messages {
		lines := strings.Split(message, "\n")

		for _, line := range lines {
			for _, ch := range line {
				termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
				x++
			}

			x = startX
			y++

			if y >= endY {
				break
			}
		}
	}
}

func drawInputBox(inpBuf *InputBuf, startX, startY, endY int) {
	x, y := startX, startY

	for _, ch := range inpBuf.Content {
		termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
		x++

		x1, _ := termbox.Size()
		if x >= x1 {
			x = startX
			y++
		}

		if y >= endY {
			break
		}
	}
}

func userInput(userInputChan chan<- string, inputbuf *InputBuf) {
	inputText := inputbuf.Content

mainLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainLoop
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				if len(inputText) > 0 {
					inputText = inputText[:len(inputText)-1]
				}
			case termbox.KeySpace:
				inputText += " "
			case termbox.KeyEnter:
				userInputChan <- inputText
				inputbuf.Content = ""
				inputText = ""
			default:
				if ev.Ch != 0 {
					inputText += string(ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		inputbuf.Content = inputText
	}
}
