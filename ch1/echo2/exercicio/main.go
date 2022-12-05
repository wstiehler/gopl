package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// s, sep := "", ""

	// for _, arg := range os.Args[0:1] {
	// 	s += sep + arg
	// 	sep = ""
	// }
	// fmt.Printf("%s\n", s)
	fmt.Println(strings.Join(os.Args[1:], ""))
	fmt.Println(os.Args[0:1])
}
