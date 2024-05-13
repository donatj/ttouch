package ttouch

import (
	"log"
	"os"
	"path/filepath"

	"github.com/robertkrimen/otto"
)

type jsfuncs struct {
	otto *otto.Otto
}

func (o jsfuncs) jsreadfile(call otto.FunctionCall) otto.Value {
	right, _ := call.Argument(0).ToString()

	result := otto.NullValue()

	c, err := os.ReadFile(right)
	if err == nil {
		result, err = otto.ToValue(string(c))
		if err != nil {
			log.Println(err)
		}
	}

	return result
}

func (o jsfuncs) glob(call otto.FunctionCall) otto.Value {
	right, _ := call.Argument(0).ToString()

	result := otto.NullValue()

	m, err := filepath.Glob(right)
	if err == nil {
		result, err = o.otto.ToValue(m)
		if err != nil {
			log.Println(err)
		}
	}

	return result
}

func (o jsfuncs) scanup(call otto.FunctionCall) otto.Value {
	right, _ := call.Argument(0).ToString()

	result, err := o.otto.ToValue(scanCwdUpForFile(right))
	if err != nil {
		log.Println(err)
	}

	return result
}

func (o jsfuncs) splitpath(call otto.FunctionCall) otto.Value {
	right, _ := call.Argument(0).ToString()

	d, f := filepath.Split(right)

	result, err := o.otto.ToValue([]string{d, f})
	if err != nil {
		log.Println(err)
	}

	return result
}
