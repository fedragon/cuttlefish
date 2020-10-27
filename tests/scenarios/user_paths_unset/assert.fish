#!/usr/bin/fish

if test -z "$fish_user_paths"
  echo '[user_paths_unset] ok'
else
  echo '[user_paths_unset] FAILED'
  exit 1
end