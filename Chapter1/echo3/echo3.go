// Simpler Echo Solution to Echo1 and Echo2
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(osArgs[1:])
}