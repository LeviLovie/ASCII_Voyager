package foo

import "fmt"

var MenuTasks = []string{"Play", "Credits", "Exit"}
var ManuCredits = []string{"LeviLovie", "Oto", "Learunaso"}

func MenuDrawTasks(chose, x, y int) {
	for i, task := range MenuTasks {
		MoveCursor(x, y+i)
		if chose == i {
			fmt.Printf(TEXT_WHITE_BOLD + task + TEXT_RESET)
		} else {
			fmt.Print(task)
		}
	}
}

func MenuDrawCredits(x, y int, prefix string) {
	for i, credit := range ManuCredits {
		MoveCursor(x, y+i)
		fmt.Printf(prefix, credit)
		fmt.Printf(prefix, credit)
	}
}

func MenuSetUp() {
	SetTerminalSize(Width, Height)
	ClearScreen()
	NotVisibleCursor()
}

func MenuDrawLogo() {
	WriteTextOnCenter("      _         ______      ______   _____   _____          ____   ____     ___     ____  ____        _          ______    ________   _______     ", Width, 5)
	WriteTextOnCenter("     / \\      .' ____ \\   .' ___  | |_   _| |_   _|        |_  _| |_  _|  .'   `.  |_  _||_  _|      / \\       .' ___  |  |_   __  | |_   __ \\    ", Width, 6)
	WriteTextOnCenter("    / _ \\     | (___ \\_| / .'   \\_|   | |     | |            \\ \\   / /   /  .-.  \\   \\ \\  / /       / _ \\     / .'   \\_|    | |_ \\_|   | |__) |   ", Width, 7)
	WriteTextOnCenter("   / ___ \\     _.____`.  | |          | |     | |             \\ \\ / /    | |   | |    \\ \\/ /       / ___ \\    | |   ____    |  _| _    |  __ /    ", Width, 8)
	WriteTextOnCenter(" _/ /   \\ \\_  | \\____) | \\ `.___.'\\  _| |_   _| |_             \\ ' /     \\  `-'  /    _|  |_     _/ /   \\ \\_  \\ `.___]  |  _| |__/ |  _| |  \\ \\_  ", Width, 9)
	WriteTextOnCenter("|____| |____|  \\______.'  `.____ .' |_____| |_____|             \\_/       `.___.'    |______|   |____| |____|  `._____.'  |________| |____| |___| ", Width, 10)
	WriteTextOnCenter("                                                                                                                                             ", Width, 11)
}
