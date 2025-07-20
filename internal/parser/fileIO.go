package parser

import (
	"errors"
	"os"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	strFile := string(file)

	return strFile, nil
}

func WriteFile(fileName, input string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		if closeErr := file.Close(); err != nil {
			err = errors.Join(err, closeErr)
		}
	}(file)

	_, err = file.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}
