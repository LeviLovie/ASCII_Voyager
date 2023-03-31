package main

import (
	"flag"
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
	var (
		runEditor bool
		noMusic   bool
		help      bool
	)

	flag.BoolVar(&runEditor, "editor", false, "Start editor")
	flag.BoolVar(&noMusic, "no-music", false, "Disable music")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.Parse()

	if help {
		fmt.Println("ASCII Voyager")
		fmt.Println("  -editor\tStart editor")
		fmt.Println("  -no-music\tDisable music")
		fmt.Println("  -help\t\tShow help")
		return
	}

	if runEditor {
		editor.Editor()
		return
	}

	foo.InitLog()
	logrus.Debug("")
	logrus.Debug("")
	logrus.Debug("")
	logrus.Debug("Starting main.go")

	if !noMusic {
		go music.Init()
	}

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
		gameName string
		result   int
		save     foo.GameWorld
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
			json.NewSave(gameName, keys)
			result, save = json.LoadSave(gameName)
			if result == 0 {
				stage, _, _ = game.Game(FPS, keys, save, gameName)
			} else {
				stage = foo.StageMenu
			}
		case foo.StageLoadGame:
			foo.ClearScreen()
			foo.MenuDrawLogo()
			foo.MoveCursor(15, 15)
			fmt.Print("Enter game name: ")
			foo.MoveCursor(15, 16)
			gameName = foo.GetString(keys)
			logrus.Infof("Loading game: '%s'", gameName)
			result, save = json.LoadSave(gameName)
			if result == 0 {
				stage, _, _ = game.Game(FPS, keys, save, gameName)
			} else {
				stage = foo.StageMenu
			}
		case foo.StageCreditsGame:
			foo.ClearScreen()
			foo.MoveCursor(15, 15)
			foo.NotVisibleCursor()
			foo.WriteTextOnCenter("Credits", foo.TerminalWidth, 4)

			foo.WriteTextOnCenter("Game Developing", foo.TerminalWidth, 6)
			foo.WriteTextOnCenter("LeviiLovie - https://github.com/LeviiLovie", foo.TerminalWidth, 7)
			foo.WriteTextOnCenter("Dreadatour - https://github.com/dreadatour", foo.TerminalWidth, 8)

			foo.WriteTextOnCenter("Game Desing", foo.TerminalWidth, 10)
			foo.WriteTextOnCenter("Oto - ", foo.TerminalWidth, 11)

			foo.WriteTextOnCenter("Testing", foo.TerminalWidth, 13)
			foo.WriteTextOnCenter("Dreadatour - https://github.com/dreadatour", foo.TerminalWidth, 14)

			foo.WriteTextOnCenter("Music", foo.TerminalWidth, 16)
			foo.WriteTextOnCenter("Learunaso - https://on.soundcloud.com/v7Wkeph5BNPTfzT96", foo.TerminalWidth, 17)

			foo.WriteTextOnCenter("The main motivator", foo.TerminalWidth, 19)
			foo.WriteTextOnCenter("Hamster of the main developer", foo.TerminalWidth, 20)

			foo.WriteTextOnCenter("Source code", foo.TerminalWidth, 22)
			foo.WriteTextOnCenter("https://github.com/LeviiLovie/ASCII_Voyager", foo.TerminalWidth, 23)

			<-keys
			stage = foo.StageMenu
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
