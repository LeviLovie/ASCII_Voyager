package json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/sirupsen/logrus"
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

func LoadSave(name string) foo.GameWorld {
	filename := fmt.Sprintf("./saves/%s.json", name)
	logrus.Debug("Json - LoadSave: ", filename)

	var save foo.GameWorld

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	logrus.Debug("Json - LoadSave: ", string(f))

	err = json.Unmarshal(f, &save)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return save
}

func SaveGame(name string, save foo.GameWorld) {
	filename := fmt.Sprintf("./saves/%s.json", name)

	f, err := json.Marshal(save)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(filename, f, 0644)
	if err != nil {
		println("error:", err)
	}
}
