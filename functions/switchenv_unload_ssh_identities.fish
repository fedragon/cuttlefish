function switchenv_unload_ssh_identities
  if test -n "$switchenv_ssh_identities"
    for ssh_identity in $switchenv_ssh_identities
      if test -n "$ssh_identity"
        ssh-add -d $ssh_identity
      end
    end
  end

  set --erase switchenv_ssh_identities
end
