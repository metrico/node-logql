default: build

compile-go:
	go build -ldflags "-s -w" -buildmode=c-shared -o logql.so src/golang/logql.go
	strip logql.so

build: compile-go
	npm run build
