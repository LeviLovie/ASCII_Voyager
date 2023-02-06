package foo

import (
	"fmt"
)

type GameWorld struct {
	FPS             int     `json:"fps"`
	NeedRedraw      bool    `json:"needRedraw"`
	Width           int     `json:"width"`
	Height          int     `json:"height"`
	World           [][]int `json:"world"` // 0 - empty, 1 - wall
	PlayerPositionX int     `json:"playerPositionX"`
	PlayerPositionY int     `json:"playerPositionY"`
}

func (g *GameWorld) Draw() {
	if !g.NeedRedraw {
		return
	}

	ClearScreen()
	MoveCursor(0, 0)
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			switch g.World[i][j] {
			case 0:
				MoveCursor(j+(g.PlayerPositionX-((125-1)/2)), i+(g.PlayerPositionY-(29-1)/2))
				fmt.Print(" ")
			case 1:
				MoveCursor(j+(g.PlayerPositionX-((125-1)/2)), i+(g.PlayerPositionY-(29-1)/2))
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	MoveCursor(g.PlayerPositionX-((125-1)/2), g.PlayerPositionY-(29-1)/2)
	fmt.Print(TEXT_CYAN + "@" + TEXT_RESET)
	g.NeedRedraw = false
}

func (g *GameWorld) SetPlayerPosition(x, y int) {
	g.PlayerPositionX = x
	g.PlayerPositionY = y
	g.NeedRedraw = true
}

func (g *GameWorld) MovePlayer(x, y int) {
	g.PlayerPositionX += x
	g.PlayerPositionY += y
	g.NeedRedraw = true
}

func (g *GameWorld) MovePlayerUp() {
	if g.PlayerPositionY < 0 {
		return
	}
	if g.PlayerPositionY > 0 && g.World[g.PlayerPositionY-1][g.PlayerPositionX] == 1 {
		return
	}
	g.SetPlayerPosition(g.PlayerPositionX, g.PlayerPositionY-1)
}

func (g *GameWorld) MovePlayerDown() {
	if g.PlayerPositionY > g.Height-1 {
		return
	}
	if g.PlayerPositionY < g.Height-1 && g.World[g.PlayerPositionY+1][g.PlayerPositionX] == 1 {
		return
	}
	g.SetPlayerPosition(g.PlayerPositionX, g.PlayerPositionY+1)
}

func (g *GameWorld) MovePlayerLeft() {
	if g.PlayerPositionX < 0 {
		return
	}
	if g.PlayerPositionX > 0 && g.World[g.PlayerPositionY][g.PlayerPositionX-1] == 1 {
		return
	}
	g.SetPlayerPosition(g.PlayerPositionX-1, g.PlayerPositionY)
}

func (g *GameWorld) MovePlayerRight() {
	if g.PlayerPositionX > g.Width-1 {
		return
	}
	if g.PlayerPositionX < g.Width-1 && g.World[g.PlayerPositionY][g.PlayerPositionX+1] == 1 {
		return
	}
	g.SetPlayerPosition(g.PlayerPositionX+1, g.PlayerPositionY)
}
