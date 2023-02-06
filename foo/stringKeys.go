package foo

import (
	"os"

	"github.com/eiannone/keyboard"
)

func GetString(keyChan chan KeyPress) string {
	var result string

	for {
		var key KeyPress

		select {
		case key = <-keyChan:
		default:
		}

		if key.Key == keyboard.KeyBackspace || key.Key == keyboard.KeyArrowLeft {
			if result != "" {
				result = result[:len(result)-1]
				print("\b \b")
			}
		} else if key.Key == keyboard.KeyEnter {
			return result
		} else if key.Key == keyboard.KeyEsc {
			os.Exit(0)
		} else if key.Char != 0 {
			result += string(key.Char)
			print(string(key.Char))
		}
	}
}
