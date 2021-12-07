# LogQL Clickhouse UDF

Ladies and Gentlemen, meet the world's first Clickhouse `LogQL-UDF`

```
SELECT cloki_func('{job="any"} |= "logging.go" | logfmt | line_format "{{.msg}}"', s)
FROM
(
    SELECT 'level=debug ts=2020-10-02T10:10:42.092268913Z caller=logging.go:66 traceID=a9d4d8a928d8db1 msg="POST /api/prom/api/v1/query_range (200) 1.5s"' AS s
)

Query id: 392a022e-d112-4847-92f7-9282b3ad535d

┌─cloki_func('{job="any"} |= "logging.go" | logfmt | line_format "{{.msg}}"', s)─┐
│ POST /api/prom/api/v1/query_range (200) 1.5s                                   │
└────────────────────────────────────────────────────────────────────────────────┘

1 rows in set. Elapsed: 0.056 sec.
```
