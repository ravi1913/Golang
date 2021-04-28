//Uses the inbuilt string join function
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
