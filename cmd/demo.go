package main

import (
	"fmt"
	pack "github.com/mlavergn/gopack/src/gopack"
	"os"
)

func main() {
	fmt.Println("GoPack Demo")

	demo := 3

	pack := pack.NewPack()
	fmt.Println(pack.Container())
	switch demo {
	case 1:
		pack.Extract()
		break
	case 2:
		fmt.Println(pack.Load())
		fmt.Println(pack.String("cmd/index.html"))
		break
	case 3:
		fmt.Println(pack.Load())
		filePath, _ := pack.File("cmd/index.html")
		defer os.Remove(*filePath)
		fmt.Println(*filePath)
		break
	}
}
