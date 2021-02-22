# wasm-hcl-to-json

Build the Go binary via:

    GOOS=js GOARCH=wasm go build -o ./main.wasm

Convert a Terraform HCL file into JSON via:

    node main.js input.tf

The original wasm_exec.js file can be found at `$(go env GOROOT)/misc/wasm/wasm_exec.js` but unfortunately exposes globals :/
