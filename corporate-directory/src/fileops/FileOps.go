package fileops

import (
	"bufio"
	"log"
	"os"
)



func ReadFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var names []string

	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	file.Close()

	return names
}
