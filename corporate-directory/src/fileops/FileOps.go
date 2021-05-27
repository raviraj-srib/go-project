package fileops

import (
	"bufio"
	"log"
	"os"
)

const (
	INPUT_FILE_PATH = "../input/employee.txt"
)

func ReadFile() []string {
	file, err := os.Open(INPUT_FILE_PATH)

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
