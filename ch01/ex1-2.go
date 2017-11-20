// 1.2 - modify the echo program to also print os.Args[0], the name of the
// command that invoked it

package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args {
		j := string(i)
		s = s + j + " " + arg + "\r\n"
	}
	fmt.Println(s)
}
