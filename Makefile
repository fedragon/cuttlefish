all: build install
.PHONY: all

current_dir = $(shell pwd)

build:
	go get
	go build

install:
	cp ./fish-switchenv /usr/local/bin/
	cp ./functions/__switchenv.fish ~/.config/fish/functions/
	cp ~/.config/fish/config.fish ~/.config/fish/config.fish_bck
	echo "\nfunctions -c __switchenv switchenv" >> ~/.config/fish/config.fish

uninstall:
	rm /usr/local/bin/fish-switchenv
	rm ~/.config/fish/functions/__switchenv.fish
	cp ~/.config/fish/config.fish_bck ~/.config/fish/config.fish

test:
	env GOOS=linux GOARCH=amd64 go build -o tests/fish-switchenv
	docker build -t fish-switchenv:latest .
	docker run --rm -v $(current_dir)/tests/scenarios:/home/swen/scenarios -w /home/swen/scenarios fish-switchenv:latest run_assertions.fish