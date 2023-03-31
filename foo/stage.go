package foo

type Stage int

const (
	StageMenu Stage = iota
	StageNewGame
	StageLoadGame
	StageCreditsGame
	StageExit
)
