#!/usr/bin/fish

if test (git config --get user.name) != "Foo Bar"
  echo '[git_name_skip_invalid] ok'
else
  echo '[git_name_skip_invalid] FAILED'
  exit 1
end