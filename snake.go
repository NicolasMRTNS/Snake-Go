package main

type SnakeBody struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (snakeBody *SnakeBody) ChangeDir(vertical, horizontal int) {
	snakeBody.Yspeed = vertical
	snakeBody.Xspeed = horizontal
}

func (snakeBody *SnakeBody) Update(width, height int) {
	snakeBody.X = (snakeBody.X + snakeBody.Xspeed) % width
	if snakeBody.X < 0 {
		snakeBody.X += width
	}

	snakeBody.Y = (snakeBody.Y + snakeBody.Yspeed) % height
	if snakeBody.Y < 0 {
		snakeBody.Y += height
	}
}
