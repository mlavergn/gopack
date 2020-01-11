[![Build Status](https://github.com/mlavergn/gopack/workflows/CI/badge.svg?branch=master)](https://github.com/mlavergn/gopack/actions)
[![Go Report](https://goreportcard.com/badge/github.com/mlavergn/gopack)](https://goreportcard.com/report/github.com/mlavergn/gopack)

# GO Pack

Lightweight dependency-free embedding of static files into Go executables

## Implementation

The implementation assumes the following binary file structure

```text
executable + zip contents + zip size
```

The logic attempts to find a string size marker (32 bit) at the end of the Go executable. This marker is used
to calculate the offset of the zip contents from the end of the executable. The zip contents are optionally buffered
and used to access the static files in the zip contents or extracted to the directory containing the executable.

## Usage

For the purposes of the included demo, the following steps generate the expected executable file format:

```bash
    zip pack cmd/index.html
    printf "%010d" `stat -f%z pack.zip` >> pack.zip
    mv TheExecutable main.pack; cat main.pack pack.zip > TheExecutable
    chmod +x TheExecutable
```

The API is basic:

```golang
    pack := gopack.NewPack()
    pack.Extract()
    // -or-
    pack.Load()
    reader := pack.String("cmd/index.html")
    // -for web-
    reader := pack.Pipe("cmd/index.html")
    ioutil.Copy(resp, )
```
