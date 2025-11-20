# ttouch

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/ttouch)](https://goreportcard.com/report/github.com/donatj/ttouch)
[![GoDoc](https://godoc.org/github.com/donatj/ttouch?status.svg)](https://godoc.org/github.com/donatj/ttouch)
[![CI](https://github.com/donatj/ttouch/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/ttouch/actions/workflows/ci.yml)

**ttouch** is a command-line tool that creates files with intelligent, context-aware content using JavaScript-powered templates. Think of it as the Unix `touch` command, but instead of creating empty files, it generates files pre-populated with boilerplate code tailored to your project.

## What is ttouch?

ttouch automatically generates starter content for new files based on:
- **File extension** (e.g., `.js`, `.go`, `.py`, `.sh`)
- **File name** (e.g., `main.go`, `README.md`)
- **Project context** (e.g., package names from existing files, composer.json for PHP)

The templates are written in JavaScript (ES2023) and can be customized per-project or used globally via the built-in templates.

## Why Use ttouch?

- **Eliminate repetitive boilerplate**: Never type `#!/bin/bash` or `package main` again
- **Context-aware generation**: Templates can read your project structure and adapt accordingly
- **Highly customizable**: Write JavaScript templates with full access to file system operations
- **Project-specific templates**: Each project can have its own `.ttouch` directory with custom templates
- **Language agnostic**: Works with any programming language or file type
- **Fast and lightweight**: Built in Go with minimal dependencies

## Installation

```bash
go install github.com/donatj/ttouch/cmd/ttouch@latest
```

## Quick Start

Create a new shell script with built-in template:

```bash
ttouch -f script.sh
```

This generates:
```sh
#!/bin/sh

set -e

```

Create a Go file that inherits the package name from existing files:

```bash
ttouch -f helper.go
```

Result (if you have other `.go` files with `package myapp`):
```go
package myapp

```

Create an executable file:

```bash
ttouch -e -f deploy.sh
```

## Usage

```
ttouch [flags] <file>...
```

**Flags:**
- `-f` - Force overwrite if file exists (required to create new files)
- `-e` - Mark file(s) as executable

**Examples:**

```bash
# Create a new markdown file
ttouch -f README.md

# Create multiple files at once
ttouch -f script.sh helper.go utils.js

# Create executable Python script
ttouch -e -f deploy.py
```

## Templates

The templating engine runs on the [modernc.org/quickjs](https://pkg.go.dev/modernc.org/quickjs) JavaScript runtime which targets ES2023 currently.

### How Templates Work

Templates are searched for in `.ttouch` directories starting in your current working directory and walking up to the root, similar to how `.git` directories are searched. This allows you to:
- Have project-specific templates in your project's `.ttouch` directory
- Share templates across multiple projects by placing them in parent directories
- Use built-in templates when no custom template is found

Templates are named using lowercase conventions:
- `{ext}.js` - Matches all files with that extension (e.g., `py.js` matches `*.py`)
- `{filename}.js` - Matches specific filenames (e.g., `readme.md.js` matches `README.md`)

**Note**: Filename matches take precedence over extension matches.

### Template Search Order

1. `.ttouch/{filename}.js` in current directory (e.g., `.ttouch/readme.md.js`)
2. `.ttouch/{filename}.js` in parent directories
3. `.ttouch/{ext}.js` in current directory (e.g., `.ttouch/md.js`)
4. `.ttouch/{ext}.js` in parent directories
5. Built-in templates (embedded in ttouch binary)

### Simple Template Examples

A minimal template that returns a string:

```js
"#!/bin/sh\n";
```

Using an Immediately Invoked Function Expression (IIFE) for more control:

```js
(function () {
	return "#!/bin/sh\n";
})();
```

Or with ES6 arrow functions:

```js
(() => {
	return "#!/usr/bin/env python3\n\ndef main():\n    pass\n";
})();
```

**Important**: The value of the template is the value of the last expression in the file.

### The VM Object

Templates have access to a `VM` object containing file information:

```js
{
  "Filename": "script.sh",           // Relative filename as passed to ttouch
  "AbsFilename": "/home/user/...",   // Absolute path to the file
  "Flags": {                         // Command-line flags passed to ttouch
    "Executable": false,
    "Overwrite": true,
    "Files": ["script.sh"]
  }
}
```

Example using VM data:

```js
(() => {
    if (VM.Flags.Executable) {
        return "#!/usr/bin/env python3\n\n# Executable script\n";
    }
    return "# Python module\n";
})();
```

### Available JavaScript Functions

ttouch provides several helper functions for template logic:

#### `ReadFile(path)`
Reads and returns the contents of a file as a string.

```js
const content = ReadFile("package.json");
const pkg = JSON.parse(content);
```

#### `Glob(pattern)`
Returns an array of filenames matching the glob pattern.

```js
const goFiles = Glob("*.go");
if (goFiles.length > 0) {
    // There are existing Go files
}
```

#### `ScanUp(filename)`
Searches for a file by walking up directories from the current working directory, returns an array of matching paths.

```js
const configs = ScanUp("package.json");
if (configs.length > 0) {
    // Found package.json in current or parent directory
}
```

#### `SplitPath(path)`
Splits a path into directory and filename components, returns an array `[directory, filename]`.

```js
const [dir, file] = SplitPath(VM.AbsFilename);
```

#### `Log(...args)`
Logs messages (useful for debugging templates).

```js
Log("Creating file:", VM.Filename);
```

### Built-in Template Examples

ttouch includes several built-in templates. You can see examples of these in the [templates](https://github.com/donatj/ttouch/tree/master/templates) directory.

#### Shell Scripts (`sh.js`)

```js
`#!/bin/sh

set -e

`;
```

**Result**: Shell script with shebang and error exit enabled.

#### Markdown Files (`md.js`)

```js
(() => {
    const parts = SplitPath(VM.AbsFilename)[0].split(/[\/\\]/g);
    const name = parts[parts.length - 2];
    return `# ${name}\n\n`;
})();
```

**Result**: Markdown with h1 heading using parent directory name.

#### Go Files (`go.js`)

```js
(() => {
    const existing = Glob("*.go");
    let pkg = "main";
    if (existing && existing.length) {
        const content = ReadFile(existing[0]);
        if (content !== null) {
            const match = content.match(/^package (\w+)$/m);
            if (match !== null && match.length) {
                pkg = match[1];
            }
        }
    }
    let contents = "";
    if (pkg == "main" && VM.Filename == "main.go") {
        contents = `func main() {\n\n}\n`;
    }
    return `package ${pkg}

${contents}`;
})();
```

**Result**: Go file with package name detected from existing `.go` files, and a `main()` function if the file is `main.go` in package `main`.

#### PHP Files (`php.js`)

```js
(() => {
    // Reads composer.json to determine PSR-4 namespace
    const composerFiles = ScanUp("composer.json");
    let namespace = "";
    
    if (composerFiles.length > 0) {
        const content = ReadFile(composerFiles[0]);
        const composer = JSON.parse(content);
        // ... namespace detection logic ...
    }
    
    let result = "";
    if (VM.Flags.Executable) {
        result += "#!/usr/bin/env php\n";
    }
    result += "<?php\n\n";
    if (namespace !== "") {
        result += `namespace ${namespace};\n\n`;
    }
    return result;
})();
```

**Result**: PHP file with optional shebang, opening tag, and PSR-4 namespace detection from `composer.json`.

## Creating Custom Templates

Create a `.ttouch` directory in your project:

```bash
mkdir .ttouch
```

Create a template file, for example `.ttouch/py.js`:

```js
(() => {
    let content = "#!/usr/bin/env python3\n";
    content += '"""Module docstring."""\n\n';
    
    // Check if this is a test file
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

Now when you run `ttouch -f test_utils.py`, it will create a Python test file with unittest boilerplate, while `ttouch -f utils.py` creates a regular Python module.

## Real-World Use Cases

- **Project initialization**: Quickly scaffold new files with correct headers and imports
- **Team consistency**: Share templates via `.ttouch` in your repository to ensure consistent file structure
- **Smart boilerplate**: Templates adapt based on existing project files (package names, namespaces, etc.)
- **Language learning**: See proper file structure for languages you're learning
- **Automation**: Use in scripts to generate multiple files with appropriate content

## Tips

- **Template debugging**: Use `Log()` function to debug template execution
- **Start simple**: Begin with static templates and add logic as needed
- **Check built-ins**: Review the [built-in templates](https://github.com/donatj/ttouch/tree/master/templates) for inspiration
- **Version control**: Commit your `.ttouch` directory to share templates with your team
- **Hierarchical templates**: Place common templates in parent directories to share across multiple projects
