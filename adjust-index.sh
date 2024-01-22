#!/bin/sh

## This script is to trap vercel in production environment :)
## It will just removes the "/public" strings from the public/index.html
## making it possible to fetch the needed files public/main.wasm and 
## public/wasm_exec.js

sed -i 's/public\/main.wasm/main.wasm/' public/index.html || exit 1
sed -i 's/public\/wasm_exec.js/wasm_exec.js/' public/index.html || exit 1
