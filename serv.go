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

	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Invalid directory: %s\n", dir)
		os.Exit(1)
	}

	log.Printf("Serving from dir: %s\n", dir)

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}
