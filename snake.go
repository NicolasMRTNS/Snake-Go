package main

type Part struct {
	X int
	Y int
}

type SnakeBody struct {
	Parts  []Part
	Xspeed int
	Yspeed int
}

func (snakeBody *SnakeBody) ChangeDir(vertical, horizontal int) {
	snakeBody.Yspeed = vertical
	snakeBody.Xspeed = horizontal
}

func (snakeBody *SnakeBody) Update(width, height int, longerSnake bool) {
	snakeBody.Parts = append(snakeBody.Parts, snakeBody.Parts[len(snakeBody.Parts)-1].GetUpdatedPart(snakeBody, width, height))
	if !longerSnake {
		snakeBody.Parts = snakeBody.Parts[1:]
	}
}

func (snakeBody *SnakeBody) ResetPos(width, height int) {
	snakeParts := []Part{
		{
			X: int(width / 2),
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 1,
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 2,
			Y: int(height / 2),
		},
	}
	snakeBody.Parts = snakeParts
	snakeBody.Xspeed = 1
	snakeBody.Yspeed = 0
}

func (snakePart *Part) GetUpdatedPart(snakeBody *SnakeBody, width, height int) Part {
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
