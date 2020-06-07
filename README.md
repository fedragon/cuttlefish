# fish-switchenv

Switch environment according to current working directory. Inspired by [direnv](https://direnv.net/) but implemented solely for [fish shell](https://fishshell.com) and not limited to environment variables. Created in order to refresh/expand my fish scripting.

## Features

- load ssh identities according to current working directory
- append elements to user paths according to current working directory

## Functions

Here is a list of the provided functions, which can be mixed & matched according to your needs.

### `switchenv_load_ssh_identities`

Adds the provided identity/ies to the ssh agent and stores it/them in the `$switchenv_ssh_identities` universal variable.

Arguments: a non-empty list of absolute paths pointing to the ssh identities to load

Usage:

```
switchenv_load_ssh_identities ~/.ssh/work_id_rsa ~/.ssh/work_cvs_id_rsa
```

### `switchenv_unload_ssh_identities`

Removes all ssh identities found in `$switchenv_ssh_identities` from the ssh agent and then erases the variable.

Arguments: none

Usage:

```
switchenv_unload_ssh_identities
```

### `switchenv_append_to_user_path`:

Adds the provided path(s) to the `$fish_user_paths` universal variable and stores them in the `$switchenv_user_path_additions` universal variable.

Arguments: a non-empty list of absolute paths

Usage:

```
switchenv_append_to_user_path /some/path ~/some/other/path ... /yet/another/path
```

### `switchenv_reset_user_path`

Removes all paths found in `$switchenv_user_path_additions` from `$fish_user_paths` and then erases `$switchenv_user_path_additions`.

Arguments: none

Usage:

```
switchenv_reset_user_path
```

## Usage

1. Copy all `switchenv_*` functions in `functions/` to `~/.config/fish/functions`

```
cp ./functions/switchenv_* ~/.config/fish/functions/
```

2. Look at the examples provided in `examples/`, then create a `~/.config/fish/functions/__switchenv.fish` function fitting your needs.

3. Add the following instruction anywhere in your own `~/.config/fish/config.fish`

```
functions -c __switchenv switchenv
```

(example provided in `./config.fish`)

4. Customize `switchenvrc.fish` according to your needs, then copy it to `/my/path/.switchenvrc.fish` (note the starting dot) and make sure it's executable by your user.
