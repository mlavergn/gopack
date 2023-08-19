###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

.PHONY: test

VERSION := 1.2.0

ver:
	@sed -i '' 's/^const Version = "[0-9]\{1,3\}.[0-9]\{1,3\}.[0-9]\{1,3\}"/const Version = "${VERSION}"/' pack.go

format:
	go fmt

lint:
	go vet

build:
	go build pack.go

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
	go test -v ./...

github:
	open "https://github.com/mlavergn/gopack"

release:
	zip -r gopack.zip LICENSE README.md Makefile cmd test go.mod *.go
	hub release create -m "${VERSION} - GoPack" -a gopack.zip -t master "v${VERSION}"
	open "https://github.com/mlavergn/gopack/releases"

st:
	open -a SourceTree .