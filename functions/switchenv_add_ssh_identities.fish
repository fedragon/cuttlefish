function switchenv_add_ssh_identities
  if test -z "$switchenv_ssh_identities"
    for ssh_identity in $argv
      if test -n "$ssh_identity"
        ssh-add $ssh_identity
      end
    end
  end

  set --universal switchenv_ssh_identities $argv
end
