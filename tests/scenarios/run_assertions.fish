#!/usr/bin/fish

cd set_globals
./assert.fish

cd ../unset_globals
./assert.fish

cd ../set_git_email
./assert.fish

cd ..
echo 'y' | ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa 2>&1 >/dev/null
cd set_ssh_identities
./assert.fish

cd ../set_user_paths
./assert.fish

cd ../unset_user_paths
./assert.fish