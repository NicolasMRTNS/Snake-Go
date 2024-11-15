package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody SnakeBody
}

func drawParts(screen tcell.Screen, parts []SnakePart, style tcell.Style) {
	for _, part := range parts {
		screen.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func (game *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	game.Screen.SetStyle(defStyle)
	width, height := game.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		game.Screen.Clear()
		game.SnakeBody.Update(width, height)
		drawParts(game.Screen, game.SnakeBody.Parts, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		game.Screen.Show()
	}
}
