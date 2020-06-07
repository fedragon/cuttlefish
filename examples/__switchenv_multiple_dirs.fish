# scenario: we have multiple "special" directories
# goal: we want to reset the environment whenever we switch from one to the other
function __switchenv --on-variable PWD
  switch $PWD/
    case "~/personal/*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
      ~/personal/.switchenvrc.fish

    case "~/work/*"
      switchenv_reset_user_path
      switchenv_reset_ssh_identities
      ~/work/.switchenvrc.fish
  end
end
