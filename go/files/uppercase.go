package main

import (
	"log"
	"os"
	"strings"
)

func readFile(dir string) []byte {
	file, err := os.ReadFile(dir)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func writeFile(dir string, lines []byte) {
	str := strings.ToUpper(string(lines))

	err := os.WriteFile(dir, []byte(str), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	readDir := "./files/input.txt"
	writeDir := "./files/output.txt"

	lines := readFile(readDir)
	if len(lines) == 0 {
		log.Fatal("Empty file: ", readDir)
	}

	writeFile(writeDir, lines)
}
