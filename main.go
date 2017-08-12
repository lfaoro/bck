package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var (
	log          = logrus.New()
	debug        = flag.Bool("debug", false, "activate debug mode.")
	restore      = flag.Bool("restore", false, "restore files in destinations to their origins.")
	settingsFlag = flag.String("settings", "$HOME/.backup.yml", "path to the settings file.")
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
	err := Sync(settings)
	if err != nil {
		log.Fatal(err)
	}

	if *restore {
		if err := Restore(settings); err != nil {
			log.Fatal(err)
		}
	}
}
