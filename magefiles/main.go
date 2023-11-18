//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/magefile/mage/mg"
)

func All() {
	mg.SerialDeps(startup, Before, Generate, After)
}

const LogPrefix = "[test-mage] "

var startTime time.Time // used to calculate target run time

func startup() error {
	startTime = time.Now()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LUTC | log.Lmsgprefix)
	log.SetOutput(os.Stdout) // simplify use of tee
	log.SetPrefix(LogPrefix)
	fmt.Println("Starting...")
	return nil
}
