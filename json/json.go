package json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/LeviiLovie/ASCII_Voyager/game"
)

//go:embed default.json
var defaultJson []byte

func NewSave(name string) {
	filename := fmt.Sprintf("./saves/%s.json", name)

	if _, err := os.Stat(filename); err == nil {
		return
	}

	err := os.WriteFile(filename, defaultJson, 0644)
	if err != nil {
		println("error:", err)
	}
}

func LoadSave(name string) game.GameWorld {
	filename := fmt.Sprintf("./saves/%s.json", name)

	var save game.GameWorld

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	err = json.Unmarshal(f, &save)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return save
}
