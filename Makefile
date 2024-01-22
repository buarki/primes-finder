gowasm:
	GOARCH=wasm GOOS=js go build -o public/main.wasm cmd/wasm/main.go

wasm_exec:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" public/wasm_exec.js

clean_public:
	rm public/main.wasm && rm public/wasm_exec.js

local: clean_public gowasm wasm_exec
	go build -o bin/web cmd/web/main.go && PORT=8080 bin/web
