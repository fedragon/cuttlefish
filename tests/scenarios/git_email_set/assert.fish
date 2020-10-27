#!/usr/bin/fish

if test (git config --get user.email) = "foo@bar.baz"
  echo '[git_email_set] ok'
else
  echo '[git_email_set] FAILED'
  exit 1
end