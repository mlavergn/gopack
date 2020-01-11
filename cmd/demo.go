package main

import (
	"fmt"
	"gopack"
)

func main() {
	fmt.Println("GoPack Demo")

	extract := true

	pack := gopack.NewPack()
	fmt.Println(pack.Executable())
	if extract {
		pack.Extract()
	} else {
		fmt.Println(pack.Load())
		// reader := pack.Pipe("cmd/index.html")
		fmt.Println(pack.String("cmd/index.html"))
	}
}
