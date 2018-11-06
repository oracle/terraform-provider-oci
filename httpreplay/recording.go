// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

// +build record

package httpreplay

import "net/http"

var proxy *Proxy

// SetScenario creates a new proxy for this scenario
func SetScenario(name string) error {
	var err error
	if proxy, err = NewProxyAsMode(name, Recording); err == nil {
		debugLogf("Making a new proxy '%s' success", name)
	} else {
		debugLogf("Making a new proxy '%s' failed, %v", name, err)
	}
	return err
}

// SaveScenario saves the recorded service calls
func SaveScenario() error {
	debugLogf("Saving the proxy")
	err := proxy.Stop()
	proxy = nil
	return err
}

// InstallRecorder puts the recording transport into the http client, then returns a type that is compatible with the SDK's HTTPRequestDispatcher
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return InstallRecorderForRecodReplay(client, proxy)
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return false
}
