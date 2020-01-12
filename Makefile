###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

VERSION := 1.0.0

lint:
	golint .

build:
	go build ./...

EXECFILE = "demo"
pack: build
	zip pack cmd/index.html
	printf "%010d" `stat -f%z pack.zip` >> pack.zip
	mv ${EXECFILE} main.pack; cat main.pack pack.zip > ${EXECFILE}
	chmod +x ${EXECFILE}
	rm pack.zip main.pack

demo: build
	go build -o demo cmd/demo.go
	$(MAKE) pack
	cp demo test
	./demo

clean:
	rm -f demo

test: build
	go test -v ./src/...

github:
	open "https://github.com/mlavergn/gopack"

release:
	zip -r gopack.zip LICENSE README.md Makefile cmd src
	hub release create -m "${VERSION} - GoPack" -a gopack.zip -t master "v${VERSION}"
	open "https://github.com/mlavergn/gopack/releases"
