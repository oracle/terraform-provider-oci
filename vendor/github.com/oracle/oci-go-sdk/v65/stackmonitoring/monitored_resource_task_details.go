// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoredResourceTaskDetails The request details for the performing the task.
type MonitoredResourceTaskDetails interface {
}

type monitoredresourcetaskdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *monitoredresourcetaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitoredresourcetaskdetails monitoredresourcetaskdetails
	s := struct {
		Model Unmarshalermonitoredresourcetaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitoredresourcetaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "IMPORT_OCI_TELEMETRY_RESOURCES":
		mm := ImportOciTelemetryResourcesTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MonitoredResourceTaskDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m monitoredresourcetaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m monitoredresourcetaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoredResourceTaskDetailsTypeEnum Enum with underlying type: string
type MonitoredResourceTaskDetailsTypeEnum string

// Set of constants representing the allowable values for MonitoredResourceTaskDetailsTypeEnum
const (
	MonitoredResourceTaskDetailsTypeImportOciTelemetryResources MonitoredResourceTaskDetailsTypeEnum = "IMPORT_OCI_TELEMETRY_RESOURCES"
)

var mappingMonitoredResourceTaskDetailsTypeEnum = map[string]MonitoredResourceTaskDetailsTypeEnum{
	"IMPORT_OCI_TELEMETRY_RESOURCES": MonitoredResourceTaskDetailsTypeImportOciTelemetryResources,
}

var mappingMonitoredResourceTaskDetailsTypeEnumLowerCase = map[string]MonitoredResourceTaskDetailsTypeEnum{
	"import_oci_telemetry_resources": MonitoredResourceTaskDetailsTypeImportOciTelemetryResources,
}

// GetMonitoredResourceTaskDetailsTypeEnumValues Enumerates the set of values for MonitoredResourceTaskDetailsTypeEnum
func GetMonitoredResourceTaskDetailsTypeEnumValues() []MonitoredResourceTaskDetailsTypeEnum {
	values := make([]MonitoredResourceTaskDetailsTypeEnum, 0)
	for _, v := range mappingMonitoredResourceTaskDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredResourceTaskDetailsTypeEnumStringValues Enumerates the set of values in String for MonitoredResourceTaskDetailsTypeEnum
func GetMonitoredResourceTaskDetailsTypeEnumStringValues() []string {
	return []string{
		"IMPORT_OCI_TELEMETRY_RESOURCES",
	}
}

// GetMappingMonitoredResourceTaskDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredResourceTaskDetailsTypeEnum(val string) (MonitoredResourceTaskDetailsTypeEnum, bool) {
	enum, ok := mappingMonitoredResourceTaskDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
