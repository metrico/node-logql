const logql = require('./');
process.argv.splice(0, 2);
var query = process.argv[0] || false;
if (!query) process.exit();

/* pipe stdin */
var lineReader = require('readline').createInterface({
  input: process.stdin
});

/* process logql */
lineReader.on('line', function (line) {
  const parsed = logql.parse(query, line);
  process.stdout.write(parsed+"\n");
});
