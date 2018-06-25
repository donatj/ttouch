package templater

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/robertkrimen/otto"
)

type templater struct {
	envflags interface{}
}

func New(envflags interface{}) *templater {
	return &templater{
		envflags: envflags,
	}
}

//go:generate go-bindata -prefix ../templates -nomemcopy -pkg templater -o templates.go ../templates/...

func (t *templater) GetTemplate(filename string) string {
	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext)
	ext = strings.Trim(ext, ". ")

	js, _ := Asset(ext + ".js")
	if js != nil && len(js) > 0 {
		out := runJSTemplate(string(js), filename, t.envflags)
		if out != "" {
			return out
		}
	}

	return ""
}

type JSFlags struct {
	Filename string
	Flags    interface{}
}

func runJSTemplate(js, filename string, vmflags interface{}) string {
	vm := otto.New()

	vm.Set("VM", &JSFlags{
		Filename: filename,
		Flags:    vmflags,
	})

	vm.Set("Filename", filename)
	vm.Set("ReadFile", jsreadfile)
	vm.Set("Glob", jsglob{vm}.glob)

	v, err := vm.Run(js)
	if err != nil {
		log.Fatal(err)
	}

	s, err := v.ToString()
	if err != nil {
		log.Fatal(err)
	}

	return s
}
