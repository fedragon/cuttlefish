# fish-switchenv

Switch environment according to current working directory. Heavily inspired by the amazing [direnv](https://direnv.net/) but not limited to environment variables and only targeting [fish shell](https://fishshell.com).

## Features

Configure any combination of the following in a simple YAML file:
- ssh identities
- paths (to append to `$fish_user_paths`)
- global variables
- git user email address

Whenever the directory in which the configuration file resides is visited, its contents are parsed and sourced.

## Install

The installation requires Go 1.13+.

```
git clone https://github.com/fedragon/fish-switchenv
cd fish-switchenv
make
```

**Note:** The installation makes small changes (= appending a line) to your `~/.config/fish/config.fish`: before doing so, though, it will copy the original file to `~/.config/fish/config.fish_bck` so that it can be restored at will (see `install` target in `Makefile` for more details).

## Usage

Create a `.switchenv.yaml` in every directory where you want to load a different environment: its syntax is described in `examples/switchenv.yaml`.

## Uninstall

```
cd fish-switchenv
make uninstall
```