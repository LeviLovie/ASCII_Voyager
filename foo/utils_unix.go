//go:build !windows
// +build !windows

package foo

import "fmt"

var (
	TEXT_BLACK         = "\033[38;5;0m"
	TEXT_RED           = "\033[38;5;1m"
	TEXT_GREEN         = "\033[38;5;2m"
	TEXT_YELLOW        = "\033[38;5;3m"
	TEXT_BLUE          = "\033[38;5;4m"
	TEXT_MAGENTA       = "\033[38;5;5m"
	TEXT_CYAN          = "\033[38;5;6m"
	TEXT_WHITE         = "\033[38;5;7m"
	TEXT_LIGHT_BLACK   = "\033[38;5;8m"
	TEXT_LIGHT_RED     = "\033[38;5;9m"
	TEXT_LIGHT_GREEN   = "\033[38;5;10m"
	TEXT_LIGHT_YELLOW  = "\033[38;5;11m"
	TEXT_LIGHT_BLUE    = "\033[38;5;12m"
	TEXT_LIGHT_MAGENTA = "\033[38;5;13m"
	TEXT_LIGHT_CYAN    = "\033[38;5;14m"
	TEXT_LIGHT_WHITE   = "\033[38;5;15m"
	TEXT_LIGHT_ORANGE  = "\033[38;5;214m"
	TEXT_WHITE_BOLD    = "\033[38;5;0;48;5;255m"
	TEXT_RESET         = "\033[0m"
)

func SetTerminalSize(x, y int) {
	fmt.Printf("\033[8;%d;%dt", y, x)
}

func ClearScreen() {
	fmt.Printf("\033[H\033[2J\033[3J")
}

func NotVisibleCursor() {
	fmt.Printf("\033[?25l")
}

func VisibleCursor() {
	fmt.Printf("\033[?25h")
}

func MoveCursor(x, y int) {
	if x < 0 {
		return
	}
	if y < 0 {
		return
	}
	fmt.Printf("\033[%d;%dH", y, x)
}

func PrintAt(x, y int, text string, args ...interface{}) {
	if x < 0 {
		return
	}
	if y < 0 {
		return
	}
	fmt.Printf("\033[%d;%dH%s", y, x, fmt.Sprintf(text, args...))
}
