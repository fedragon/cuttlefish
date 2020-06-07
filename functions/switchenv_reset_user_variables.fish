function switchenv_reset_user_variables
  if test -n "$switchenv_user_variables"
    for user_var in $switchenv_user_variables
      set --erase $user_var
    end
  end

  set --erase --universal switchenv_user_variables
end
