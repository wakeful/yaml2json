all: build

dependencies:
	go get -d -v -t

build: dependencies
	mkdir -p out
	go build -o out/yaml2json .

install: dependencies
	go install .
