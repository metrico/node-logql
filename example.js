const logql = require('./');
process.argv.splice(0, 2);

var query = process.argv[0] || `{job="cortex-ops/query-frontend"} |= "logging.go" | logfmt | line_format "{{.msg}}"`
var logline = process.argv[1] || `level=debug ts=2020-10-02T10:10:42.092268913Z caller=logging.go:66 traceID=a9d4d8a928d8db1 msg="POST /api/prom/api/v1/query_range (200) 1.5s"`
try {
  const parsed = logql.parse(query, logline);
  console.log("query:", query);
  console.log("log:", logline);
  console.log("output:", parsed);
} catch(e){
  console.log(e);
}
