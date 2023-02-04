package menu

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
)

func keyBoard(chose foo.MenuItem, keyPress foo.KeyPress) (foo.MenuItem, bool) {
	switch keyPress.Key {
	case keyboard.KeyEnter:
		return chose, true
	case keyboard.KeyArrowUp:
		return chose.Prev(), false
	case keyboard.KeyArrowDown:
		return chose.Next(), false
	}

	switch keyPress.Char {
	case 'w', 'W':
		return chose.Prev(), false
	case 's', 'S':
		return chose.Next(), false
	}

	return chose, false
}

func Menu(FPS int, keys chan foo.KeyPress) int {
	logrus.Debugf("Starting, menu/menu.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

	var (
		chose    = foo.MenuItemNewGame
		selected = false
	)

	logrus.Debugf("Menu - Main loop starting")
	for {
		foo.ClearScreen()
		foo.MenuDrawLogo()
		foo.MenuDrawTasks(chose, 15, 15)

		var keyPress foo.KeyPress
		select {
		case keyPress = <-keys:
		default:
		}

		chose, selected = keyBoard(chose, keyPress)

		if selected {
			switch chose {
			case foo.MenuItemNewGame:
				logrus.Debugf("Menu - Starting - New Game")
				return 2
			case foo.MenuItemLoadGame:
				logrus.Debugf("Menu - Starting - Load Game")
				return 3
			case foo.MenuItemExit:
				logrus.Debugf("Menu - Exit")
				foo.ClearScreen()
				foo.MoveCursor(0, 0)
				fmt.Println("Goodbye!")
				return 0
			}
		}
		time.Sleep(time.Second / 30)
	}
}
