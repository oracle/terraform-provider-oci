// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceDetails An object that represents the source of the flow defined by the connector.
// An example source is the VCNFlow logs within the NetworkLogs group.
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.cloud.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type SourceDetails interface {
}

type sourcedetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *sourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcedetails sourcedetails
	s := struct {
		Model Unmarshalersourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "logging":
		mm := LoggingSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "monitoring":
		mm := MonitoringSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "plugin":
		mm := PluginSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SourceDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m sourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceDetailsKindEnum Enum with underlying type: string
type SourceDetailsKindEnum string

// Set of constants representing the allowable values for SourceDetailsKindEnum
const (
	SourceDetailsKindLogging    SourceDetailsKindEnum = "logging"
	SourceDetailsKindMonitoring SourceDetailsKindEnum = "monitoring"
	SourceDetailsKindStreaming  SourceDetailsKindEnum = "streaming"
	SourceDetailsKindPlugin     SourceDetailsKindEnum = "plugin"
)

var mappingSourceDetailsKindEnum = map[string]SourceDetailsKindEnum{
	"logging":    SourceDetailsKindLogging,
	"monitoring": SourceDetailsKindMonitoring,
	"streaming":  SourceDetailsKindStreaming,
	"plugin":     SourceDetailsKindPlugin,
}

var mappingSourceDetailsKindEnumLowerCase = map[string]SourceDetailsKindEnum{
	"logging":    SourceDetailsKindLogging,
	"monitoring": SourceDetailsKindMonitoring,
	"streaming":  SourceDetailsKindStreaming,
	"plugin":     SourceDetailsKindPlugin,
}

// GetSourceDetailsKindEnumValues Enumerates the set of values for SourceDetailsKindEnum
func GetSourceDetailsKindEnumValues() []SourceDetailsKindEnum {
	values := make([]SourceDetailsKindEnum, 0)
	for _, v := range mappingSourceDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceDetailsKindEnumStringValues Enumerates the set of values in String for SourceDetailsKindEnum
func GetSourceDetailsKindEnumStringValues() []string {
	return []string{
		"logging",
		"monitoring",
		"streaming",
		"plugin",
	}
}

// GetMappingSourceDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceDetailsKindEnum(val string) (SourceDetailsKindEnum, bool) {
	enum, ok := mappingSourceDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
