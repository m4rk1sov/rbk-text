package main

import (
	"github.com/m4rk1sov/rbk-text/internal/jsonlog"
	"github.com/m4rk1sov/rbk-text/internal/parser"
	"github.com/m4rk1sov/rbk-text/internal/token"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed to open the logs file: %v", err)
	}
	defer func(logFile *os.File) {
		if err = logFile.Close(); err != nil {
			log.Fatalf("failed to close the logs file: %v", err)
		}
	}(f)

	logInfo := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	logFile := jsonlog.New(f, jsonlog.LevelTrace)
	logError := jsonlog.New(os.Stderr, jsonlog.LevelError)

	if len(os.Args) < 3 {
		logError.PrintFatal("Must use the 3 arguments (example: go run ./cmd/cli input.txt output.txt", nil)
	}
	argIn := os.Args[1]
	argOut := os.Args[2]

	raw, err := ReadFile(argIn)
	if err != nil {
		logFile.PrintError("failed to read file", map[string]string{"error: ": err.Error()})
		logError.PrintFatal("failed to read file", map[string]string{"error: ": err.Error()})
	}

	clean := token.Normalize(raw)

	tokens, err := token.Tokenize(clean)
	if err != nil {
		logFile.PrintError("failed to tokenize the text", map[string]string{"error: ": err.Error()})
		logError.PrintError("failed to tokenize the text", map[string]string{"error: ": err.Error()})
	}

	transformed, err := parser.TransformTokens(tokens)
	if err != nil {
		logFile.PrintError("failed to transform the text", map[string]string{"error: ": err.Error()})
		logError.PrintError("failed to transform the text", map[string]string{"error: ": err.Error()})
	}

	formatted := parser.JoinTokens(transformed)

	formattedNormalized := token.Normalize(formatted)

	err = WriteFile(argOut, formattedNormalized)
	if err != nil {
		logFile.PrintError("failed to write tokens to a file", map[string]string{"error: ": err.Error()})
		logError.PrintError("failed to write tokens to a file", map[string]string{"error: ": err.Error()})
	}
	logFile.PrintInfo("successfully wrote to a file", nil)
	logInfo.PrintInfo("successfully wrote to a file", nil)

}
