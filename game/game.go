package game

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
)

func keyBoard(keyPress foo.KeyPress, world *foo.GameWorld) int {
	switch keyPress.Key {
	case keyboard.KeyEsc:
		return 0
	}

	switch keyPress.Char {
	case 'w', 'W':
		world.MovePlayerUp()
		world.NeedRedraw = true
	case 's', 'S':
		world.MovePlayerDown()
		world.NeedRedraw = true
	case 'a', 'A':
		world.MovePlayerLeft()
		world.NeedRedraw = true
	case 'd', 'D':
		world.MovePlayerRight()
		world.NeedRedraw = true
	}

	return 1
}

func Game(FPS int, keys chan foo.KeyPress, save foo.GameWorld, gameName string) (int, string, foo.GameWorld) {
	logrus.Debugf("Starting - Game.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

	var world *foo.GameWorld = &save
	logrus.Debugf("Game - save = '%s'", save)

	logrus.Debugf("Game - Done - Init world")
	logrus.Infof("Game - World size: %dx%d", world.Width, world.Height)
	logrus.Infof("Game - Player position: %dx%d", world.PlayerPositionX, world.PlayerPositionY)

	logrus.Debugf("Game - Main loop starting")
	world.NeedRedraw = true
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
			return 1, gameName, *world
		}
		time.Sleep(time.Second / 30)
	}
}
