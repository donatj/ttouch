// Code generated by go-bindata.
// sources:
// templates/go.js
// templates/php.js
// DO NOT EDIT!

package ttouch

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _goJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x4d\x4b\x33\x31\x14\x85\xd7\x93\x5f\x71\x08\x2f\x25\x79\x07\x12\x2d\xee\x86\x59\xaa\x2b\x41\x44\xdc\xb4\x15\xd2\x78\x9b\x0e\xa6\x89\x34\x99\xb6\x20\xfd\xef\x92\xd8\xfa\xb5\xbb\xe1\x3e\xe7\x3e\x39\x5a\xc3\x2c\x2d\x7a\x4c\xd1\x22\xef\xe3\xbd\x1f\x93\x98\xca\x8e\x69\x0d\x1b\x43\x8a\x9e\x94\x8f\x4e\xf0\xc7\x35\x61\x67\xfc\x48\x88\xab\x1a\x19\x12\x38\xda\x32\xca\x0e\x5a\xe3\xea\x6f\xe4\xe9\x4e\xdd\x78\xe3\x92\xba\x3e\x90\x1d\xb3\x59\x7a\x92\x1d\x63\x3b\xb3\xc5\x96\xd2\xe8\x33\x7a\x70\xad\xeb\x95\xc2\x0e\x9e\x82\xd9\x10\x5a\xf0\x79\x98\x07\x7e\x62\xe9\x30\xa4\x3c\x04\x87\x1e\xb7\x3e\x2e\x05\xff\xaf\x5c\xe4\x92\x31\x36\xac\xc4\xf7\x76\x32\xf9\x9a\x95\xa7\xe0\xf2\x1a\x12\xef\xac\x29\x27\x6c\x0c\x99\x42\xf1\x3d\x90\x79\x29\x22\x71\x66\x67\x17\x8b\xf2\xa9\x8a\x6d\x4c\xb6\x6b\xf4\x67\x5c\xd5\xb7\xd0\xcf\x6f\xc6\xbe\x1a\x47\x10\xf3\x7d\x2b\xff\xe9\x8d\xec\x58\x53\xdc\x75\xff\x4b\xd6\x9c\x8a\xb5\x3d\xf8\x39\x55\xea\x55\x72\x76\xb9\xf8\xd1\xad\x39\xb2\x23\x63\x9f\x7c\xf7\x11\x00\x00\xff\xff\x43\x8c\xce\x58\x86\x01\x00\x00"

func goJsBytes() ([]byte, error) {
	return bindataRead(
		_goJs,
		"go.js",
	)
}

func goJs() (*asset, error) {
	bytes, err := goJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go.js", size: 390, mode: os.FileMode(420), modTime: time.Unix(1529956720, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _phpJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4b\x2c\x52\x28\x4a\x2d\x2e\xcd\x29\x51\xb0\x55\x50\x52\xe2\xe2\xca\x4c\x53\xd0\x08\xf3\xd5\x73\xcb\x49\x4c\x2f\xd6\x73\xad\x48\x4d\x2e\x2d\x49\x4c\xca\x49\xd5\x54\xa8\xe6\xe2\x44\x28\x54\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x28\xc8\x28\x88\xc9\x53\xb2\xe6\xaa\xe5\xe2\x82\xaa\xd1\xb6\x55\x50\xb2\xb1\x07\x8b\x83\x65\xa0\xe2\xd6\x80\x00\x00\x00\xff\xff\x8f\x7e\x70\xc7\x70\x00\x00\x00"

func phpJsBytes() ([]byte, error) {
	return bindataRead(
		_phpJs,
		"php.js",
	)
}

func phpJs() (*asset, error) {
	bytes, err := phpJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "php.js", size: 112, mode: os.FileMode(420), modTime: time.Unix(1529958135, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"go.js": goJs,
	"php.js": phpJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"go.js": &bintree{goJs, map[string]*bintree{}},
	"php.js": &bintree{phpJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

