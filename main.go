package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/LeviiLovie/ASCII_Voyager/game"
	"github.com/LeviiLovie/ASCII_Voyager/menu"
	"github.com/LeviiLovie/ASCII_Voyager/state"
)

const FPS = 30

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println()
	log.Println()
	log.Println()

	log.Println("Main - Starting")
	var currentState = state.Menu
	log.Println("Main - CurrentState: ", currentState)
	for {
		switch currentState {
		case state.Menu:
			log.Println("Main - Menu case")
			currentState = menu.Menu()
		case state.Game:
			log.Println("Main - Game case")
			fmt.Println("Game 2")
			currentState = game.General(FPS)
		case state.Exit:
			log.Println("Main - Exit case")
			foo.ClearScreen()
			foo.VisibleCursor()
			fmt.Printf("Goodbye!\n")
			return
		}
	}
}
