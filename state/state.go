package state

type State int

const (
	Menu State = iota
	Game
	Exit
)
