package main

import (
	"flag"

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
	*debug = false
	// dev mode

	if *debug {
		log.Level = logrus.DebugLevel
	}
}

func main() {

	settings := NewSettings()

	err := Sync(settings)
	if err != nil {
		log.Fatal(err)
	}
}
