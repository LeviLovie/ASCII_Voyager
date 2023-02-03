package menu

import "github.com/LeviiLovie/ASCII_Voyager/foo"

// var Tasks = []string{"New Game", "Load Game", "Settings", "Credits", "Exit"}
var Tasks = []string{"Open Demo", "Credits", "Exit"}

func SetUp() {
	foo.SetTerminalSize(foo.Width, foo.Height)
	foo.ClearScreen()
	foo.NotVisibleCursor()
}

func DrawLogo() {
	foo.WriteTextOnCenter("      _         ______      ______   _____   _____          ____   ____     ___     ____  ____        _          ______    ________   _______     ", foo.Width, 5)
	foo.WriteTextOnCenter("     / \\      .' ____ \\   .' ___  | |_   _| |_   _|        |_  _| |_  _|  .'   `.  |_  _||_  _|      / \\       .' ___  |  |_   __  | |_   __ \\    ", foo.Width, 6)
	foo.WriteTextOnCenter("    / _ \\     | (___ \\_| / .'   \\_|   | |     | |            \\ \\   / /   /  .-.  \\   \\ \\  / /       / _ \\     / .'   \\_|    | |_ \\_|   | |__) |   ", foo.Width, 7)
	foo.WriteTextOnCenter("   / ___ \\     _.____`.  | |          | |     | |             \\ \\ / /    | |   | |    \\ \\/ /       / ___ \\    | |   ____    |  _| _    |  __ /    ", foo.Width, 8)
	foo.WriteTextOnCenter(" _/ /   \\ \\_  | \\____) | \\ `.___.'\\  _| |_   _| |_             \\ ' /     \\  `-'  /    _|  |_     _/ /   \\ \\_  \\ `.___]  |  _| |__/ |  _| |  \\ \\_  ", foo.Width, 9)
	foo.WriteTextOnCenter("|____| |____|  \\______.'  `.____ .' |_____| |_____|             \\_/       `.___.'    |______|   |____| |____|  `._____.'  |________| |____| |___| ", foo.Width, 10)
	foo.WriteTextOnCenter("                                                                                                                                             ", foo.Width, 11)
}
