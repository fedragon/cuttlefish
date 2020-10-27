#!/usr/bin/fish

if test (ssh-add -L | grep ssh-rsa | wc -l) -eq 0
  echo '[ssh_identities_skip_invalid] ok'
else
  echo '[ssh_identities_skip_invalid] FAILED'
  exit 1
end