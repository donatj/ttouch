package ttouch

import (
	"log"
	"os"
	"path/filepath"
)

func jsReadfile(right string) string {
	c, err := os.ReadFile(right)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(c)
}

func jsGlob(right string) []string {
	m, err := filepath.Glob(right)
	if err != nil {
		log.Println(err)
	}
	return m
}

func jsScanUp(right string) []string {
	return scanCwdUpForFile(right)
}

func jsSplitpath(right string) []string {
	d, f := filepath.Split(right)
	return []string{d, f}
}
