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

// ScriptedRestMonitorConfiguration Configuration details for the SCRIPTED_REST monitor type.
type ScriptedRestMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m ScriptedRestMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m ScriptedRestMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ScriptedRestMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptedRestMonitorConfiguration ScriptedRestMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScriptedRestMonitorConfiguration
	}{
		"SCRIPTED_REST_CONFIG",
		(MarshalTypeScriptedRestMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
