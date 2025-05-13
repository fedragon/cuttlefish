#!/usr/bin/fish

if test (git config --get user.name) = "Foo Bar"
  echo '[git_name_set] ok'
else
  echo '[git_name_set] FAILED'
  exit 1
end