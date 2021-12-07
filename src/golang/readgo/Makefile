NAME?=udf-reader

all:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o $(NAME) 
	#go build -a -ldflags '-extldflags "-static"' -o $(NAME) 

debug:
	go build -o $(NAME) 
	
modules:
	go get ./...

.PHONY: clean
clean:
	rm -fr $(NAME)
