package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/donatj/ttouch"
)

type EnvFlags struct {
	Executable bool
	Overwrite  bool
	Files      []string
}

var envf = EnvFlags{}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <file>...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&envf.Executable, "e", false, "mark the out file(s) executable")
	flag.BoolVar(&envf.Overwrite, "o", false, "overwrite the file(s) if they exists")
	flag.Parse()

	envf.Files = flag.Args()
	if len(envf.Files) == 0 {
		fmt.Fprintln(os.Stderr, "error: at least one <file> is required")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	tmpr := ttouch.New(envf)

	for _, f := range tmpr.Flags.Files {
		_, err := os.Stat(f)
		if errors.Is(err, fs.ErrNotExist) {
			if !tmpr.Flags.Overwrite {
				continue
			}
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error: file %q - %v\n", f, err)
			continue
		}

		t, err := tmpr.GetTemplate(f)
		if err != nil && !errors.Is(err, ttouch.ErrTemplateNotFound) {
			log.Fatal(err)
		}

		mode := os.FileMode(0644)
		if envf.Executable {
			mode = os.FileMode(0755)
		}

		os.WriteFile(f, []byte(t), mode)

	}

}
