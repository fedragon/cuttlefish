#!/usr/bin/fish

if test "$fish_user_paths" = "/usr /tmp"
  echo '[user_pahts_set] ok'
else
  echo '[user_pahts_set] FAILED'
  exit 1
end