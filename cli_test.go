package ttouch_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var ttouchBin string

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "ttouch-bin-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmp)

	ttouchBin = filepath.Join(tmp, "ttouch")
	if runtime.GOOS == "windows" {
		ttouchBin += ".exe"
	}

	// Build the binary from the cmd/ttouch directory.
	build := exec.Command("go", "build", "-o", ttouchBin, "./cmd/ttouch")
	build.Stdout = os.Stderr
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		panic("failed to build ttouch: " + err.Error())
	}

	os.Exit(m.Run())
}

// runInDir runs the ttouch binary inside dir with the given extra args.
func runInDir(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command(ttouchBin, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("ttouch %v failed: %v\n%s", args, err, out)
	}
}

// readFile reads a file from dir and returns its content as a string.
func readFile(t *testing.T, dir, name string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(content)
}

// composerJSON is a minimal representation used to write fixture composer.json files.
type composerJSON struct {
	Autoload struct {
		PSR4 map[string]string `json:"psr-4,omitempty"`
		PSR0 map[string]string `json:"psr-0,omitempty"`
	} `json:"autoload"`
}

func writeComposerJSON(t *testing.T, dir string, psr4 map[string]string) {
	t.Helper()
	var cj composerJSON
	cj.Autoload.PSR4 = psr4
	data, err := json.MarshalIndent(cj, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "composer.json"), data, 0644); err != nil {
		t.Fatal(err)
	}
}

// ---- sh template ----

func TestShTemplate(t *testing.T) {
	dir := t.TempDir()
	runInDir(t, dir, "script.sh")
	got := readFile(t, dir, "script.sh")
	want := "#!/bin/sh\n\nset -e\n\n"
	if got != want {
		t.Errorf("sh template:\ngot:  %q\nwant: %q", got, want)
	}
}

// ---- dot template ----

func TestDotTemplate(t *testing.T) {
	dir := t.TempDir()
	runInDir(t, dir, "graph.dot")
	got := readFile(t, dir, "graph.dot")
	want := "digraph {\n\n}\n"
	if got != want {
		t.Errorf("dot template:\ngot:  %q\nwant: %q", got, want)
	}
}

// ---- go template ----

func TestGoTemplateMainInEmptyDir(t *testing.T) {
	dir := t.TempDir()
	runInDir(t, dir, "main.go")
	got := readFile(t, dir, "main.go")
	want := "package main\n\nfunc main() {\n\n}\n"
	if got != want {
		t.Errorf("go template (main.go, empty dir):\ngot:  %q\nwant: %q", got, want)
	}
}

func TestGoTemplateNonMainInEmptyDir(t *testing.T) {
	dir := t.TempDir()
	runInDir(t, dir, "util.go")
	got := readFile(t, dir, "util.go")
	want := "package main\n\n"
	if got != want {
		t.Errorf("go template (non-main.go, empty dir):\ngot:  %q\nwant: %q", got, want)
	}
}

func TestGoTemplateExistingPackage(t *testing.T) {
	dir := t.TempDir()
	// Seed the directory with an existing Go file declaring a named package.
	err := os.WriteFile(filepath.Join(dir, "existing.go"), []byte("package mypackage\n\nfunc Existing() {}\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	runInDir(t, dir, "newfile.go")
	got := readFile(t, dir, "newfile.go")
	want := "package mypackage\n\n"
	if got != want {
		t.Errorf("go template (existing package):\ngot:  %q\nwant: %q", got, want)
	}
}

// ---- md template ----

func TestMdTemplateHeadingFromDir(t *testing.T) {
	// The md template derives the heading from the name of the containing directory.
	parent := t.TempDir()
	dir := filepath.Join(parent, "myproject")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}
	runInDir(t, dir, "README.md")
	got := readFile(t, dir, "README.md")
	want := "# myproject\n\n"
	if got != want {
		t.Errorf("md template:\ngot:  %q\nwant: %q", got, want)
	}
}

// ---- php template ----

func TestPhpTemplateNoComposer(t *testing.T) {
	dir := t.TempDir()
	runInDir(t, dir, "MyClass.php")
	got := readFile(t, dir, "MyClass.php")
	want := "<?php\n\n"
	if got != want {
		t.Errorf("php template (no composer.json):\ngot:  %q\nwant: %q", got, want)
	}
}

func TestPhpTemplatePsr4Root(t *testing.T) {
	// composer.json at the project root; files directly inside src/ use the root namespace.
	root := t.TempDir()
	writeComposerJSON(t, root, map[string]string{"MyApp\\": "src/"})
	srcDir := filepath.Join(root, "src")
	if err := os.Mkdir(srcDir, 0755); err != nil {
		t.Fatal(err)
	}
	runInDir(t, srcDir, "MyClass.php")
	got := readFile(t, srcDir, "MyClass.php")
	want := "<?php\n\nnamespace MyApp;\n\n"
	if got != want {
		t.Errorf("php template (psr-4 root):\ngot:  %q\nwant: %q", got, want)
	}
}

func TestPhpTemplatePsr4Subdir(t *testing.T) {
	// Files in a subdirectory under the PSR-4 root get a derived sub-namespace.
	root := t.TempDir()
	writeComposerJSON(t, root, map[string]string{"MyApp\\": "src/"})
	subDir := filepath.Join(root, "src", "Sub")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	runInDir(t, subDir, "MyClass.php")
	got := readFile(t, subDir, "MyClass.php")
	want := "<?php\n\nnamespace MyApp\\Sub;\n\n"
	if got != want {
		t.Errorf("php template (psr-4 subdir):\ngot:  %q\nwant: %q", got, want)
	}
}
