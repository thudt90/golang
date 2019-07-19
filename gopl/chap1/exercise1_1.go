// Echo prints its comman-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, step string
	fmt.Println(len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}
