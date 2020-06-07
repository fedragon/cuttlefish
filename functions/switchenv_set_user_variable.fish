function switchenv_set_user_variable
  set -l user_var $argv[1]
  set -l values $argv[2..(count $argv)]

  set --global $user_var $values

  if not contains $user_var $switchenv_user_variables
    set --universal --append switchenv_user_variables $user_var
  end
end
