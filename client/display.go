package client

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type InputBuf struct {
	//多线程读取用户输入，上锁处理
	Content string
}

// TODO 实现客户端消息展示功能
func Display() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var inputBuf InputBuf = InputBuf{
		Content: "",
	}

	// 创建一个管道用于接收消息
	messageReceiveChan := make(chan string)
	messageInputChan := make(chan string)
	//messageInputBufChan := make(chan string)

	go receiveMessages(messageReceiveChan, messageInputChan) // 启动协程接收消息
	go userInput(messageInputChan, &inputBuf)                // 启动协程接收用户输入
	//go getUserInput(messageInputChan, reader) // 启动协程接收用户输入

	for {
		drawUI(messageReceiveChan, &inputBuf) // 绘制UI界面
	}
	//for {
	//	fmt.Println(messageInputChan)
	//	fmt.Println("inputBuf.Content:", inputBuf.Content)
	//}
	//for {
	//	drawInputBox(&inputBuf, 1, 1, 50)
	//}
}

func receiveMessages(messages chan<- string, messageInputChan <-chan string) {
	//for message := range messageInputChan {
	//	messages <- message
	//}
	//这里模拟接收消息的过程，可以根据实际需求进行修改
	for i := 1; ; i++ {
		message := fmt.Sprintf("This is receive %d ", i)
		//messages <- message
		messages <- message
	}
}

func drawUI(recv <-chan string, inp *InputBuf) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	width, height := termbox.Size()

	// 计算分割线的位置
	dividerPos := height / 2

	// 绘制分割线
	for x := 0; x < width; x++ {
		termbox.SetCell(x, dividerPos, '-', termbox.ColorDefault, termbox.ColorDefault)
	}

	drawMessages(recv, 0, 0, dividerPos-1)
	//// 绘制消息
	//go drawMessages(inp, 0, dividerPos+1, height)
	//drawInputBox(inp, 0, dividerPos+1, height)
	termbox.Flush()
}

func drawMessages(messages <-chan string, startX, startY, endY int) {
	x, y := startX, startY

	for message := range messages {
		for _, ch := range message {
			termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}

		// 换行
		x = startX
		y++

		// 如果超出终端上方区域，停止绘制
		if y >= endY {
			break
		}
	}
}

func drawInputBox(inpBuf *InputBuf, startX, startY, endY int) {

	x, y := startX, startY

	for _, ch := range inpBuf.Content {
		termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
		x++

		// 换行
		x = startX
		y++

		// 如果超出终端上方区域，停止绘制
		if y >= endY {
			break
		}
	}
}

func userInput(inpText chan string, inputbuf *InputBuf) {

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
				inputbuf.Content = ""
				inputText = ""
				//inpText <- inputText

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
