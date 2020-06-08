# fish-switchenv

Switch environment according to current working directory. Inspired by [direnv](https://direnv.net/) but implemented solely for [fish shell](https://fishshell.com) and not limited to environment variables. Pet project created in order to improve my fish scripting skills.

## Features

Enables user to configure any combination of the following actions, which will be automatically triggered according to the current working directory:
- set global variables
- add ssh identities to ssh agent
- append path(s) to `$fish_user_paths`

## Functions

Here is a list of the provided functions, which can be mixed & matched according to your needs.

### `switchenv_add_ssh_identities`

Adds the provided identity/ies to the ssh agent and stores it/them in the `$switchenv_ssh_identities` universal variable.

Arguments: a non-empty list of absolute paths pointing to the ssh identities to load

Usage:

```
switchenv_add_ssh_identities ~/.ssh/work_id_rsa ~/.ssh/work_cvs_id_rsa
```

### `switchenv_reset_ssh_identities`

Removes all ssh identities found in `$switchenv_ssh_identities` from the ssh agent and then erases the variable.

Arguments: none

Usage:

```
switchenv_reset_ssh_identities
```

### `switchenv_set_user_variable`

Creates a global variable with provided name and value(s) and appends its name to the `$switchenv_user_variables` universal variable.

Arguments: `variable name` `value1` `value2` `valueN`

Usage:

```
switchenv_set_user_variable my_home /opt/home/bin
switchenv_set_user_variable foo bar baz
```

### `switchenv_reset_user_variables`

Erases all variables found in `$switchenv_user_variables` and then the universal variable itself.

Arguments: none

Usage:

```
switchenv_reset_user_variables
```

### `switchenv_append_to_user_path`

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

4. Customize `examples/switchenvrc.fish` according to your needs, then copy it to `/my/path/.switchenvrc.fish` (note the starting dot) and make sure it's executable by your user.
