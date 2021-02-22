const fs = require('fs');
require('./wasm_exec');

(async () => {
  try {
    const buf = fs.readFileSync('./main.wasm');
    const go = new Go();
    const result = await WebAssembly.instantiate(buf, go.importObject);
    go.run(result.instance)
    process.stdout.write(hclToJSON(fs.readFileSync(process.argv[2]).toString('utf-8')));
    process.stdout.write("\n");
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
})();
