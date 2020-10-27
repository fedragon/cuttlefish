#!/usr/bin/fish

if test (git config --get user.email) != "nemo@cuttle.fish"
  echo '[git_email_skip_invalid] ok'
else
  echo '[git_email_skip_invalid] FAILED'
  exit 1
end