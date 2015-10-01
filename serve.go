package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: serve [dir]\n")
		os.Exit(1)
	}

	dir := args[0]
	log.Printf("Serving from dir: %s\n", dir)

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}
