// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//go:build !record && !replay
// +build !record,!replay

package httpreplay

import (
	"net/http"
)

// InstallRecorder does no-op.
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return client, nil
}

// SetScenario lets the recorder know which test is currently executing
func SetScenario(name string) error {
	debugLogf("Not recording. %s", name)
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

// ModeRecordReplay returns true in record and replay
func ModeRecordReplay() bool {
	return false
}
