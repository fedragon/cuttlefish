#!/usr/bin/fish

if test -z "$fish_user_paths"
  echo '[unset_user_paths] ok'
else
  echo '[unset_user_paths] FAILED'
  exit 1
end