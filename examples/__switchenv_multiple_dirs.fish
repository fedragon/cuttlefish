# scenario: we have multiple "special" directories
# goal: we want to reset the environment whenever we switch from one to the other
function __switchenv --on-variable PWD
  switch $PWD/
    case "~/personal/*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
      ~/personal/.switchenvrc.fish
      set --universal switchenv_last_visited "~/personal/*"

    case "~/work/*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
      ~/work/.switchenvrc.fish
      set --universal switchenv_last_visited "~/work/*"
  end
end
