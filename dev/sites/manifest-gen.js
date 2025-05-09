const fs = require("fs/promises");
const path = require("path");
const util = require('util');
const exec = util.promisify(require('child_process').exec);

async function main(dir) {
  let manifest = {};

  let files = await fs.readdir(dir, {withFileTypes: true});
  for (let i = 0; i < files.length; i++) {
    if (!files[i].isFile()) {
      continue;
    }

    let filename = files[i].name;
    let filepath = path.join(dir, filename);
    const { stdout, stderr } = await exec(`sha256sum ${filepath} | cut -d " " -f 1`);
    if (stderr) {
      throw new Error(stderr);
    }
    
    let hash = stdout.trim();
    manifest[filename] = hash;
  }

  console.log(JSON.stringify(manifest));
}

main(process.argv[2])
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
