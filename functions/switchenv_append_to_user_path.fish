function switchenv_append_to_user_path
  if test -z "$switchenv_user_path_additions"
    for user_path in $argv
      contains $user_path $fish_user_paths; or set --universal --append fish_user_paths $user_path
    end
  end

  set --universal switchenv_user_path_additions $argv
end
