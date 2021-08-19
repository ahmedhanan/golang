package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(rw, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":3000", nil)
}