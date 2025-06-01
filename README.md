# ttouch

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/ttouch)](https://goreportcard.com/report/github.com/donatj/ttouch)
[![GoDoc](https://godoc.org/github.com/donatj/ttouch?status.svg)](https://godoc.org/github.com/donatj/ttouch)

Unix touch with JavaScript driven templates.

## Installation

```bash
go install github.com/donatj/ttouch/cmd/ttouch@latest
```

## Templates

The templating engine runs on the [modernc.org/quickjs](https://pkg.go.dev/modernc.org/quickjs) JavaScript runtime which targets ES2023 currently.

Templates are searched for in `.ttouch` directories staring in your current
working directory down to the root, similar to how `.git` directories are searched for.

The templates are JavaScript files that return a string which is used as the content of the file being touched. The template can use the `file` variable to access information about the file being touched, such as its name, path, and extension. You can see examples of these in the [templates](https://github.com/donatj/ttouch/tree/master/templates) directory.

Templates are named all lowercase `{ext}.js` or `{filename}.js` - for example `md.js` for all `.md` files, or `readme.md.js` to match `README.md`. Filename matches take precidence over extension matches.

A very simple template could be:

```js
"#!/bin/sh\n";
```

If you want a little more control or clarity, you could wrap that in a IIFE.

```js
(function () {
	return "#!/bin/sh\n";
})();
```

The value of the template is the value of the last expression in the file.
