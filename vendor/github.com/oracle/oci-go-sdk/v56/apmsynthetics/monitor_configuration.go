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

// MonitorConfiguration Details of monitor configuration.
type MonitorConfiguration interface {

	// If isFailureRetried is enabled, then a failed call will be retried.
	GetIsFailureRetried() *bool
}

type monitorconfiguration struct {
	JsonData         []byte
	IsFailureRetried *bool  `mandatory:"false" json:"isFailureRetried"`
	ConfigType       string `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *monitorconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitorconfiguration monitorconfiguration
	s := struct {
		Model Unmarshalermonitorconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsFailureRetried = s.Model.IsFailureRetried
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitorconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "SCRIPTED_REST_CONFIG":
		mm := ScriptedRestMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCRIPTED_BROWSER_CONFIG":
		mm := ScriptedBrowserMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_CONFIG":
		mm := RestMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BROWSER_CONFIG":
		mm := BrowserMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsFailureRetried returns IsFailureRetried
func (m monitorconfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m monitorconfiguration) String() string {
	return common.PointerString(m)
}

// MonitorConfigurationConfigTypeEnum Enum with underlying type: string
type MonitorConfigurationConfigTypeEnum string

// Set of constants representing the allowable values for MonitorConfigurationConfigTypeEnum
const (
	MonitorConfigurationConfigTypeBrowserConfig         MonitorConfigurationConfigTypeEnum = "BROWSER_CONFIG"
	MonitorConfigurationConfigTypeScriptedBrowserConfig MonitorConfigurationConfigTypeEnum = "SCRIPTED_BROWSER_CONFIG"
	MonitorConfigurationConfigTypeRestConfig            MonitorConfigurationConfigTypeEnum = "REST_CONFIG"
	MonitorConfigurationConfigTypeScriptedRestConfig    MonitorConfigurationConfigTypeEnum = "SCRIPTED_REST_CONFIG"
)

var mappingMonitorConfigurationConfigType = map[string]MonitorConfigurationConfigTypeEnum{
	"BROWSER_CONFIG":          MonitorConfigurationConfigTypeBrowserConfig,
	"SCRIPTED_BROWSER_CONFIG": MonitorConfigurationConfigTypeScriptedBrowserConfig,
	"REST_CONFIG":             MonitorConfigurationConfigTypeRestConfig,
	"SCRIPTED_REST_CONFIG":    MonitorConfigurationConfigTypeScriptedRestConfig,
}

// GetMonitorConfigurationConfigTypeEnumValues Enumerates the set of values for MonitorConfigurationConfigTypeEnum
func GetMonitorConfigurationConfigTypeEnumValues() []MonitorConfigurationConfigTypeEnum {
	values := make([]MonitorConfigurationConfigTypeEnum, 0)
	for _, v := range mappingMonitorConfigurationConfigType {
		values = append(values, v)
	}
	return values
}
