package foo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type KeyPress struct {
	Char rune
	Key  keyboard.Key
}

func GetFilesInDir() []string {
	root := "./saves"
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) != ".json" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func InitLog() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	closeLogFile := func() {
		if f != nil {
			f.Close()
		}
	}
	logrus.RegisterExitHandler(closeLogFile)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(f)
	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "15:04:05",
		LogFormat:       "%time% [%lvl%] %msg%\n",
	})
}

func DrawVerticalSplitLine(height int) {
	for y := 0; y < height-1; y++ {
		MoveCursor(80, y)
		fmt.Print("║")
	}
}

func WriteTextOnCenter(text string, width, y int) {
	var x = (width - len(text)) / 2
	MoveCursor(x, y)
	fmt.Print(text)
}
