const path = require('path');
const fs = require('fs');

const shim = {...global};
const proxy = new Proxy(shim, {
  set(target, prop, value) {
    target[prop] = value;
    return true;
  }
});

((global) => {
  //REPLACE//

  // Needed to shim the "global" variables.
  var fs = global.fs;
  var Go = global.Go;
  var crypto = global.crypto;
  var performance = global.performance;
})(proxy);

module.exports = async function load() {
  const { Go } = shim;
  const buf = fs.readFileSync(path.join(__dirname, 'hcltojson.wasm'));
  const go = new Go();
  const result = await WebAssembly.instantiate(buf, go.importObject);
  go.run(result.instance)
  return (input) => JSON.parse(shim.hclToJSON(input));
}
