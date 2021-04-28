//dup1 prints the text of each line that appears more than once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter lines to check for duplicates (type q to stop entering): ")

	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "q" {
			break
		}
	}
	//Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
