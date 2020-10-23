all: build install
.PHONY: all

current_dir = $(shell pwd)

build:
	go get
	go build

install:
	cp ./cuttlefish /usr/local/bin/
	cp ./functions/__cuttlefish.fish ~/.config/fish/functions/
	cp ~/.config/fish/config.fish ~/.config/fish/config.fish_bck
	echo "\nfunctions -c __cuttlefish cuttlefish" >> ~/.config/fish/config.fish

uninstall:
	rm /usr/local/bin/cuttlefish
	rm ~/.config/fish/functions/__cuttlefish.fish
	cp ~/.config/fish/config.fish_bck ~/.config/fish/config.fish

test:
	env GOOS=linux GOARCH=amd64 go build -o tests/cuttlefish
	docker build -t cuttlefish:latest .
	docker run --rm -v $(current_dir)/tests/scenarios:/home/nemo/scenarios -w /home/nemo/scenarios cuttlefish:latest run_assertions.fish