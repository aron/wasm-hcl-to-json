# wasm-hcl-to-json

Uses the Go [hcl2json](https://github.com/tmccombs/hcl2json/convert) package to expose a function to JavaScript via WebAssembly.

Unfortunately the Go WebAssembly implementation currently depends on exporting functions into the global scope. We use a hacky shim to wrap the provided wasm_exec.js file and ensure that our functions do not pollute the Node environment.

Build the package via:

    make build

This will generate a hcltojson.wasm file and an index.js.

Usage:

    const fs = require('fs');
    const load = require('hcltojson');
    const hcltojson = await load();
    const result = hcltojson(fs.readFileSync("myterraform.tf").toString("utf-8"));
    console.log(result);

Convert a Terraform HCL file into JSON via:

    node index.js input.tf

The original wasm_exec.js file can be found at `$(go env GOROOT)/misc/wasm/wasm_exec.js`

## To build a CLI Tool

    make hcl2json

This will spit out a ~5mb go binary that can be used as follows:

    <myterraform.tf hcl2json
    > ... some json output ...

This binary can be pulled down to 1.5mb via building with the flags
`-ldflags="-s -w"` and then using the upx packer tool.
