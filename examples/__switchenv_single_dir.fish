# scenario: we have only one "special" directory
# goal: we want to reset the environment whenever we move out of it
function __switchenv --on-variable PWD
  if not test -f /tmp/switchenv
    switchenv_erase
  end

  switch $PWD/
    case "$HOME/personal/*"
      if test -z "$switchenv_last_visited"
        $HOME/personal/.switchenvrc.fish

        set --universal switchenv_last_visited "$HOME/personal"
      end

    case "*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
  end

  touch /tmp/switchenv
end
