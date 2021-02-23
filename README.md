# wasm-hcl-to-json

Build the package via:

  make build

This will generate a hcltojson.wasm file and an index.js.

Usage:

    const load = require('hcltojson');
    const hcltojson = await load();
    const result = hcltojson(hclSource);
    console.log(result);

Convert a Terraform HCL file into JSON via:

    node index.js input.tf

The original wasm_exec.js file can be found at `$(go env GOROOT)/misc/wasm/wasm_exec.js`
