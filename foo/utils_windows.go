//go:build windows
// +build windows

package foo

import (
	"github.com/AllenDang/w32"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var (
	kernel32             = windows.NewLazySystemDLL("kernel32.dll")
	setConsoleCursorInfo = kernel32.NewProc("SetConsoleCursorInfo")
)

var (
	TEXT_BLACK         = func() string { return colorText(0, 0) }
	TEXT_RED           = func() string { return colorText(0, 1) }
	TEXT_GREEN         = func() string { return colorText(0, 2) }
	TEXT_YELLOW        = func() string { return colorText(0, 3) }
	TEXT_BLUE          = func() string { return colorText(0, 4) }
	TEXT_MAGENTA       = func() string { return colorText(0, 5) }
	TEXT_CYAN          = func() string { return colorText(0, 6) }
	TEXT_WHITE         = func() string { return colorText(0, 7) }
	TEXT_LIGHT_BLACK   = func() string { return colorText(1, 0) }
	TEXT_LIGHT_RED     = func() string { return colorText(1, 1) }
	TEXT_LIGHT_GREEN   = func() string { return colorText(1, 2) }
	TEXT_LIGHT_YELLOW  = func() string { return colorText(1, 3) }
	TEXT_LIGHT_BLUE    = func() string { return colorText(1, 4) }
	TEXT_LIGHT_MAGENTA = func() string { return colorText(1, 5) }
	TEXT_LIGHT_CYAN    = func() string { return colorText(1, 6) }
	TEXT_LIGHT_WHITE   = func() string { return colorText(1, 7) }
	TEXT_LIGHT_ORANGE  = func() string { return colorText(1, 6) }
	TEXT_WHITE_BOLD    = func() string { return colorText(1, 7) }
	TEXT_RESET         = ""
)

func SetTerminalSize(x, y int) {
	console := w32.GetConsoleWindow()
	w32.MoveWindow(console, 100, 100, x, y, true)
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func NotVisibleCursor() {
	w32.ShowCursor(false)
}

func VisibleCursor() {
	w32.ShowCursor(true)
}

func MoveCursor(x, y int) {
	if x < 0 {
		return
	}
	if y < 0 {
		return
	}
	w32.SetConsoleCursorPosition(w32.GetStdHandle(w32.STD_OUTPUT_HANDLE), w32.COORD{X: int16(x), Y: int16(y)})
}

func PrintAt(x, y int, text string, args ...interface{}) {
	if x < 0 {
		return
	}
	if y < 0 {
		return
	}
	MoveCursor(x, y)
	fmt.Printf(text, args...)
}
