package cutscenes

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/eiannone/keyboard"
)

//go:embed startGame.dat
var startGame []byte

var cuts [][]string

func keyBoard(keyPress foo.KeyPress) int {
	switch keyPress.Key {
	case keyboard.KeyEnter:
		return 1
	case keyboard.KeyEsc:
		return 2
	}
	return 0
}

func SatrtGameCuts(keys chan foo.KeyPress) {
	scanner := bufio.NewScanner(bytes.NewReader(startGame))
	var i int
	var cut []string
	for scanner.Scan() {
		if i%30 == 0 && i != 0 {
			cuts = append(cuts, cut)
		} else {
			cut = append(cut, scanner.Text())
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var result int
	for {
		var keyPress foo.KeyPress
		select {
		case keyPress = <-keys:
		default:
		}

		foo.ClearScreen()
		for _, line := range cuts[0] {
			fmt.Println(line)
		}
		cuts = cuts[1:]

		result = keyBoard(keyPress)

		switch result {
		case 1:
			fmt.Println("Start game")
		case 2:
			fmt.Println("Exit game")
			os.Exit(0)
		}
		time.Sleep(time.Second / 30)
	}
}
