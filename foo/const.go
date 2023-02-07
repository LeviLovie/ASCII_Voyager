package foo

const (
	TerminalWidth  = 175
	TerminalHeight = 30

	GameMenuWidth = 50

	GameLeft   = 0
	GameRight  = 0
	GameTop    = 0
	GameBottom = 0

	GameWidth  = TerminalWidth - GameMenuWidth - GameLeft - GameRight
	GameHeight = TerminalHeight - GameTop - GameBottom
)
