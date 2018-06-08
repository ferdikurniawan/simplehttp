package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	sourceDir := flag.String("d", "", "Source Dir")
	flag.Parse()

	if *sourceDir == "" {
		fmt.Println("Use -d to specify the source directory \n")
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.Dir("./"+*sourceDir)))
	fmt.Println("serving on :3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
