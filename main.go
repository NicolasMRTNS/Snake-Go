package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snakeBody := SnakeBody{
		X:      5,
		Y:      10,
		Xspeed: 1,
		Yspeed: 0,
	}

	game := Game{
		Screen:    screen,
		SnakeBody: snakeBody,
	}

	go game.Run()

	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()

		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.SnakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				game.SnakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.SnakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.SnakeBody.ChangeDir(0, 1)
			}
		}
	}
}