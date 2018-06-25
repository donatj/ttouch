package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/donatj/ttouch/templater"
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
	tmpr := templater.New(envf)

	for _, f := range envf.Files {
		_, err := os.Stat(f)
		if os.IsNotExist(err) {
			t := tmpr.GetTemplate(f)

			mode := os.FileMode(0644)
			ioutil.WriteFile(f, []byte(t), mode)
		} else {
			// UPDATE MODIFIED LATER
		}
	}

}
