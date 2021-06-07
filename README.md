cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./lib

GOOS=js GOARCH=wasm go build -o lib/main.wasm

$env:GOOS='js';$env:GOARCH='wasm';go build -o lib/main.wasm
