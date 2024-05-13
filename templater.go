package ttouch

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/donatj/ttouch/templates"
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

// ErrTemplateNotFound is returned when no template was found for the given
// type in any of the template locations
var ErrTemplateNotFound = errors.New("tempate not found")

func (t *templater) GetTemplate(filename string) (string, error) {
	_, file := filepath.Split(filename)
	file = strings.ToLower(file)
	out, err := t.getTemplateFor(file+".js", filename)
	if err != ErrTemplateNotFound {
		if err != nil {
			return "", err
		}

		return out, nil
	}

	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext)
	ext = strings.Trim(ext, ". ")

	return t.getTemplateFor(ext+".js", filename)
}

func (t *templater) getTemplateFor(tmpFname, filename string) (string, error) {
	tpls := scanCwdUpForFile(filepath.Join(".ttouch", tmpFname))
	for _, tpl := range tpls {
		js, err := os.ReadFile(tpl)
		if err != nil {
			log.Println(tpl, err)
			continue
		}

		out := runJSTemplate(string(js), filename, t.envflags)
		if out != "" {
			return out, nil
		}
	}

	js, _ := templates.Content.ReadFile(tmpFname)
	if js != nil && len(js) > 0 {
		out := runJSTemplate(string(js), filename, t.envflags)
		if out != "" {
			return out, nil
		}
	}

	return "", ErrTemplateNotFound
}

func scanCwdUpForFile(fname string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return []string{}
	}

	return scanUpForFile(cwd, fname)
}

func scanUpForFile(dir, fname string) []string {
	cwdParts := strings.Split(dir, string(os.PathSeparator))
	tmpls := []string{}

	for n := len(cwdParts) - 1; n >= 0; n-- {
		p := append([]string{"/"}, cwdParts[0:n+1]...)
		p = append(p, fname)

		tp := filepath.Join(p...)

		if _, err := os.Stat(tp); err == nil {
			tmpls = append(tmpls, tp)
		}
	}

	return tmpls
}

type JSFlags struct {
	Filename    string
	AbsFilename string
	Flags       interface{}
}

func runJSTemplate(js, filename string, vmflags interface{}) string {
	vm := otto.New()

	abs, _ := filepath.Abs(filename)

	vm.Set("VM", &JSFlags{
		Filename:    filename,
		AbsFilename: abs,
		Flags:       vmflags,
	})

	jsf := jsfuncs{vm}

	vm.Set("SplitPath", jsf.splitpath)
	vm.Set("ReadFile", jsf.jsreadfile)
	vm.Set("Glob", jsf.glob)
	vm.Set("ScanUp", jsf.scanup)

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
