//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"os"
	// mg contains helpful utility functions, like Deps
)

// Delete generated files.
func Clean() {
	fmt.Println("Cleaning...")
	log.Println("Deleting README.generated.md")
	os.RemoveAll("README.generated.md")
}
