package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
			// Another way to handle error
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oops"))
			// return
		}

		// log.Printf("Data: %s\n", d)
		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
