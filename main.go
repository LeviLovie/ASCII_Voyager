package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/LeviiLovie/ASCII_Voyager/game"
	"github.com/LeviiLovie/ASCII_Voyager/json"
	"github.com/LeviiLovie/ASCII_Voyager/menu"
	"github.com/LeviiLovie/ASCII_Voyager/music"
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
	go music.Init()
	// json.NewSave("test")

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

	var save game.GameWorld
	for {
		logrus.Debugf("Stage %d", stage)
		switch stage {
		case 1:
			stage = menu.Menu(FPS, keys)
		case 2:
			json.NewSave("test")
			save = json.LoadSave("test")
			stage = game.Game(FPS, keys, save)
		case 3:
			save = json.LoadSave("test")
			stage = game.Game(FPS, keys, save)
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
