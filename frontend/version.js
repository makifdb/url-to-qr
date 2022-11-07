var fs = require('fs');
var pjson = require('./package.json');
var nversion = process.argv[2];
pjson.version = nversion;
fs.writeFileSync('./package.json', JSON.stringify(pjson, null, 2));