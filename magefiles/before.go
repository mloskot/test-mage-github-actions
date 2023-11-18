//go:build mage
// +build mage

package main

import (
	"log"
	"os"

	"github.com/magefile/mage/mg"
)

func Before() error {
	mg.SerialDeps(startup)
	log.Println("Cleaning before...")

	_, err := os.Stat("generated.txt")
	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove("generated.txt")
	if err != nil {
		panic(err)
	}
	return nil
}
