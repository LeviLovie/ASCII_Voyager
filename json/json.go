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

func NewSave(name string, keys chan foo.KeyPress) int {
	// cutscenes.SatrtGameCuts(keys)
	filename := fmt.Sprintf("./saves/%s.dat", name)

	if _, err := os.Stat(filename); err == nil {
		return 1
	}

	err := os.WriteFile(filename, defaultJson, 0644)
	if err != nil {
		logrus.Error("Json - NewSave: ", err)
		return 1
	}

	return 0
}

func LoadSave(name string) (int, foo.GameWorld) {
	filename := fmt.Sprintf("./saves/%s.dat", name)
	logrus.Debug("Json - LoadSave: ", filename)

	var save foo.GameWorld

	f, err := os.ReadFile(filename)
	if err != nil {
		logrus.Error("Json - LoadSave: ", err)
		return 1, save
	}

	logrus.Debug("Json - LoadSave: ", string(f))

	err = json.Unmarshal(f, &save)
	if err != nil {
		logrus.Error("Json - LoadSave: ", err)
		return 1, save
	}

	return 0, save
}

func SaveGame(name string, save foo.GameWorld) int {
	filename := fmt.Sprintf("./saves/%s.dat", name)

	f, err := json.Marshal(save)
	if err != nil {
		return 1
	}

	err = os.WriteFile(filename, f, 0644)
	if err != nil {
		return 1
	}

	return 0
}
