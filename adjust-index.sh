#!/bin/sh

sed -i 's/public\/main.wasm/main.wasm/' public/index.html || exit 1
sed -i 's/public\/wasm_exec.js/wasm_exec.js/' public/index.html || exit 1
