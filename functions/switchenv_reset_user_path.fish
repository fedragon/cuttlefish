function switchenv_reset_user_path
  if test -n "$switchenv_user_path_additions"
    for user_path in $switchenv_user_path_additions
      if set -l index (contains -i $user_path $fish_user_paths)
        set --erase --universal fish_user_paths[$index]
      end
    end
  end

  set --erase --universal switchenv_user_path_additions
end
