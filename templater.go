package ttouch

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/robertkrimen/otto"
)

//go:generate go-bindata -prefix templates -nomemcopy -pkg ttouch -o templates.go templates/...

type templater struct {
	envflags interface{}
}

func New(envflags interface{}) *templater {
	return &templater{
		envflags: envflags,
	}
}

// ErrTemplateNotFound is returned when no template was found for the given
// type in any of the template locations
var ErrTemplateNotFound = errors.New("tempate not found")

func (t *templater) GetTemplate(filename string) (string, error) {
	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext)
	ext = strings.Trim(ext, ". ")

	tmpFname := ext + ".js"

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	tpls := scanUpForTemplates(cwd, tmpFname)
	for _, tpl := range tpls {
		js, err := ioutil.ReadFile(tpl)
		if err != nil {
			log.Println(tpl, err)
			continue
		}

		out := runJSTemplate(string(js), filename, t.envflags)
		if out != "" {
			return out, nil
		}
	}

	js, _ := Asset(tmpFname)
	if js != nil && len(js) > 0 {
		out := runJSTemplate(string(js), filename, t.envflags)
		if out != "" {
			return out, nil
		}
	}

	return "", ErrTemplateNotFound
}

func scanUpForTemplates(dir, tmpFname string) []string {
	cwdParts := strings.Split(dir, string(os.PathSeparator))
	tmpls := []string{}

	for n := len(cwdParts) - 1; n >= 0; n-- {
		p := append([]string{"/"}, cwdParts[0:n+1]...)
		p = append(p, ".ttouch", tmpFname)

		tp := filepath.Join(p...)

		if _, err := os.Stat(tp); err == nil {
			tmpls = append(tmpls, tp)
		}
	}

	return tmpls
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
