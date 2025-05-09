// usage node generateNodes.js <numContentNodes>
// This script will generate <numContentNodes> inexorable-content nodes in the dev.json file

const fs = require('fs');
const path = require('path');

async function generateContentNodes (numContentNodes) {
  const dev = require(__dirname + '/dev.base.json');

  for (let nodeId = 1; nodeId <= numContentNodes; nodeId++) {
    const port = 80;
    console.log("Adding node with port", port, "nodeId", nodeId);
    dev.EarthfastNodes.nodes.push({
      "id": `0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3b${30082 + nodeId}`,
      "operatorId": "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
      "host": `dev-content-${nodeId}-1__${port}.armadalocal.test`,
      "region": "xx",
      "disabled": false,
      "prices": [
        "0.0",
        "0.0"
      ],
      "projectIds": [
        "0x0000000000000000000000000000000000000000000000000000000000000000",
        "0x0000000000000000000000000000000000000000000000000000000000000000"
      ]
    });
  }

  dev.EarthfastRegistry.nonce = (parseInt(dev.EarthfastRegistry.nonce, 10) + numContentNodes).toString(10);

  fs.writeFileSync(__dirname + '/dev.json', JSON.stringify(dev, null, 2));
}

// take first user arg as parameter
const numContentNodes = parseInt(process.argv[2], 10) || 1;
console.log("numContentNodes", numContentNodes, "type", typeof numContentNodes);
generateContentNodes(numContentNodes);
