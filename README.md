# ttouch

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/ttouch)](https://goreportcard.com/report/github.com/donatj/ttouch)
[![GoDoc](https://godoc.org/github.com/donatj/ttouch?status.svg)](https://godoc.org/github.com/donatj/ttouch)

Unix touch with JavaScript driven templates.

## Installation

```bash
go install github.com/donatj/ttouch/cmd/ttouch@latest
```

## Templates

The templating engine runs on the [otto](https://github.com/robertkrimen/otto) JavaScript runtime which targets ES5.

Templates are searched for in `.ttouch` directories staring in your current
working directory down to the root. 

Templates are named all lowercase `{ext}.js` or `{filename}.js` - for example `md.js` for all `.md` files, or `readme.md.js` to match `README.md`. Filename matches take precidence over extension matches.

The template string is the result of the final execution of the JS file. For example, a dead simple template could be:

```js
"#!/bin/sh\n";
```

If you want a little more control or clarity, you can wrap that in a function call:

```js
(function () {
	return "#!/bin/sh\n";
})();
```