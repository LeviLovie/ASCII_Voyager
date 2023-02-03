package main

import (
	"fmt"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/LeviiLovie/ASCII_Voyager/menu"
	"github.com/eiannone/keyboard"
)

const FPS = 3

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
	foo.WriteToLogFile("")
	foo.WriteToLogFile("Starting - main.go")

	var keys = make(chan foo.KeyPress, 32)
	var stage = 1

	foo.SetTerminalSize(175, 30)
	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()
	foo.WriteToLogFile("Done - ClearScreen, InvisibleCursor")

	go keyBoardRead(keys)
	foo.WriteToLogFile("Started - keyBoardRead()")

	// for {
	// 	for val := range keys {
	// 		switch val.Code {
	// 		case keyboard.KeyEsc:
	// 			foo.WriteToLogFile("Exiting - main.go")
	// 			foo.ClearScreen()
	// 			foo.MoveCursor(0, 0)
	// 			fmt.Println("Goodbye!")
	// 			return
	// 		}
	// 	}
	// }

	switch stage {
	case 1:
		stage = menu.Menu(keys)

	case 0:
		foo.WriteToLogFile("Exiting - main.go")
		foo.ClearScreen()
		foo.MoveCursor(0, 0)
		fmt.Println("Goodbye!")
		return
	}

	// log.Println("Main - Starting")
	// var currentState = state.Menu
	// log.Println("Main - CurrentState: ", currentState)
	// for {
	// 	switch currentState {
	// 	case state.Menu:
	// 		log.Println("Main - Menu case")
	// 		// currentState = menu.Menu()
	// 	case state.Game:
	// 		log.Println("Main - Game case")
	// 		fmt.Println("Game 2")
	// 		currentState = game.General(FPS)
	// 	case state.Exit:
	// 		log.Println("Main - Exit case")
	// 		foo.ClearScreen()
	// 		foo.VisibleCursor()
	// 		fmt.Printf("Goodbye!\n")
	// 		return
	// 	}
	// }
}
