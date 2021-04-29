all:
	GOOS=js GOARCH=wasm go build -o php-strip.wasm

test:
	GOOS=js GOARCH=wasm go test -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec"