//go:build mage
// +build mage

package main

import (
	"log"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Generate() error {
	mg.SerialDeps(startup, Before)
	log.Println("Generating...")
	var _, err = os.OpenFile("generated.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return sh.Run("dir")
}
