default: build

build:
	go build -v -o wikicc cmd/wikicc/main.go
	go build -v -o wikicsort cmd/wikicsort/main.go

linux: clean
	GOOS=linux GOARCH=amd64 go build -v -o wikicc cmd/wikicc/main.go
	GOOS=linux GOARCH=amd64 go build -v -o wikicsort cmd/wikicsort/main.go

clean:
	go clean
