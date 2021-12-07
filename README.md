<img src="https://user-images.githubusercontent.com/1423657/139434383-98287329-74ce-4061-aabb-a19e500a986c.png" width=180 />

# node-logql

> Native node binding for LogQL


### Build Module
```console
make
```

### Usage

See [example.js](https://raw.githubusercontent.com/metrico/node-logql/main/example.js) for a full query example

```javascript
const logql = require('node-metricsql');
const parsed = logql.parse(query, log);
```

Pipe stdin through [logql.js](https://raw.githubusercontent.com/metrico/node-logql/main/logql.js)
```console
echo "level=debug ts=2020-10-02T10:10:42.092268913Z ABC" | node logql.js '{job="stdin"} | logfmt | line_format "{{.level}}"'
```

#### Todo
- [x] go binding
- [x] function mapping
- [ ] optimize size
