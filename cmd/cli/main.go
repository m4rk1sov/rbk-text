package main

import (
	"github.com/m4rk1sov/rbk-text/internal/jsonlog"
	"github.com/m4rk1sov/rbk-text/internal/parser"
	"github.com/m4rk1sov/rbk-text/internal/token"
	"log"
	"os"
)

//type application struct {
//	logger *jsonlog.Logger
//}

func main() {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed to open the logs file: %v", err)
	}
	defer func(logFile *os.File) {
		if err = logFile.Close(); err != nil {
			log.Fatalf("failed to close the logs file: %v", err)
		}
	}(logFile)

	logger := jsonlog.New(logFile, jsonlog.LevelTrace)
	logErr := jsonlog.New(os.Stderr, jsonlog.LevelError)

	if len(os.Args) < 3 {
		logErr.PrintFatal("Must use the 3 arguments (example: go run ./cmd/cli input.txt output.txt", nil)
	}
	argIn := os.Args[1]
	argOut := os.Args[2]

	input, err := parser.ReadFile(argIn)
	if err != nil {
		logger.PrintError("failed to read file", nil)
		logErr.PrintError("failed to read file", nil)
	}

	output, err := token.Tokenize(input)
	if err != nil {
		logger.PrintError("failed to tokenize the text", nil)
		logErr.PrintError("failed to tokenize the text", nil)
	}

	success, err := parser.WriteFile(argOut, output)
	if err != nil {
		logger.PrintError("failed to write tokens to a file", nil)
		logErr.PrintError("failed to write tokens to a file", nil)
	}
	if success {
		logger.PrintInfo("successfully wrote to a file", nil)
	}
}
