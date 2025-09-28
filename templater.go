package ttouch

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/donatj/ttouch/templates"
	"modernc.org/quickjs"
)

type templater[T any] struct {
	Flags T
}

func New[T any](envflags T) *templater[T] {
	return &templater[T]{
		Flags: envflags,
	}
}

// ErrTemplateNotFound is returned when no template was found for the given
// type in any of the template locations
var ErrTemplateNotFound = errors.New("tempate not found")

func (t *templater[T]) GetTemplate(filename string) (string, error) {
	_, file := filepath.Split(filename)
	file = strings.ToLower(file)
	out, err := t.getTemplateFor(file+".js", filename)
	if !errors.Is(err, ErrTemplateNotFound) {
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

func (t *templater[T]) getTemplateFor(tmpFname, filename string) (string, error) {
	tpls := scanCwdUpForFile(filepath.Join(".ttouch", tmpFname))
	for _, tpl := range tpls {
		js, err := os.ReadFile(tpl)
		if err != nil {
			log.Println(tpl, err)
			continue
		}

		out, err := runJSTemplate(string(js), filename, t.Flags)
		if err != nil {
			return "", err
		}

		if out != "" {
			return out, nil
		}
	}

	js, _ := templates.Content.ReadFile(tmpFname)
	if js != nil && len(js) > 0 {
		out, err := runJSTemplate(string(js), filename, t.Flags)
		if err != nil {
			return "", err
		}

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

type JSFlags[T any] struct {
	Filename    string
	AbsFilename string
	Flags       T
}

func runJSTemplate[T any](js, filename string, vmflags T) (string, error) {
	vm, err := quickjs.NewVM()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Close()

	abs, _ := filepath.Abs(filename)

	vm.RegisterFunc("Log", log.Println, false)
	vm.RegisterFunc("ReadFile", jsReadfile, false)
	vm.RegisterFunc("Glob", jsGlob, false)
	vm.RegisterFunc("ScanUp", jsScanUp, false)
	vm.RegisterFunc("SplitPath", jsSplitpath, false)

	j, err := json.Marshal(&JSFlags[T]{
		Filename:    filename,
		AbsFilename: abs,
		Flags:       vmflags,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal JS flags: %w", err)
	}

	_, err = vm.Eval(fmt.Sprintf("const VM = %s;", j), quickjs.EvalGlobal)
	if err != nil {
		return "", fmt.Errorf("failed to define virtual machine flags: %w", err)
	}

	r, err := vm.Eval(js, quickjs.EvalGlobal)
	if err != nil {
		return "", fmt.Errorf("failed to run template: %w", err)
	}

	return fmt.Sprint(r), nil
}
