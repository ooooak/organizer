package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"../org"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please Specify path you want to organize.")
		os.Exit(0)
	}

	basePath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	org.Run(basePath)
}
