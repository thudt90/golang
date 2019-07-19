// Echo prints its comman-line arguments

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))

	// os.Args[1:]

}
