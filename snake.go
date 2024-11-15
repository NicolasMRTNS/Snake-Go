package main

type SnakePart struct {
	X int
	Y int
}

type SnakeBody struct {
	Parts  []SnakePart
	Xspeed int
	Yspeed int
}

func (snakeBody *SnakeBody) ChangeDir(vertical, horizontal int) {
	snakeBody.Yspeed = vertical
	snakeBody.Xspeed = horizontal
}

func (snakeBody *SnakeBody) Update(width, height int) {
	snakeBody.Parts = append(snakeBody.Parts, snakeBody.Parts[len(snakeBody.Parts)-1].GetUpdatedPart(snakeBody, width, height))
	snakeBody.Parts = snakeBody.Parts[1:]
}

func (snakePart *SnakePart) GetUpdatedPart(snakeBody *SnakeBody, width, height int) SnakePart {
	newPart := *snakePart
	newPart.X = (newPart.X + snakeBody.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}
	newPart.Y = (newPart.Y + snakeBody.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}
	return newPart
}
