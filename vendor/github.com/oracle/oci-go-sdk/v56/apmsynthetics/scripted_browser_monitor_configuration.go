// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ScriptedBrowserMonitorConfiguration Configuration details for the SCRIPTED_BROWSER monitor type.
type ScriptedBrowserMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	// If certificate validation is enabled, then the call will fail in case of certification errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m ScriptedBrowserMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m ScriptedBrowserMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ScriptedBrowserMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptedBrowserMonitorConfiguration ScriptedBrowserMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScriptedBrowserMonitorConfiguration
	}{
		"SCRIPTED_BROWSER_CONFIG",
		(MarshalTypeScriptedBrowserMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
