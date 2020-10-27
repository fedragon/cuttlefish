#!/usr/bin/fish

if test (ssh-add -L | grep ssh-rsa | wc -l) -eq 1
  echo '[ssh_identities_set] ok'
else
  echo '[ssh_identities_set] FAILED'
  exit 1
end