// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fmt.Println("got here")
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("after inputs")
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Println(n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
