# wasm-hcl-to-json

Build the Go binary via:

    GOOS=js GOARCH=wasm go build -o ./main.wasm

Convert a Terraform HCL file into JSON via:

    node main.js input.tf
