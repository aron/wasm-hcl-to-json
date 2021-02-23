goroot = $(shell go env GOROOT)

build: index.js

index.js: template.js wasm_exec.tmp.js
	sed -e '/REPLACE/ {' -e 'r wasm_exec.tmp.js' -e 'd' -e '}' $< > $@

hcltojson.wasm: main.go
	GOOS=js GOARCH=wasm go build -o $@

wasm_exec.tmp.js:
	cat "$(goroot)/misc/wasm/wasm_exec.js" > $@

.PHONY: test
test: index.js
	node test/test.js test/input.tf
