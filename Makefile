default: build

build:
	go build -v -o wikicc cmd/wikicc/main.go

linux: clean
	GOOS=linux GOARCH=amd64 go build -v -o wikicc cmd/wikicc/main.go

clean:
	go clean
