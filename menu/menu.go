package menu

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
)

var chose = 0

func keyBoard(keyPress foo.KeyPress) int {
	switch keyPress.Key {
	case keyboard.KeyEnter:
		return chose
	case keyboard.KeyArrowUp:
		if chose > 0 {
			chose--
		} else {
			chose = len(foo.MenuTasks) - 1
		}
	case keyboard.KeyArrowDown:
		if chose < len(foo.MenuTasks)-1 {
			chose++
		} else {
			chose = 0
		}
	}

	switch keyPress.Char {
	case 'w':
		if chose > 0 {
			chose--
		} else {
			chose = len(foo.MenuTasks) - 1
		}
	case 'W':
		if chose > 0 {
			chose--
		} else {
			chose = len(foo.MenuTasks) - 1
		}
	case 's':
		if chose < len(foo.MenuTasks)-1 {
			chose++
		} else {
			chose = 0
		}
	case 'S':
		if chose < len(foo.MenuTasks)-1 {
			chose++
		} else {
			chose = 0
		}
	}
	return 1
}

func Menu(FPS int, keys chan foo.KeyPress) int {
	logrus.Debugf("Starting, menu/menu.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

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

		switch keyBoard(keyPress) {
		case 0:
			logrus.Debugf("Menu - Starting - Game")
			return 2
		case 2:
			logrus.Debugf("Menu - Exit")
			foo.ClearScreen()
			foo.MoveCursor(0, 0)
			fmt.Println("Goodbye!")
			return 0
		}
		time.Sleep(time.Second / 30)
	}
}
