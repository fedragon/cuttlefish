#!/usr/bin/fish

cd git_email_set
./assert.fish

cd ../git_email_skip_invalid
./assert.fish

cd ../globals_skip_invalid
./assert.fish

cd ../globals_set
./assert.fish

cd ../globals_unset
./assert.fish

cd ..
echo 'y' | ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa 2>&1 >/dev/null
cd ssh_identities_set
./assert.fish

cd ../ssh_identities_skip_invalid
./assert.fish

cd ../user_paths_set
./assert.fish

cd ../user_paths_unset
./assert.fish

cd ../user_paths_skip_invalid
./assert.fish