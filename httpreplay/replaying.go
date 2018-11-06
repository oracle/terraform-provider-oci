// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

// +build replay

package httpreplay

import (
	"net/http"
)

var proxy *Proxy

// SetScenario creates a new proxy for this scenario
func SetScenario(name string) error {
	var err error
	if proxy, err = NewProxyAsMode(name, Replaying); err == nil {
		proxy.SetMatcher(matcher)
		proxy.SetTransformer(transformer)
	}
	return err
}

// SaveScenario does nothing when recording
func SaveScenario() error {
	proxy = nil
	return nil
}

// InstallRecorder puts the recording transport into the http client, then returns a type that is compatible with the SDK's HTTPRequestDispatcher
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return InstallRecorderForRecodReplay(client, proxy)
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return true
}
