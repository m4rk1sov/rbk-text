package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) (string, error) {
	//file, err := os.Open(fileName)
	//if err != nil {
	//	return "", err
	//}
	//defer func(file *os.File) {
	//	closeErr := file.Close()
	//	if err != nil {
	//		err = errors.Join(err, closeErr)
	//	}
	//}(file)
	//
	//data, err := io.ReadAll(file)
	//if err != nil {
	//	return "", err
	//}
	
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	
	strFile := string(file)
	
	return strFile, nil
}

func WriteFile(fileName string, tokens []string) (bool, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	
	defer func(file *os.File) {
		if closeErr := file.Close(); err != nil {
			err = errors.Join(err, closeErr)
		}
	}(file)
	
	tokens = TransformTokens(tokens)
	
	for i, token := range tokens {
		if i > 0 && !isPunctuation(token) {
			_, err := fmt.Fprint(file, " ")
			if err != nil {
				return false, err
			}
		}
		_, err := fmt.Fprint(file, token)
		if err != nil {
			return false, err
		}
	}
	
	return true, nil
}

func isPunctuation(s string) bool {
	return len(s) == 1 && strings.ContainsRune(",.?!;:", rune(s[0]))
}
