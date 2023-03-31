package editor

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/LeviiLovie/ASCII_Voyager/foo"
	"github.com/sirupsen/logrus"
)

var (
	//go:embed static
	static    embed.FS
	templates = make(map[string]*template.Template)
)

func index(t *template.Template) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var save foo.GameWorld

		f, err := os.ReadFile("json/default.json")
		if err != nil {
			logrus.Error("Error reading file: ", err)
		}

		err = json.Unmarshal(f, &save)
		if err != nil {
			logrus.Error("Error unmarshalling: ", err)
		}

		saveRaw, err := json.Marshal(save)
		if err != nil {
			logrus.Error("Error marshalling: ", err)
		}

		data := map[string]interface{}{
			"width":   save.Width,
			"height":  save.Height,
			"world":   save.World,
			"playerX": save.Player.X,
			"playerY": save.Player.Y,
			"blocks":  save.Blocks,
			"version": save.Version,
			"saveRaw": string(saveRaw),
		}

		t.Execute(w, data)
	}
}

func Editor() {
	tmplFiles, err := fs.ReadDir(static, "static")
	if err != nil {
		logrus.Error("Error reading static directory: ", err)
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() || !strings.HasSuffix(tmpl.Name(), ".html") {
			continue
		}

		pt, err := template.ParseFS(static, path.Join("static", tmpl.Name()))
		if err != nil {
			logrus.Error("Error parsing template: ", err)
		}

		templates[tmpl.Name()] = pt
	}

	http.HandleFunc("/", index(templates["index.html"]))

	fmt.Println("Starting editor on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Error("Error starting editor: ", err)
	}
}
