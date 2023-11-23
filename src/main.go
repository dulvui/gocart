// SPDX-FileCopyrightText: 2023 Simon Dalvai <info@simondalvai.org>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type LogEntry struct {
	Ts    float64
	Level string
}

func main() {
	filename, err := argsValidation(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Let's first read the `config.json` file
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	log.Printf("file: %s\n", file)

	var logEntry LogEntry
	err = json.Unmarshal(file, &logEntry)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("ts: %f\n", logEntry.Ts)
	log.Printf("level: %s\n", logEntry.Level)
}

func argsValidation(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("log file missing")
	}
	// if !strings.Contains(args[1], ".log") {
	// 	return "", errors.New("not a .log file")
	// }
	return args[1], nil
}
