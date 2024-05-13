package templates

import "embed"

// Content is the embedded content of the templates directory.
//
//go:embed *.js
var Content embed.FS
