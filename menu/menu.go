package menu

import (
	"github.com/LeviiLovie/ASCII_Voyager/foo"
)

var chose = 0

func Menu(keys chan foo.KeyPress) {
	foo.WriteToLogFile("Starting, menu/menu.go")

	foo.ClearScreen()
	foo.NotVisibleCursor()
	defer foo.VisibleCursor()
	foo.MenuDrawLogo()
	foo.WriteToLogFile("Menu - Done - ClearScreen, NotVisibleCursor, DrawLogo")

	foo.WriteToLogFile("Menu - Main loop starting")
	for {
		foo.MenuDrawTasks(0, 15, 15)
	}
}

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	foo "github.com/LeviiLovie/ASCII_Voyager/foo"
// 	keyboard "github.com/LeviiLovie/ASCII_Voyager/foo/keyBoard"
// 	menu "github.com/LeviiLovie/ASCII_Voyager/foo/menu"
// 	"github.com/LeviiLovie/ASCII_Voyager/state"
// 	keyboardAPI "github.com/eiannone/keyboard"
// )

// var chose = 0
// var FPS = 24

// func Redraw() {
// 	log.Println()
// 	log.Println("Redraw - ClearScreen, NotVisibleCursor, DrawLogo: ")
// 	foo.ClearScreen()
// 	foo.NotVisibleCursor()
// 	menu.DrawLogo()
// 	log.Printf("done")

// 	log.Println("Redraw - DrawTasks: ")
// 	for i := 0; i < len(menu.Tasks); i++ {
// 		if chose == i {
// 			foo.MoveCursor(15, 15+i)
// 			fmt.Printf(foo.BOLD_TEXT)
// 			fmt.Printf(menu.Tasks[i])
// 			fmt.Printf(foo.RESET_TEXT)
// 		} else {
// 			foo.MoveCursor(15, 15+i)
// 			fmt.Printf(menu.Tasks[i])
// 		}
// 	}
// 	log.Printf("done")
// }

// func Menu() state.State {
// 	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("error opening file: %v", err)
// 	}
// 	defer f.Close()
// 	log.SetOutput(f)
// 	log.Println()

// 	log.Println("45, Menu - Starting")
// 	log.Println("46, Menu - SetUp, DrawLogo: ")
// 	menu.SetUp()
// 	menu.DrawLogo()
// 	log.Printf("done")

// 	log.Println("51, Menu - Opening chanels: ")
// 	keyCh := make(chan keyboard.KeyPress)
// 	exit := make(chan bool)
// 	log.Printf("done")

// 	log.Println("56, Menu - defer exit function: ")
// 	defer func() {
// 		exit <- true
// 		fmt.Println("exit: ", <-exit)
// 	}()
// 	log.Printf("done")

// 	log.Println("64, Menu - go func")
// 	go func() {
// 		err := keyboardAPI.Open()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer keyboardAPI.Close()

// 		for {
// 			if <-exit == true {
// 				return
// 			} else {
// 				key, keyCode, err := keyboardAPI.GetKey()
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 				keyCh <- keyboard.KeyPress{
// 					Key:  key,
// 					Code: keyCode,
// 				}
// 			}
// 			// select {
// 			// case <-exit:
// 			// 	keyboardAPI.Close()
// 			// 	fmt.Println("return exit")
// 			// 	return
// 			// default:
// 			// 	key, keyCode, err := keyboardAPI.GetKey()
// 			// 	if err != nil {
// 			// 		fmt.Println(err)
// 			// 		return
// 			// 	}
// 			// 	keyCh <- keyboard.KeyPress{
// 			// 		Key:  key,
// 			// 		Code: keyCode,
// 			// 	}
// 			// }
// 		}
// 	}()

// 	log.Println("104, Menu - for loop")
// 	for {
// 		log.Println("106, Menu - Redraw: ")
// 		Redraw()
// 		log.Printf("done")
// 		// log.Printf("commented")

// 		log.Println("119, Menu - select (keyCh): ", keyCh)
// 		var key keyboard.KeyPress
// 		key = <-keyCh
// 		keyCode := key.Code
// 		log.Println("120, Menu - select (keyCh.Code): ", *keyCode)
// 		select {
// 		case key := <-keyCh:
// 			// fmt.Printf("Key pressed: %x\n", key)
// 			log.Println("123, Menu - select (keyCh): ", key)
// 			switch true {
// 			case key.Code == keyboardAPI.KeyArrowUp || key.Key == 'w' || key.Key == 'W':
// 				log.Println("122, W - switch")
// 				if chose > 0 {
// 					chose--
// 				} else {
// 					chose = len(menu.Tasks) - 1
// 				}
// 				log.Printf("done")
// 			case key.Code == keyboardAPI.KeyArrowDown || key.Key == 's' || key.Key == 'S':
// 				log.Println("130, S - switch")
// 				if chose < len(menu.Tasks)-1 {
// 					chose++
// 				} else {
// 					chose = 0
// 				}
// 				log.Printf("done")
// 			case key.Code == keyboardAPI.KeyEnter:
// 				log.Println("138, Enter - switch")
// 				switch chose {
// 				case len(menu.Tasks) - 1:
// 					return state.Exit
// 				case 0:
// 					fmt.Println("Game 1")
// 					return state.Game
// 				}
// 				log.Printf("done")
// 			}

// 			log.Println("151, Menu - select (time.After): ")
// 			time.Sleep(time.Duration(FPS) * time.Millisecond)
// 			log.Printf("done")
// 		}
// 	}
// }
