package ttouch

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/robertkrimen/otto"
)

func jsreadfile(call otto.FunctionCall) otto.Value {
	right, _ := call.Argument(0).ToString()

	result := otto.NullValue()

	c, err := ioutil.ReadFile(right)
	if err == nil {
		result, err = otto.ToValue(string(c))
		if err != nil {
			log.Println(err)
		}
	}

	return result
}

type jsglob struct {
	otto *otto.Otto
}

func (o jsglob) glob(call otto.FunctionCall) otto.Value {
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
