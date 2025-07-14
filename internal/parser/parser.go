package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if err != nil {
			err = errors.Join(err, closeErr)
		}
	}(file)
	
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	
	//f, err := os.ReadFile(fileName)
	//if err != nil {
	//	return "", err
	//}
	
	strFile := string(data)
	
	return strFile, nil
}

func WriteFile(fileName string, input []string) (bool, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	
	defer func(file *os.File) {
		if closeErr := file.Close(); err != nil {
			err = errors.Join(err, closeErr)
		}
	}(file)
	
	for _, s := range input {
		if _, err := fmt.Fprint(file, s); err != nil {
			return false, err
		}
	}
	
	return true, nil
}
