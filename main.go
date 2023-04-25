package main

import (
	"fmt"
	"os"

	"github.com/noah-spahn/prj-pkgsize"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		os.Exit(1)
	}

	rootDir := os.Args[1]

	// Compute the package sizes
	pkgSizes, err := pkgsize.ComputePackageSizes(rootDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the summary output
	for pkgName, pkgSize := range pkgSizes {
		fmt.Printf("%s: %d lines\n", pkgName, pkgSize)
	}
}
