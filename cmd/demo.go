package main

import (
	"fmt"
	"os"

	pack "github.com/mlavergn/gopack"
)

func main() {
	fmt.Println("GoPack Demo")

	demo := 3

	pack := pack.NewPack()
	fmt.Println(pack.Container())
	switch demo {
	case 1:
		_, err := pack.Extract()
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	case 2:
		_, err := pack.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(pack.String("cmd/index.html"))
		break
	case 3:
		_, err := pack.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		filePath, _ := pack.File("cmd/index.html")
		defer os.Remove(*filePath)
		fmt.Println(*filePath)
		break
	}
}
