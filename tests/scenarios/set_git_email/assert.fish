#!/usr/bin/fish

if test (git config --get user.email) = "foo@bar.baz"
  echo '[set_git_email] ok'
else
  echo '[set_git_email] FAILED'
  exit 1
end