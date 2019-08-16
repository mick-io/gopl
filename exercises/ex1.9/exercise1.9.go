// Exercise 1.9: Modify fetch to also print the HTTP status code, found in
// resp.Status.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http.Get returns the result in the response struct resp. The Body
		// field of resp contains the server response as a readable stream.
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		// ioutil.ReadAll reads the entire response; the result is stored in b.
		b, err := ioutil.ReadAll(resp.Body)
		// The Body stream is closed to avoid leaking resources.
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// Printf writes the response to the standard output.
		fmt.Printf("%s", b)
	}
}
