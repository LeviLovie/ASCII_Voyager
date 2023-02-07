package foo

type GameWorld struct {
	FPS        int     `json:"fps"`
	NeedRedraw bool    `json:"needRedraw"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	World      [][]int `json:"world"`
	Player     Player  `json:"player"`
}

func (g *GameWorld) Draw() {
	if !g.NeedRedraw {
		return
	}

	ClearScreen()

	screenPlayerX := GameWidth/2 + GameLeft
	screenPlayerY := GameHeight/2 + GameTop

	screenWorldX := screenPlayerX - g.Player.X
	if screenWorldX < GameLeft {
		screenWorldX = GameLeft
	}
	screenWorldY := screenPlayerY - g.Player.Y
	if screenWorldY < GameTop {
		screenWorldY = GameTop
	}

	for i := GameLeft; i < GameLeft+GameWidth; i++ {
		for j := GameTop; j < GameTop+GameHeight; j++ {
			if j < screenWorldY || j >= screenWorldY+g.Height || i < screenWorldX || i >= screenWorldX+g.Width {
				continue
			}
			MoveCursor(i, j)
			switch g.World[j-screenWorldY][i-screenWorldX] {
			case 0:
				PrintAt(i, j, ".")
			case 1:
				PrintAt(i, j, "#")
			}
		}
	}

	PrintAt(screenPlayerX, screenPlayerY, TEXT_CYAN+"@"+TEXT_RESET)
	g.NeedRedraw = false
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
