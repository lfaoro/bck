package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log   = logrus.New()
	debug = flag.Bool("debug", false, "activate debug mode.")
)

func init() {
	flag.Parse()
	log.Level = logrus.InfoLevel

	// dev mode
	*debug = true
	// dev mode

	if *debug {
		log.Level = logrus.DebugLevel
	}
}

func main() {

	settings := NewSettings()

	if err := Restore(settings); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

	err := Sync(settings)
	if err != nil {
		log.Fatal(err)
	}
}
