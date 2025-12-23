package main

import (
	"embed"
	"fmt"
)

// with `go:embed PATH` it's possible to link static folder even the binary file moved

//go:embed static/*
var staticfiles embed.FS

func main() {
	dirEntries, err := staticfiles.ReadDir("static")
	//dirEntries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}
}
