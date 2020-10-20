#!/usr/bin/fish

if test "$FOO" = "bar" -a "$BAR" = "baz"
  echo '[set_globals] ok'
else
  echo '[set_globals] FAILED'
  exit 1
end