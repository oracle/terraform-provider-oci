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

// TargetDetails An object that represents the target of the flow defined by the connector.
// An example target is a stream (Streaming service).
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.cloud.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type TargetDetails interface {
}

type targetdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *targetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetdetails targetdetails
	s := struct {
		Model Unmarshalertargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "notifications":
		mm := NotificationsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorage":
		mm := ObjectStorageTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "monitoring":
		mm := MonitoringTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "functions":
		mm := FunctionsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "loggingAnalytics":
		mm := LoggingAnalyticsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m targetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetDetailsKindEnum Enum with underlying type: string
type TargetDetailsKindEnum string

// Set of constants representing the allowable values for TargetDetailsKindEnum
const (
	TargetDetailsKindFunctions        TargetDetailsKindEnum = "functions"
	TargetDetailsKindLogginganalytics TargetDetailsKindEnum = "loggingAnalytics"
	TargetDetailsKindMonitoring       TargetDetailsKindEnum = "monitoring"
	TargetDetailsKindNotifications    TargetDetailsKindEnum = "notifications"
	TargetDetailsKindObjectstorage    TargetDetailsKindEnum = "objectStorage"
	TargetDetailsKindStreaming        TargetDetailsKindEnum = "streaming"
)

var mappingTargetDetailsKindEnum = map[string]TargetDetailsKindEnum{
	"functions":        TargetDetailsKindFunctions,
	"loggingAnalytics": TargetDetailsKindLogginganalytics,
	"monitoring":       TargetDetailsKindMonitoring,
	"notifications":    TargetDetailsKindNotifications,
	"objectStorage":    TargetDetailsKindObjectstorage,
	"streaming":        TargetDetailsKindStreaming,
}

var mappingTargetDetailsKindEnumLowerCase = map[string]TargetDetailsKindEnum{
	"functions":        TargetDetailsKindFunctions,
	"logginganalytics": TargetDetailsKindLogginganalytics,
	"monitoring":       TargetDetailsKindMonitoring,
	"notifications":    TargetDetailsKindNotifications,
	"objectstorage":    TargetDetailsKindObjectstorage,
	"streaming":        TargetDetailsKindStreaming,
}

// GetTargetDetailsKindEnumValues Enumerates the set of values for TargetDetailsKindEnum
func GetTargetDetailsKindEnumValues() []TargetDetailsKindEnum {
	values := make([]TargetDetailsKindEnum, 0)
	for _, v := range mappingTargetDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDetailsKindEnumStringValues Enumerates the set of values in String for TargetDetailsKindEnum
func GetTargetDetailsKindEnumStringValues() []string {
	return []string{
		"functions",
		"loggingAnalytics",
		"monitoring",
		"notifications",
		"objectStorage",
		"streaming",
	}
}

// GetMappingTargetDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDetailsKindEnum(val string) (TargetDetailsKindEnum, bool) {
	enum, ok := mappingTargetDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
