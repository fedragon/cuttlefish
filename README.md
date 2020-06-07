# fish-switchenv

Switch environment according to current working directory. Inspired by [direnv](https://direnv.net/) but implemented specifically for [fish shell](https://fishshell.com) and not limited to environment variables.

## Features

- load ssh identities according to current working directory
- append elements to user paths according to current working directory

If multiple "special" working directories are defined, switching from one to another will undo the changes introduced by the previous one, unloading any custom ssh identities and resetting user paths.

## Usage

1. Copy all `switchenv_*` functions in `functions/` to `~/.config/fish/functions`

```
cp ./functions/switchenv_* ~/.config/fish/functions/
```

2. Customize `functions/__switchenv` according to your needs, then copy it to `~/.config/fish/functions/`

3. Add the following instruction anywhere in your own `~/.config/fish/config.fish`

```
functions -c __switchenv switchenv
```

(example provided in `./config.fish`)

4. Customize `switchenv.fish` according to your needs, then copy it to `/my/path/.switchenv.fish` (note the starting dot)
