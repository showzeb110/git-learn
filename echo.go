package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[0:] {
		fmt.Print(idx)
		fmt.Println(" " + arg)
	}
}
