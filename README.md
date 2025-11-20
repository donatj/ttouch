# ttouch

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/ttouch)](https://goreportcard.com/report/github.com/donatj/ttouch)
[![GoDoc](https://godoc.org/github.com/donatj/ttouch?status.svg)](https://godoc.org/github.com/donatj/ttouch)
[![CI](https://github.com/donatj/ttouch/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/ttouch/actions/workflows/ci.yml)

Unix `touch` with JavaScript templates. Creates files with context-aware boilerplate instead of empty content.

## Features

- Generates content based on file extension, name, and project context
- JavaScript templates (ES2023) with file system access
- Per-project or global template customization
- Built-in templates for common file types
- Fast, lightweight Go binary

## Installation

```bash
go install github.com/donatj/ttouch/cmd/ttouch@latest
```

## Quick Start

Create a shell script:

```bash
ttouch script.sh
```

Generates:
```sh
#!/bin/sh

set -e

```

Create a Go file (detects package from existing files):

```bash
ttouch helper.go
```

Result with existing `package myapp`:
```go
package myapp

```

Create executable:

```bash
ttouch -e deploy.sh
```

## Usage

```
ttouch [flags] <file>...
```

Flags:
- `-f` - Overwrite if file exists
- `-e` - Make executable

Examples:

```bash
# Create new file
ttouch README.md

# Create multiple files
ttouch script.sh helper.go utils.js

# Create executable
ttouch -e deploy.py

# Overwrite existing file
ttouch -f existing-script.sh
```

## Templates

Uses [modernc.org/quickjs](https://pkg.go.dev/modernc.org/quickjs) JavaScript runtime (ES2023).

### Template Discovery

Templates live in `.ttouch` directories. Search starts in current directory and walks up to root (like `.git`).

Naming:
- `{ext}.js` - Matches extension (e.g., `py.js` for `*.py`)
- `{filename}.js` - Matches filename (e.g., `readme.md.js` for `README.md`)

Filename matches take precedence over extension matches.

Search order:
1. `.ttouch/{filename}.js` in current directory
2. `.ttouch/{filename}.js` in parent directories
3. `.ttouch/{ext}.js` in current directory
4. `.ttouch/{ext}.js` in parent directories
5. Built-in templates

### Writing Templates

Minimal template:

```js
"#!/bin/sh\n";
```

With IIFE:

```js
(function () {
	return "#!/bin/sh\n";
})();
```

Arrow function:

```js
(() => {
	return "#!/usr/bin/env python3\n\ndef main():\n    pass\n";
})();
```

Template value is the last expression in the file.

### VM Object

Templates access file information via `VM`:

```js
{
  "Filename": "script.sh",           // Relative path
  "AbsFilename": "/home/user/...",   // Absolute path
  "Flags": {                         // CLI flags
    "Executable": false,
    "Overwrite": true,
    "Files": ["script.sh"]
  }
}
```

Example:

```js
(() => {
    if (VM.Flags.Executable) {
        return "#!/usr/bin/env python3\n\n# Executable script\n";
    }
    return "# Python module\n";
})();
```

### Helper Functions

#### `ReadFile(path)`
Read file contents.

```js
const content = ReadFile("package.json");
const pkg = JSON.parse(content);
```

#### `Glob(pattern)`
Match files by pattern.

```js
const goFiles = Glob("*.go");
```

#### `ScanUp(filename)`
Find file in current or parent directories.

```js
const configs = ScanUp("package.json");
```

#### `SplitPath(path)`
Split path into `[directory, filename]`.

```js
const [dir, file] = SplitPath(VM.AbsFilename);
```

#### `Log(...args)`
Debug output.

```js
Log("Creating file:", VM.Filename);
```

### Built-in Templates

- [dot.js](https://github.com/donatj/ttouch/blob/master/templates/dot.js) - Dotfiles
- [go.js](https://github.com/donatj/ttouch/blob/master/templates/go.js) - Go files with package detection
- [md.js](https://github.com/donatj/ttouch/blob/master/templates/md.js) - Markdown with directory heading
- [php.js](https://github.com/donatj/ttouch/blob/master/templates/php.js) - PHP with PSR-4 namespaces
- [sh.js](https://github.com/donatj/ttouch/blob/master/templates/sh.js) - Shell scripts

## Custom Templates

Create `.ttouch/py.js`:

```js
(() => {
    let content = "#!/usr/bin/env python3\n";
    content += '"""Module docstring."""\n\n';
    
    if (VM.Filename.includes("test_")) {
        content += "import unittest\n\n";
        content += "class TestCase(unittest.TestCase):\n";
        content += "    def test_example(self):\n";
        content += "        pass\n\n";
        content += 'if __name__ == "__main__":\n';
        content += "    unittest.main()\n";
    } else {
        content += "def main():\n";
        content += "    pass\n\n";
        content += 'if __name__ == "__main__":\n';
        content += "    main()\n";
    }
    
    return content;
})();
```

Run `ttouch test_utils.py` for unittest template, `ttouch utils.py` for standard module.

## Use Cases

- Scaffold files with correct headers and imports
- Maintain consistent file structure across teams
- Generate boilerplate that adapts to existing code
- Learn language file conventions
- Automate multi-file generation

## Tips

- Use `Log()` for debugging
- Start with static templates, add logic later
- Review [built-in templates](https://github.com/donatj/ttouch/tree/master/templates) for examples
- Commit `.ttouch` directory for team sharing
- Place common templates in parent directories for project-wide use
