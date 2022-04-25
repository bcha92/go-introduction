// This program uses recursion over HTML node tree to print the structure of the tree in the outline.
// As it encounters each element, it pushes the element's tag onto a stack and then prints the stack.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// 1. If you already have the "fetch.go" executable ready on root, skip Step #2a
// 2a. Open Terminal and type in:
// $ go build ./<folder names>/fetch.go
// 2b. Type in this command to the terminal:
// $ go build ./<folder names>/outline.go
// $ ./fetch https://golang.org | ./outline
