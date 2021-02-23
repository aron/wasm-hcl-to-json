const fs = require('fs');
const load = require('../index');
const assert = require('assert');

(async () => {
  try {
    const hclToJSON = await load();
    const result = hclToJSON(fs.readFileSync(process.argv[2]).toString('utf-8'));

    assert.equal(JSON.stringify(result, null, "  "), JSON.stringify({
      "resource": {
        "aws_sqs_queue": {
          "sqs_2": [
            {
              "kms_data_key_reuse_period_seconds": 300,
              "kms_master_key_id": "alias/aws/sqs",
              "name": "terraform-example-queue"
            }
          ]
        }
      }
    }, null, "  "))

    console.log("PASSED");
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
})();
