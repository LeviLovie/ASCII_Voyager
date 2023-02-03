package menu

import (
	"fmt"
	"time"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/eiannone/keyboard"
)

var chose = 0

func KeyBoard(keys chan foo.KeyPress) int {
	var keyPress foo.KeyPress
	select {
	case keyPress = <-keys:
	}

	switch keyPress.Key {
	case keyboard.KeyEnter:
		return 0
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

func Menu(keys chan foo.KeyPress) int {
	foo.WriteToLogFile("Starting, menu/menu.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()
	foo.WriteToLogFile("Menu - Done - ClearScreen, NotVisibleCursor, DrawLogo")

	foo.WriteToLogFile("Menu - Main loop starting")
	var result int
	for {
		foo.ClearScreen()
		foo.MenuDrawLogo()
		foo.MenuDrawTasks(chose, 15, 15)

		result = KeyBoard(keys)
		if result == 0 {
			switch chose {
			case len(foo.MenuTasks) - 1:
				foo.WriteToLogFile("Menu - Exit")
				foo.ClearScreen()
				foo.MoveCursor(0, 0)
				fmt.Println("Goodbye!")
				return 0
			}
		}
		time.Sleep(time.Second / 30)
	}
}
