// Server1 is a minimal "echo" server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler Function echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Running Server on the Background; Open Terminal and Enter the Commands
// NOTE: Add the ampersand (&) as shown below if running on Mac OSX or Linux. Remove for Windows.
// $ go run ./<folders>/server1.go &
// NOTE: This also runs in tandem with the fetch.go program created earlier
// $ go build ./<folders>/server1.go
// $ ./fetch http://localhost:8000
// $ ./fetch http://localhost:8000/help
