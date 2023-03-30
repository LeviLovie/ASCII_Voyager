.PHONY: run editor build mod


# Run un normal mode
run:
	@go run main.go


# Run with flags: no-music, editor
no-music:
	@go run main.go --no-music
editor:
	@go run main.go --editor


# Run go mod ...
mod:
	@go mod tidy
	@go mod vendor


# Build game
# Game
objectGameBuild:
	@mkdir -p build/$(VERSION)/game
	@go build -o build/$(VERSION)/game/ASCII_Voyager main.go
exeGameBuild:
	@mkdir -p build/$(VERSION)/game
	@GOOS=windows go build -o build/$(VERSION)/game/ASCII_Voyager.exe main.go

# Editor
objectEditorBuild:
	@mkdir -p build/$(VERSION)/editor
	@go build -o build/$(VERSION)/editor/ASCII_Voyager_Editor editor/editor.go
exeEditorBuild:
	@mkdir -p build/$(VERSION)/editor
	@GOOS=windows go build -o build/$(VERSION)/editor/ASCII_Voyager_Editor.exe editor/editor.go

# Sort Commands
buildGame: objectGameBuild exeGameBuild
buildEditor: objectEditorBuild exeEditorBuild
build: buildGame buildEditor