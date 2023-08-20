# Go Pack

[![Build Status](https://github.com/mlavergn/gopack/actions/workflows/build.yml/badge.svg)](https://github.com/mlavergn/gopack/actions/workflows/build.yml)
[![Go Report](https://goreportcard.com/badge/github.com/mlavergn/gopack)](https://goreportcard.com/report/github.com/mlavergn/gopack)
[![GoDoc](https://godoc.org/github.com/mlavergn/gopack/src/gopack?status.svg)](https://godoc.org/github.com/mlavergn/gopack/src/gopack)

Lightweight dependency-free embedding of static files into Go executables.

This is a more popular "embedding" type module here:

- [statik](https://github.com/rakyll/statik)

However, that implementation does not fit the use case I am targeting which is very resource constrained devices on readonly file systems. This implementation avoids external dependencies and aims for minimum overhead, features are intentonally minimal.

NOTE: Go Pack currently breaks when using code signing on macOS, there is a fix possible but it is incompatble with the existing packaging process.

## Implementation

The implementation assumes the following binary file structure

```text
executable + zip contents + zip size

offset 0
Executable Data
offset x - y
Zip Data
offset x - 10
Zip Size (y)
offset x
```

The logic attempts to find a string represented size marker (10 bytes) at the end of the Go executable. This marker is used
to calculate the offset of the zip contents from the end of the executable. The zip contents are optionally buffered
and used to access the static files in the zip contents or extracted to the directory containing the executable.

## Usage

For the demo, the following steps were used to generate the expected executable file format:

```bash
zip pack cmd/index.html
printf "%010d" `stat -f%z pack.zip` >> pack.zip
mv TheExecutable main.pack; cat main.pack pack.zip > TheExecutable
chmod +x TheExecutable
```

The API is simply:

```golang
package main

import "github.com/mlavergn/gopack/src/pack"

func main() {
    pack := gopack.NewPack()
    // A) extract to working directory
    pack.Extract()
    // -or-
    // B) read from memory buffer
    pack.Load()
    // b1) string value
    stringValue := pack.String("cmd/index.html")
    // b2) []byte value
    byteValue := pack.Bytes("cmd/index.html")
    // b3) pipe value (eg. http.resp)
    reader := pack.Pipe("cmd/index.html")
    ioutil.Copy(resp, reader)
}
```
