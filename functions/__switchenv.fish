function __switchenv --on-variable PWD
  switch $PWD/
    case "/my/path/*" # change me!
      switchenv_reset_user_path
      switchenv_unload_ssh_identities

      /my/path/.switchenv.fish # change me!

    # other cases as needed
  end
end
