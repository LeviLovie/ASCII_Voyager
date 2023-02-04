package game

import (
	"fmt"
	"time"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/eiannone/keyboard"
)

func keyBoard(keyPress foo.KeyPress, world *GameWorld) int {
	switch keyPress.Key {
	case keyboard.KeyEsc:
		return 0
	}

	switch keyPress.Char {
	case 'w', 'W':
		world.MovePlayerUp()
	case 's', 'S':
		world.MovePlayerDown()
	case 'a', 'A':
		world.MovePlayerLeft()
	case 'd', 'D':
		world.MovePlayerRight()
	}

	return 1
}

func Game(FPS int, keys chan foo.KeyPress) int {
	foo.WriteToLogFile("Starting - Game.go")
	fmt.Println("Game")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()
	foo.WriteToLogFile("Game - Done - ClearScreen, NotVisibleCursor, DrawLogo")

	foo.WriteToLogFile("Game - Init world")
	var world = &GameWorld{}
	world.Init()
	foo.WriteToLogFile("Game - Done - Init world")

	foo.WriteToLogFile("Game - Main loop starting")
	for {
		world.Draw()

		var keyPress foo.KeyPress
		select {
		case keyPress = <-keys:
		default:
		}

		result := keyBoard(keyPress, world)
		if result == 0 {
			foo.WriteToLogFile("Game - Exit")
			foo.ClearScreen()
			foo.MoveCursor(0, 0)
			fmt.Println("Goodbye!")
			foo.WriteToLogFile("Game - Send 1")
			return 1
		}
		time.Sleep(time.Second / 30)
	}
}
