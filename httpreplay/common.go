// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package httpreplay

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func init() {
	SetDebugLogger(DefaultLogger())
}

// HTTPRecordingClient wraps the execution of a http request, adding record/replay functionality.  It is meant to be compatible with oci-go-sdk's client.HTTPRequestDispatcher interface.
type HTTPRecordingClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var currentLogger *log.Logger

// DebugLogf logs a formatted string to stderr.  It should be considered temporary, and in a release version either removed or replaced with a passed-in logging interface.
func debugLogf(format string, v ...interface{}) {
	if currentLogger != nil {
		err := currentLogger.Output(2, fmt.Sprintf(format, v...))
		if err != nil {
			log.Println("[WARN] Failed to write to current logger")
		}
	}
}

func SetDebugLogger(logger *log.Logger) {
	currentLogger = logger
}

func DefaultLogger() *log.Logger {
	return log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func saveOrLog(d interface{}, fn string) {
	if err := save(d, fn); err != nil {
		debugLogf("Error: %v", err)
	}
}

// Save writes some data on disk for future re-use
func save(d interface{}, fn string) error {
	// Create directory for scenario if missing
	scenarioDir := filepath.Dir(fn)
	if _, err := os.Stat(scenarioDir); os.IsNotExist(err) {
		if err = os.MkdirAll(scenarioDir, 0755); err != nil {
			return err
		}
	}

	// Marshal to YAML and save interactions
	data, err := yaml.Marshal(d)
	if err != nil {
		return err
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Honor the YAML structure specification
	// http://www.yaml.org/spec/1.2/spec.html#id2760395
	_, err = f.Write([]byte("---\n"))
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
