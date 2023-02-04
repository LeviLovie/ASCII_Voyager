package game

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
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

func Game(FPS int, keys chan foo.KeyPress, save GameWorld) int {
	logrus.Debugf("Starting - Game.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

	var world *GameWorld = &save

	world.Init()
	logrus.Debugf("Game - Done - Init world")
	logrus.Infof("Game - World size: %dx%d", world.Width, world.Height)
	logrus.Infof("Game - Player position: %dx%d", world.PlayerPositionX, world.PlayerPositionY)

	logrus.Debugf("Game - Main loop starting")
	for {
		world.Draw()

		var keyPress foo.KeyPress
		select {
		case keyPress = <-keys:
		default:
		}

		result := keyBoard(keyPress, world)
		if result == 0 {
			logrus.Debugf("Game - Exit")
			foo.ClearScreen()
			foo.MoveCursor(0, 0)
			fmt.Println("Goodbye!")
			return 1
		}
		time.Sleep(time.Second / 30)
	}
}
