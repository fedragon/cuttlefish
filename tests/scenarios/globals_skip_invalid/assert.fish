#!/usr/bin/fish

for key in A B C
    if test -z "$$key"
        echo '[globals_skip_invalid] ok'
    else
        echo "[globals_skip_invalid] FAILED: global variable $$key is set"
        exit 1
    end
end