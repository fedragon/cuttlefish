#!/usr/bin/fish

if test "$FOO" = "bar" -a "$BAR" = "baz"
  echo '[globals_set] ok'
else
  echo '[globals_set] FAILED'
  exit 1
end