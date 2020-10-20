#!/usr/bin/fish

if test "$fish_user_paths" = "/usr /tmp"
  echo '[set_user_paths] ok'
else
  echo '[set_user_paths] FAILED'
  exit 1
end