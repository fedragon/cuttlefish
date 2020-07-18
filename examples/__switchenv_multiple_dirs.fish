# scenario: we have multiple "special" directories
# goal: we want to reset the environment whenever we switch from one to the other
function __switchenv --on-variable PWD
  if not test -f /tmp/switchenv
    switchenv_erase
  end

  switch $PWD/
    case "$HOME/personal/*"
      if test "$HOME/personal" != "$switchenv_last_visited"
        switchenv_reset_user_path
        switchenv_reset_ssh_identities
        $HOME/personal/.switchenvrc.fish

        set --universal switchenv_last_visited "$HOME/personal"
      end

    case "$HOME/work/*"
      if test "$HOME/work" != "$switchenv_last_visited"
        switchenv_reset_user_path
        switchenv_reset_ssh_identities
        $HOME/work/.switchenvrc.fish

        set --universal switchenv_last_visited "$HOME/work"
      end
  end

  touch /tmp/switchenv
end
