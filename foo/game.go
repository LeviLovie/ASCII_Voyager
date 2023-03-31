package foo

import (
	"strconv"
)

type GameWorld struct {
	FPS        int      `json:"fps"`
	NeedRedraw bool     `json:"needRedraw"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	World      [][]int  `json:"world"`
	Player     Player   `json:"player"`
	Version    float32  `json:"save_version"`
	Blocks     []string `json:"blocks"`
}

func (g *GameWorld) DrawGame() {
	if !g.NeedRedraw {
		return
	}

	ClearScreen()

	screenPlayerX := GameWidth/2 + GameLeft
	screenPlayerY := GameHeight/2 + GameTop

	screenWorldX := screenPlayerX - g.Player.X
	screenWorldY := screenPlayerY - g.Player.Y

	for i := GameLeft; i < GameLeft+GameWidth; i++ {
		for j := GameTop; j < GameTop+GameHeight; j++ {
			if j < screenWorldY || j >= screenWorldY+g.Height || i < screenWorldX || i >= screenWorldX+g.Width {
				continue
			}
			MoveCursor(i, j)

			for k := 0; k < len(g.Blocks); k++ {
				if g.World[j-screenWorldY][i-screenWorldX] == k {
					PrintAt(i, j, g.Blocks[k])
				}
			}
		}
	}

	PrintAt(screenPlayerX, screenPlayerY, TEXT_LIGHT_CYAN+"@"+TEXT_RESET)
	g.NeedRedraw = false
}

func (g *GameWorld) DrawMenu() {
	for i := 1; i < TerminalHeight+1; i++ {
		PrintAt(GameWidth+GameLeft+GameRight, i, "║")
	}
	PrintAt(GameWidth+GameLeft+GameRight+1+((GameMenuWidth-4)/2), 1, TEXT_LIGHT_GREEN+"MENU"+TEXT_RESET)
	if g.Player.HP > 66 {
		PrintAt(GameWidth+GameLeft+GameRight+3, 3, "Health: "+TEXT_GREEN+strconv.Itoa(g.Player.HP)+TEXT_RESET)
	} else if g.Player.HP > 33 {
		PrintAt(GameWidth+GameLeft+GameRight+3, 3, "Health: "+TEXT_YELLOW+strconv.Itoa(g.Player.HP)+TEXT_RESET)
	} else if g.Player.HP > 0 {
		PrintAt(GameWidth+GameLeft+GameRight+3, 3, "Health: "+TEXT_RED+strconv.Itoa(g.Player.HP)+TEXT_RESET)
	} else {
		PrintAt(GameWidth+GameLeft+GameRight+3, 3, "Health: "+TEXT_MAGENTA+strconv.Itoa(g.Player.HP)+TEXT_RESET)
	}
}

func (g *GameWorld) SetPlayerPosition(x, y int) {
	g.Player.X = x
	g.Player.Y = y
	g.NeedRedraw = true
}

func (g *GameWorld) MovePlayer(x, y int) {
	g.Player.X += x
	g.Player.Y += y
	g.NeedRedraw = true
}

func (g *GameWorld) MovePlayerUp() {
	if g.Player.Y < 0 {
		return
	}
	if g.Player.Y > 0 && g.World[g.Player.Y-1][g.Player.X] == 1 {
		return
	}
	g.SetPlayerPosition(g.Player.X, g.Player.Y-1)
}

func (g *GameWorld) MovePlayerDown() {
	if g.Player.Y > g.Height-1 {
		return
	}
	if g.Player.Y < g.Height-1 && g.World[g.Player.Y+1][g.Player.X] == 1 {
		return
	}
	g.SetPlayerPosition(g.Player.X, g.Player.Y+1)
}

func (g *GameWorld) MovePlayerLeft() {
	if g.Player.X < 0 {
		return
	}
	if g.Player.X > 0 && g.World[g.Player.Y][g.Player.X-1] == 1 {
		return
	}
	g.SetPlayerPosition(g.Player.X-1, g.Player.Y)
}

func (g *GameWorld) MovePlayerRight() {
	if g.Player.X > g.Width-1 {
		return
	}
	if g.Player.X < g.Width-1 && g.World[g.Player.Y][g.Player.X+1] == 1 {
		return
	}
	g.SetPlayerPosition(g.Player.X+1, g.Player.Y)
}
