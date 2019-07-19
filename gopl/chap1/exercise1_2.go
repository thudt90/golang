// Echo1 prints its comman-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("index:", i, "value:", os.Args[i])
	}
}
