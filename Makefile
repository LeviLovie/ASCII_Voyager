run:
	@go run main.go

objectBuild:
	@mkdir -p build/$(VERSION)
	@go build -o build/$(VERSION)/ASCII_Voyager main.go

exeBuild:
	@mkdir -p build/$(VERSION)
	@GOOS=windows go build -o build/$(VERSION)/ASCII_Voyager.exe main.go

build: objectBuild exeBuild

mod:
	@go mod tidy
	@go mod vendor