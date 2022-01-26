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

// UpdateMonitorDetails Details of the request body used to update a monitor.
type UpdateMonitorDetails struct {

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A list of vantage points from which to execute the monitor.
	// Use /publicVantagePoints to fetch public vantage points.
	VantagePoints []string `mandatory:"false" json:"vantagePoints"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the script.
	// scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null.
	ScriptId *string `mandatory:"false" json:"scriptId"`

	// Enables or disables the monitor.
	Status MonitorStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Interval in seconds after the start time when the job should be repeated.
	// Minimum repeatIntervalInSeconds should be 300 seconds.
	RepeatIntervalInSeconds *int `mandatory:"false" json:"repeatIntervalInSeconds"`

	// If runOnce is enabled, then the monitor will run once.
	IsRunOnce *bool `mandatory:"false" json:"isRunOnce"`

	// Timeout in seconds. Timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors.
	// Also, timeoutInSeconds should be a multiple of 60. Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Specify the endpoint on which to run the monitor.
	// For BROWSER and REST monitor types, target is mandatory.
	// If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint.
	// If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is.
	Target *string `mandatory:"false" json:"target"`

	// List of script parameters in the monitor.
	// This is valid only for SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null.
	// Example: `[{"paramName": "userid", "paramValue":"testuser"}]`
	ScriptParameters []MonitorScriptParameter `mandatory:"false" json:"scriptParameters"`

	Configuration MonitorConfiguration `mandatory:"false" json:"configuration"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMonitorDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMonitorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		VantagePoints           []string                          `json:"vantagePoints"`
		ScriptId                *string                           `json:"scriptId"`
		Status                  MonitorStatusEnum                 `json:"status"`
		RepeatIntervalInSeconds *int                              `json:"repeatIntervalInSeconds"`
		IsRunOnce               *bool                             `json:"isRunOnce"`
		TimeoutInSeconds        *int                              `json:"timeoutInSeconds"`
		Target                  *string                           `json:"target"`
		ScriptParameters        []MonitorScriptParameter          `json:"scriptParameters"`
		Configuration           monitorconfiguration              `json:"configuration"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.VantagePoints = make([]string, len(model.VantagePoints))
	for i, n := range model.VantagePoints {
		m.VantagePoints[i] = n
	}

	m.ScriptId = model.ScriptId

	m.Status = model.Status

	m.RepeatIntervalInSeconds = model.RepeatIntervalInSeconds

	m.IsRunOnce = model.IsRunOnce

	m.TimeoutInSeconds = model.TimeoutInSeconds

	m.Target = model.Target

	m.ScriptParameters = make([]MonitorScriptParameter, len(model.ScriptParameters))
	for i, n := range model.ScriptParameters {
		m.ScriptParameters[i] = n
	}

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(MonitorConfiguration)
	} else {
		m.Configuration = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
