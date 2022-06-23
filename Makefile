all: build install
.PHONY: all

current_dir = $(shell pwd)

build:
	go get
	go build

clean:
	rm -r dist

install:
	sudo cp ./cuttlefish /usr/local/bin/
	cp ./functions/__cuttlefish.fish ~/.config/fish/functions/
	echo "Backing up your ~/.config/fish/config.fish to ~/.config/fish/config.fish_bck"
	cp ~/.config/fish/config.fish ~/.config/fish/config.fish_bck
	echo "Appending cuttlefish initialization to ~/.config/fish/config.fish"
	echo "\nfunctions -c __cuttlefish cuttlefish" >> ~/.config/fish/config.fish

uninstall:
	sudo rm /usr/local/bin/cuttlefish
	rm ~/.config/fish/functions/__cuttlefish.fish
	cp ~/.config/fish/config.fish_bck ~/.config/fish/config.fish

test:
	env GOOS=linux GOARCH=amd64 go build -o tests/cuttlefish
	docker build -t cuttlefish:latest .
	docker run --rm -v $(current_dir)/tests/scenarios:/home/nemo/scenarios -w /home/nemo/scenarios cuttlefish:latest run_assertions.fish

release: clean
	mkdir -p dist/darwin-amd64 dist/darwin-arm64
	go get
	# amd64
	env GOOS=darwin GOARCH=amd64 go build -o dist/darwin-amd64/cuttlefish
	cp -r functions dist/darwin-amd64/
	cp Makefile dist/darwin-amd64/
	tar -cvzf binaries-darwin-amd64.tar.gz -C dist darwin-amd64
	# arm64
	env GOOS=darwin GOARCH=arm64 go build -o dist/darwin-arm64/cuttlefish
	cp -r functions dist/darwin-arm64/
	cp Makefile dist/darwin-arm64/
	tar -cvzf binaries-darwin-arm64.tar.gz -C dist darwin-arm64
