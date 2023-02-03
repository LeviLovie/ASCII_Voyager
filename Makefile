run:
	@go run main.go

build:
	@go build -o build/build main.go

winBuild:
	@GOOS=windows go build -o build/build.exe main.go