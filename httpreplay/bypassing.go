// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

// +build !record
// +build !replay

package httpreplay

import (
	"net/http"
)

// InstallRecorder in bypassing mode return original client
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return client, nil
}

// SetScenario lets the proxy know which test is currently executing
func SetScenario(name string) error {
	debugLogf("Not recording.", name)
	return nil
}

// SaveScenario saves the recorded service calls for the current scenario
func SaveScenario() error {
	return nil
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return false
}
