// SPDX-FileCopyrightText: 2023 Simon Dalvai <info@simondalvai.org>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"bufio"
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

// var wg sync.WaitGroup
var uriMap map[string]int

func main() {
	uriMap = make(map[string]int)

	filename, err := argsValidation(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	reader, _ := os.Open(filename)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// wg.Add(1)
		processLog(scanner.Text())
	}
	// wg.Wait()

	for key, value := range uriMap {
		log.Printf("%s counter: %d\n", key, value)
	}

}

func argsValidation(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("log file missing")
	}
	return args[1], nil
}

func processLog(logLine string) {
	var caddyLog CaddyLog
	err := json.Unmarshal([]byte(logLine), &caddyLog)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// log.Printf("Ts: %f\n", caddyLog.Ts)
	// log.Printf("Level: %s\n", caddyLog.Level)
	// log.Printf("Status: %d\n", caddyLog.Status)
	// log.Printf("Request: %s\n", caddyLog.Request)

	uriMap[caddyLog.Request.Uri]++
	// wg.Done()
}
