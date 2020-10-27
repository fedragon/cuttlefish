#!/usr/bin/fish

if test "$BAZ" = "buzz" -a -z "$BAR" -a -z "$FOO"
  echo '[globals_unset] ok'
else
  echo '[globals_unset] FAILED'
  exit 1
end