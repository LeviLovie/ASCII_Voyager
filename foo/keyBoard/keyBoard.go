package keyboard

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type KeyPress struct {
	Key  rune
	Code keyboard.Key
}

func Quit(width, height int) {
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
		if char == 'q' || key == keyboard.KeyEsc {
			break
		}
	}
}

func GetKeyCodePress() (res KeyPress) {
	err := keyboard.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer keyboard.Close()

	key, keyCode, err := keyboard.GetKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	return KeyPress{
		Key:  key,
		Code: keyCode,
	}
}
