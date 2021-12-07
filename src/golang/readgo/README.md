<img src="https://user-images.githubusercontent.com/1423657/139434383-98287329-74ce-4061-aabb-a19e500a986c.png" width=180 /><img src="https://img1.picmix.com/output/stamp/normal/5/9/9/7/1577995_45d25.gif" width=200/ >


# LogQL UDF for Clickhouse

> Native LogQL Clickhouse UDF

Ladies and Gentlemen, meet the world's first Clickhouse `LogQL-UDF`

```
SELECT cloki_func('{job="any"} |= "logging.go" | logfmt | line_format "{{.msg}}"', log) as out
FROM
(
    SELECT 'level=debug ts=2020-10-02T10:10:42.092268913Z caller=logging.go:66 traceID=a9d4d8a928d8db1 msg="POST /api/prom/api/v1/query_range (200) 1.5s"' AS log
) HAVING notEmpty(out)

Query id: 392a022e-d112-4847-92f7-9282b3ad535d

┌─cloki_func('{job="any"} |= "logging.go" | logfmt | line_format "{{.msg}}"', s)─┐
│ POST /api/prom/api/v1/query_range (200) 1.5s                                   │
└────────────────────────────────────────────────────────────────────────────────┘

1 rows in set. Elapsed: 0.056 sec.
```
