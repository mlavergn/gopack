package main

import (
	"fmt"

	pack "github.com/mlavergn/gopack/src/gopack"
)

func main() {
	fmt.Println("GoPack Demo")

	extract := true

	pack := pack.NewPack()
	fmt.Println(pack.Container())
	if extract {
		pack.Extract()
	} else {
		fmt.Println(pack.Load())
		// reader := pack.Pipe("cmd/index.html")
		fmt.Println(pack.String("cmd/index.html"))
	}
}
