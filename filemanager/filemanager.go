package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{inputPath, outputPath}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)
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

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputPath)

	var closeErr error
	defer func() {
		if cerr := file.Close(); cerr != nil {
			closeErr = fmt.Errorf("error closing file: %v", cerr)
		}
	}()

	if err != nil {
		return errors.New(fmt.Sprintf("Error creating file: %v", err))
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return errors.New(fmt.Sprintf("Error encoding file to json: %v", err))
	}

	if closeErr != nil {
		return closeErr
	}

	return nil
}
