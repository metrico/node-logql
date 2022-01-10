<img src='https://user-images.githubusercontent.com/1423657/147935343-598c7dfd-1412-4bad-9ac6-636994810443.png' style="margin-left:-10px" width=220>

# node-logql

> Native node binding for LogQL


### Build Module
```console
make
```

### Install
```
npm install -g node-logql
```

### Usage

See [example.js](https://raw.githubusercontent.com/metrico/node-logql/main/example.js) for a full query example

```javascript
const logql = require('node-logql');
const parsed = logql.parse(query, log);
```

Pipe stdin through [logql.js](https://raw.githubusercontent.com/metrico/node-logql/main/logql.js)
```console
echo "level=debug ts=2020-10-02T10:10:42.092268913Z ABC" | logql.js '{job="stdin"} | logfmt | line_format "{{.level}}"'
```

#### Todo
- [x] go binding
- [x] function mapping
- [ ] optimize size
