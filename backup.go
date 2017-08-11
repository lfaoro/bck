package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Settings struct {
	Origins      []string
	Destinations []string
	Options      []string
}

func NewSettings() *Settings {

	s := &Settings{}
	s.parseSettings()
	s.expandEnvs()
	s.sanityCheck()

	return s
}

// Sync will sync the data from origins to destionations.
func Sync(s *Settings) error {

	rsync := mustHave("rsync")

	flags := "--relative --copy-links --recursive --update --force --progress"
	origins := fmt.Sprintf("%s", strings.Join(s.Origins, " "))

	wg := &sync.WaitGroup{}
	for _, dst := range s.Destinations {
		input := fmt.Sprintf("%s %s %s %s", rsync, flags, origins, dst)
		wg.Add(1)
		go execCmd(input, wg)
	}
	wg.Wait()

	return nil

}

// Restore recovers data from the destionations and places it back into
// the origins.
func Restore(s *Settings) error {

	rsync := mustHave("rsync")

	flags := "--copy-links --recursive --update --force --progress"

	wg := &sync.WaitGroup{}
	var input string
	for _, o := range s.Origins {
		str := fmt.Sprintf("%s%s/", s.Destinations[0], o)
		input = fmt.Sprintf("%s %s %s %s", rsync, flags, str, o)

		wg.Add(1)
		go execCmd(input, wg)
	}

	wg.Wait()

	return nil
}

func (s *Settings) parseSettings() {
	file, err := os.Open(".backup.yml")
	if err != nil {
		log.Fatalf("unable to find \".backup.yml\" %s", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(data, s); err != nil {
		log.Fatal(err)
	}

}

func (s *Settings) expandEnvs() {
	for i, o := range s.Origins {
		s.Origins[i] = os.ExpandEnv(o)
	}
}

func (s *Settings) sanityCheck() {
	// TODO: get rid of goto
ORIG:
	for i, o := range s.Origins {
		_, err := os.Stat(o)
		if os.IsNotExist(err) {
			log.Errorf("excluding origin | %s", o)
			log.Debugf("removing elem: %s", s.Origins[i])
			s.Origins = append(s.Origins[:i], s.Origins[i+1:]...)
			goto ORIG
		}
	}

DEST:
	for i, o := range s.Destinations {
		str := os.ExpandEnv(o)
		_, err := os.Stat(str)
		if os.IsNotExist(err) {
			log.Errorf("excluding destination | %s", str)
			log.Debugf("removing elem: %s", s.Destinations[i])
			s.Destinations = append(s.Destinations[:i], s.Destinations[i+1:]...)
			goto DEST
		}
	}

	if len(s.Origins) == 0 {
		log.Fatal("no Origins exist")
	}
	if len(s.Destinations) == 0 {
		log.Fatal("no Destinations exist")
	}
}

func mustHave(bin string) string {
	// check if rsync command is available
	bin, err := exec.LookPath(bin)
	if err != nil {
		log.Fatal(err)
	}

	return bin
}

func execCmd(input string, wg *sync.WaitGroup) {

	fields := strings.Fields(input)
	cmd := exec.Command(fields[0], fields[1:len(fields)]...)
	log.Debug(cmd.Args)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Debug("\n", string(b))
		log.Fatal(err)
	}
	log.Infof("sync to %s | complete", fields[len(fields)-1])
	log.Debug("\n", string(b))

	wg.Done()
}
