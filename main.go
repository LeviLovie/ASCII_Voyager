package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/LeviiLovie/ASCII_Voyager/game"
	"github.com/LeviiLovie/ASCII_Voyager/menu"
)

const FPS = 30

func keyBoardRead(keys chan foo.KeyPress) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		keys <- foo.KeyPress{
			Char: char,
			Key:  key,
		}
	}
}

func main() {
	foo.InitLog()
	logrus.Debug("")
	logrus.Debug("")
	logrus.Debug("")
	logrus.Debug("Starting main.go")

	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("Panic: %v", r)
			os.Exit(1)
		}
	}()

	var keys = make(chan foo.KeyPress, 32)
	var stage = 1

	foo.SetTerminalSize(175, 30)
	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()

	go keyBoardRead(keys)
	logrus.Debugf("Started - keyBoardRead()")

	for {
		logrus.Debugf("Stage %d", stage)
		switch stage {
		case 1:
			stage = menu.Menu(FPS, keys)
		case 2:
			stage = game.Game(FPS, keys)
		case 0:
			logrus.Debugf("Exiting - main.go")
			foo.ClearScreen()
			foo.MoveCursor(0, 0)
			fmt.Println("Goodbye!")
			foo.SetTerminalSize(125, 25)
			return
		}
	}
}
