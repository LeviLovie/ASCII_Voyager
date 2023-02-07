package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"

	"github.com/LeviiLovie/ASCII_Voyager/editor"
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
	if len(os.Args) > 1 {
		if os.Args[1] == "--editor" {
			editor.Editor()
		}
	} else {
		var playMusic bool = true
		foo.InitLog()
		logrus.Debug("")
		logrus.Debug("")
		logrus.Debug("")
		logrus.Debug("Starting main.go")

		if playMusic {
			go music.Init()
		}
		// json.NewSave("test")

		defer func() {
			if r := recover(); r != nil {
				logrus.Fatalf("Panic: %v", r)
				os.Exit(1)
			}
		}()

		var keys = make(chan foo.KeyPress, 32)
		var stage = foo.StageMenu

		foo.SetTerminalSize(foo.TerminalWidth, foo.TerminalHeight)
		foo.ClearScreen()
		foo.NotVisibleCursor()
		defer foo.VisibleCursor()

		go keyBoardRead(keys)
		logrus.Debugf("Started - keyBoardRead()")

		var (
			gameName    string
			gameNameNew string
			save        foo.GameWorld
			saveNew     foo.GameWorld
		)
		for {
			logrus.Debugf("Stage %d", stage)
			switch stage {
			case foo.StageMenu:
				stage = menu.Menu(FPS, keys)
			case foo.StageNewGame:
				foo.ClearScreen()
				foo.MenuDrawLogo()
				foo.MoveCursor(15, 15)
				fmt.Print("Enter game name: ")
				foo.MoveCursor(15, 16)
				gameName = foo.GetString(keys)
				json.NewSave(gameName)
				save = json.LoadSave(gameName)
				stage, gameNameNew, saveNew = game.Game(FPS, keys, save, gameName)
				json.SaveGame(gameNameNew, saveNew)
			case foo.StageLoadGame:
				foo.ClearScreen()
				foo.MenuDrawLogo()
				foo.MoveCursor(15, 15)
				fmt.Print("Enter game name: ")
				foo.MoveCursor(15, 16)
				gameName = foo.GetString(keys)
				logrus.Infof("Loading game: '%s'", gameName)
				save = json.LoadSave(gameName)
				stage, gameNameNew, saveNew = game.Game(FPS, keys, save, gameName)
				json.SaveGame(gameNameNew, saveNew)
			case foo.StageExit:
				logrus.Debugf("Exiting - main.go")
				foo.ClearScreen()
				foo.MoveCursor(0, 0)
				fmt.Println("Goodbye!")
				foo.SetTerminalSize(125, 25)
				return
			}
		}
	}
}
