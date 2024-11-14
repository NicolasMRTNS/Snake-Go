package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody SnakeBody
}

func (game *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	game.Screen.SetStyle(defStyle)
	width, height := game.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		game.Screen.Clear()
		game.SnakeBody.Update(width, height)
		game.Screen.SetContent(game.SnakeBody.X, game.SnakeBody.Y, ' ', nil, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		game.Screen.Show()
	}
}
