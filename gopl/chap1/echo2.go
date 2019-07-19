// Echo1 prints its comman-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, step string
	s, step := "", ""
	// fmt.Println(len(os.Args))
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}
