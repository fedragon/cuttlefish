# scenario: we have only one "special" directory
# goal: we want to reset the environment whenever we move out of it
function __switchenv --on-variable PWD
  switch $PWD/
    case "~/personal/*"
      ~/personal/.switchenvrc.fish
      set --universal switchenv_last_visited "~/personal/*"

    case "*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
  end
end
