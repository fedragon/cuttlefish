#!/usr/local/bin/fish

switch $PWD/
  case "$switchenv_last_visited/*"
    exit 0
  case "*"
    switchenv_add_ssh_identities ~/.ssh/my_id_rsa ~/.ssh/another_id_rsa
    switchenv_append_to_user_path /some/path/ /some/other/path
    switchenv_set_user_variable my_home /somewhere
end
