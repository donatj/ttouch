package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/donatj/ttouch"
)

type envflags struct {
	Executable bool
	Files      []string
}

var envf = envflags{}

func init() {
	flag.BoolVar(&envf.Executable, "e", false, "mark the file executable")
	flag.Parse()

	envf.Files = flag.Args()
}

func main() {
	tmpr := ttouch.New(envf)

	for _, f := range envf.Files {
		_, err := os.Stat(f)
		if os.IsNotExist(err) {
			t, err := tmpr.GetTemplate(f)
			if err != nil && err != ttouch.ErrTemplateNotFound {
				log.Fatal(err)
			}

			mode := os.FileMode(0644)
			ioutil.WriteFile(f, []byte(t), mode)
		} else {
			// UPDATE MODIFIED LATER
		}
	}

}
