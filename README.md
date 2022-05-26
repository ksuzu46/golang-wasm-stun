# Golang wasm stun

This repository is based on:
- https://github.com/ianfinch/golang-wasm-calc
- https://github.com/pion/stun

## Get started

#### Copy `wasm_exec.js`:

Use Go:
```bash
cp `go env GOROOT`/misc/wasm/wasm_exec.js public/js/
```

Use TinyGo:
```bash
wget https://raw.githubusercontent.com/tinygo-org/tinygo/release/targets/wasm_exec.js -O public/js/wasm_exec.js
```

#### Build:

Go:
```bash
GOOS=js GOARCH=wasm go build -o public/js/main.wasm src/main.go
```
TinyGo:
```bash
tinygo build -o public/js/main.wasm -target wasm src/main.go
```

#### Run server:
```bash
go run src/server.go
```

TODO:
- Use WASI since Wasm does not support OS features like file I/O or Networking

Stuff which may be related to this:
- https://stackoverflow.com/questions/55880920/dial-tcp-protocol-not-available-go-webassembly-test
- https://github.com/WebAssembly/design/issues/1251
- https://emscripten.org/docs/introducing_emscripten/about_emscripten.html
- https://github.com/wasmerio/wasmer-js/tree/main/examples/node
- https://text.baldanders.info/golang/wasi-with-tinygo/