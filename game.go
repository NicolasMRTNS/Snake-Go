package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody SnakeBody
	FoodPos   Part
	Score     int
	GameOver  bool
}

func drawParts(screen tcell.Screen, snakeParts []Part, foodPos Part, snakeStyle tcell.Style, foodStyle tcell.Style) {
	screen.SetContent(foodPos.X, foodPos.Y, '\u25CF', nil, foodStyle)
	for _, part := range snakeParts {
		screen.SetContent(part.X, part.Y, ' ', nil, snakeStyle)
	}
}

func drawText(screen tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, r := range text {
		screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func checkCollison(parts []Part, otherPart Part) bool {
	for _, part := range parts {
		if part.X == otherPart.X && part.Y == otherPart.Y {
			return true
		}
	}
	return false
}

func (game *Game) UpdateFoodPos(width, height int) {
	game.FoodPos.X = rand.Intn(width)
	game.FoodPos.Y = rand.Intn(height)
	if game.FoodPos.Y == 1 && game.FoodPos.X < 10 {
		game.UpdateFoodPos(width, height)
	}
}

func (game *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	game.Screen.SetStyle(defStyle)
	width, height := game.Screen.Size()
	game.SnakeBody.ResetPos(width, height)
	game.UpdateFoodPos(width, height)
	game.GameOver = false
	game.Score = 0
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		longerSnake := false
		game.Screen.Clear()

		if checkCollison(game.SnakeBody.Parts[len(game.SnakeBody.Parts)-1:], game.FoodPos) {
			game.UpdateFoodPos(width, height)
			longerSnake = true
			game.Score++
		}
		if checkCollison(game.SnakeBody.Parts[:len(game.SnakeBody.Parts)-1], game.SnakeBody.Parts[len(game.SnakeBody.Parts)-1]) {
			break
		}
		game.SnakeBody.Update(width, height, longerSnake)
		drawParts(game.Screen, game.SnakeBody.Parts, game.FoodPos, snakeStyle, defStyle)
		drawText(game.Screen, 1, 1, 8+len(strconv.Itoa(game.Score)), 1, "Score: "+strconv.Itoa(game.Score))
		time.Sleep(60 * time.Millisecond)
		game.Screen.Show()
	}
	game.GameOver = true
	drawText(game.Screen, width/2-20, height/2, width/2+20, height/2, "Game Over, Score: "+strconv.Itoa(game.Score)+", Play Again? y/n")
	game.Screen.Show()
}
