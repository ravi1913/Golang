//Echo1 prints command-line argumets.
package main

import (
	"fmt"
	"os"
)

func main() {
	//var s, sep string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
		//		s += sep + os.Args[i]
		//		sep = " "
	}
	//fmt.Println(s)
}
