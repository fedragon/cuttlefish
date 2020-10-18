# fish-switchenv

Switch environment according to current working directory. Concept inspired by [direnv](https://direnv.net/) but not limited to environment variables and targeting only [fish shell](https://fishshell.com).

## Features

Configure any combination of the following in a simple YAML file:
- ssh identities
- paths (to append to `$fish_user_paths`)
- global variables
- git user email address

Whenever the directory in which the configuration file resides is visited, its contents are parsed and sourced.

## Install

The installation requires Go 1.13+.

**Note:** `make install` makes small changes (= append a line) to `~/.config/fish/config.fish`: to be on the safe side, it will create a copy of the original file to `~/.config/fish/config.fish_bck` so that it can be restored at will.

```
git clone https://github.com/fedragon/fish-switchenv
cd fish-switchenv
make build; and make install
```

## Usage

Create a `.switchenv.yaml` in every directory where you want to load a different environment: its syntax is described in `examples/switchenv.yaml`.

## Uninstall

```
cd fish-switchenv
make uninstall
```