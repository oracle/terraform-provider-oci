// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//go:build record
// +build record

package httpreplay

import "net/http"

var recorder *Recorder

// SetScenario creates a new recorder for this scenario
func SetScenario(name string) error {
	var err error
	if recorder, err = NewRecorderAsMode(name, ModeRecording); err == nil {
		debugLogf("Making a new recorder '%s' success", name)
	} else {
		debugLogf("Making a new recorder '%s' failed, %v", name, err)
	}
	return err
}

// SaveScenario saves the recorded service calls
func SaveScenario() error {
	debugLogf("Saving the recorder")
	err := recorder.Stop()
	recorder = nil
	return err
}

// InstallRecorder puts the recording transport into the http client, then returns a type that is compatible with the SDK's HTTPRequestDispatcher
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return InstallRecorderForRecodReplay(client, recorder)
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return false
}

// ModeRecordReplay returns true in record and replay
func ModeRecordReplay() bool {
	return true
}
