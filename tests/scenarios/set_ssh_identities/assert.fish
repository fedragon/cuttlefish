#!/usr/bin/fish

if test (ssh-add -L | grep ssh-rsa | wc -l) -eq 1
  echo '[set_ssh_identities] ok'
else
  echo '[set_ssh_identities] FAILED'
  exit 1
end