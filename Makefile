###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

GOPATH = "${PWD}"

lint:
	GOPATH=${GOPATH} ~/go/bin/golint .

build:
	GOPATH=${GOPATH} go build ./...

EXECFILE = "demo"
pack: build
	zip pack cmd/index.html
	printf "%010d" `stat -f%z pack.zip` >> pack.zip
	mv ${EXECFILE} main.pack; cat main.pack pack.zip > ${EXECFILE}
	chmod +x ${EXECFILE}
	rm pack.zip main.pack

demo: build
	GOPATH=${GOPATH} go build -o demo cmd/demo.go
	$(MAKE) pack
	./demo

clean:
	rm -f demo

test: build
	GOPATH=${GOPATH} go test -v ./src/...

github:
	open "https://github.com/mlavergn/gopack"