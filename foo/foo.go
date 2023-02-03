package foo

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

var Width = 175
var GameWidth = Width - (Width - 80)
var Height = 40

type KeyPress struct {
	Key  rune
	Code keyboard.Key
}

var TEXT_RED = "\033[38;5;196m"
var TEXT_CYAN = "\033[38;5;51m"
var TEXT_CYAN_LIGHT = "\033[38;5;87m"
var TEXT_WHITE_BOLD = "\033[38;5;0;48;5;255m"
var TEXT_RESET = "\033[38;5;255;48;5;0m"

func SetTerminalSize(x, y int) {
	fmt.Printf("\033[8;%d;%dt", y, x)
}

func ClearScreen() {
	fmt.Printf("\033[H\033[2J")
}

func WriteToLogFile(text string) {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(text)
}

func NotVisibleCursor() {
	fmt.Printf("\033[?25l")
}

func VisibleCursor() {
	fmt.Printf("\033[?25h")
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func DrawVerticalSplitLine(height int) {
	for y := 0; y < height-1; y++ {
		MoveCursor(80, y)
		fmt.Print("║")
	}
}

func WriteTextOnCenter(text string, width, y int) {
	var x = (width - len(text)) / 2
	MoveCursor(x, y)
	fmt.Print(text)
}

func SetColor(color string) {
	fmt.Printf("\033[%sm", color)
}

func ResetColor() {
	fmt.Printf("\033[0m")
}

func PrintAt(x, y int, text string) {
	MoveCursor(x, y)
	fmt.Print(text)
}

func PrintAtColor(x, y int, color string, text string) {
	MoveCursor(x, y)
	SetColor(color)
	fmt.Print(text)
	ResetColor()
}

func PrintAtColorBG(x, y int, color, bgcolor string, text string) {
	MoveCursor(x, y)
	SetColor(color)
	SetColor(bgcolor)
	fmt.Print(text)
	ResetColor()
}
