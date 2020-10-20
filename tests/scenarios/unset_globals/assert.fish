#!/usr/bin/fish

if test "$BAZ" = "buzz" -a -z "$BAR" -a -z "$FOO"
  echo '[unset_globals] ok'
else
  echo '[unset_globals] FAILED'
  exit 1
end