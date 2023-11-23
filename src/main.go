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

type CaddyLog struct {
	Level   string
	Logger  string
	Message string
	Ts      float64
	Status  int
	Request Request
}

type Request struct {
	RemoteIp   string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	ClientIp   string `json:"client_ip"`
	Method     string
	Proto      string
	Host       string
	Uri        string
	Headers    Headers
}

type Headers struct {
	UserAgent []string `json:"User-Agent"`
}

func main() {
	filename, err := argsValidation(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	var caddyLog CaddyLog
	err = json.Unmarshal(file, &caddyLog)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("Ts: %f\n", caddyLog.Ts)
	log.Printf("Level: %s\n", caddyLog.Level)
	log.Printf("Status: %d\n", caddyLog.Status)
	log.Printf("Request: %s\n", caddyLog.Request)
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
