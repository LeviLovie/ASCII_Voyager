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
	@mkdir -p build/$(V)
	go build -o build/$(V)/ASCII_Voyager main.go
exeGameBuild:
	@mkdir -p build/$(V)
	GOOS=windows go build -o build/$(V)/ASCII_Voyager.exe main.go
build: objectGameBuild exeGameBuild