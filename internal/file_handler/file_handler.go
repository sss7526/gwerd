package file_handler

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"fmt"
)
func ReadLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return lines, nil
}

func ResolveFilePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ReadBlock(file *os.File) ([]string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return []string{string(content)}, nil

}