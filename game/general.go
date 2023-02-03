package game

import (
	"fmt"
	"time"

	"github.com/LeviiLovie/ASCII_Voyager/state"
)

func General(FPS int) state.State {
	// foo.ClearScreen()

	// keyCh := make(chan keyboard.KeyPress)
	// exit := make(chan bool)
	// defer func() {
	// 	exit <- true
	// }()

	// go func() {
	// 	err := keyboardAPI.Open()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer keyboardAPI.Close()

	// 	for {
	// 		select {
	// 		case <-exit:
	// 			return
	// 		default:
	// 			key, keyCode, err := keyboardAPI.GetKey()
	// 			if err != nil {
	// 				fmt.Println(err)
	// 				return
	// 			}
	// 			keyCh <- keyboard.KeyPress{
	// 				Key:  key,
	// 				Code: keyCode,
	// 			}
	// 		}
	// 	}
	// }()

	fmt.Println("Game 3")
	// for {
	// Redraw()

	// select {
	// case key := <-keyCh:
	// 	// fmt.Printf("Key pressed: %x\n", key)
	// 	switch true {
	// 	case key.Code == keyboardAPI.KeyArrowUp || key.Key == 'w' || key.Key == 'W':
	// 		if chose > 0 {
	// 			chose--
	// 		} else {
	// 			chose = len(menu.Tasks) - 1
	// 		}
	// 	case key.Code == keyboardAPI.KeyArrowDown || key.Key == 's' || key.Key == 'S':
	// 		if chose < len(menu.Tasks)-1 {
	// 			chose++
	// 		} else {
	// 			chose = 0
	// 		}
	// 	case key.Code == keyboardAPI.KeyEnter:
	// 		switch chose {
	// 		case len(menu.Tasks) - 1:
	// 			return state.Exit
	// 		case 0:
	// 			return state.Game
	// 		}
	// 	}

	// 	time.Sleep(time.Duration(FPS) * time.Millisecond)
	// }
	// }

	time.Sleep(1 * time.Second)
	return state.Exit
}
