#!/bin/sh

## This script is to trap vercel in production environment :)
## It will just removes the "/public" strings from the public/index.html
## making it possible to fetch the needed files public/main.wasm and 
## public/wasm_exec.js

sed -i 's/public\/\([^ ]*\)/\1/' public/index.html || exit 1
