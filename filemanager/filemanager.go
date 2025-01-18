package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error opening file: %v", err))
	}

	var closeErr error
	defer func() {
		if cerr := file.Close(); cerr != nil {
			closeErr = fmt.Errorf("error closing file: %v", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error scanning file: %v", err))
	}

	if closeErr != nil {
		return nil, closeErr
	}

	return lines, nil
}
