.PHONY: all build run

all: build run

build b: 
	go build -o ./parker ./main.go

run:
	./parker

test:
	go test ./...

$(V).SILENT: