package game

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/LeviiLovie/ASCII_Voyager/json"
)

func keyBoard(keyPress foo.KeyPress, world *foo.GameWorld) int {
	switch keyPress.Key {
	case keyboard.KeyEsc:
		return 0
	case keyboard.KeyTab:
		return 1
	case keyboard.KeyF3:
		return -3
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

	return -1
}

func Game(FPS int, keys chan foo.KeyPress, save foo.GameWorld, gameName string) (foo.Stage, string, foo.GameWorld) {
	logrus.Debugf("Starting - Game.go")
	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

	var world *foo.GameWorld = &save
	json.CheckVersions(world)

	var (
		isPauseMenu bool = false
		isDevMode   bool = false
	)

	logrus.Debugf("Game - Main loop starting")
	world.NeedRedraw = true
	for {
		world.DrawGame()
		world.DrawMenu(isDevMode, isPauseMenu)

		var keyPress foo.KeyPress
		select {
		case keyPress = <-keys:
		default:
		}

		result := keyBoard(keyPress, world)
		if result == 0 {
			if isPauseMenu {
				logrus.Debugf("Game - Exit")
				foo.ClearScreen()
				foo.MoveCursor(0, 0)
				fmt.Println("Goodbye!")
				return foo.StageMenu, gameName, *world
			}
		}
		if result == 1 {
			foo.ClearScreen()
			world.NeedRedraw = true
			isPauseMenu = !isPauseMenu
		}
		if result == -3 {
			foo.ClearScreen()
			world.NeedRedraw = true
			isDevMode = !isDevMode
		}
		time.Sleep(time.Second / 30)
	}
}
