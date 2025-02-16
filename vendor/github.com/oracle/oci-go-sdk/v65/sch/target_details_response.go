// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetDetailsResponse An object that represents the target of the flow defined by the connector.
// An example target is a stream (Streaming service).
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type TargetDetailsResponse interface {
	GetPrivateEndpointMetadata() *PrivateEndpointMetadata
}

type targetdetailsresponse struct {
	JsonData                []byte
	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`
	Kind                    string                   `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *targetdetailsresponse) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetdetailsresponse targetdetailsresponse
	s := struct {
		Model Unmarshalertargetdetailsresponse
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PrivateEndpointMetadata = s.Model.PrivateEndpointMetadata
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetdetailsresponse) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "monitoring":
		mm := MonitoringTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "loggingAnalytics":
		mm := LoggingAnalyticsTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "functions":
		mm := FunctionsTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorage":
		mm := ObjectStorageTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "notifications":
		mm := NotificationsTargetDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TargetDetailsResponse: %s.", m.Kind)
		return *m, nil
	}
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m targetdetailsresponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m targetdetailsresponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetdetailsresponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetDetailsResponseKindEnum Enum with underlying type: string
type TargetDetailsResponseKindEnum string

// Set of constants representing the allowable values for TargetDetailsResponseKindEnum
const (
	TargetDetailsResponseKindFunctions        TargetDetailsResponseKindEnum = "functions"
	TargetDetailsResponseKindLogginganalytics TargetDetailsResponseKindEnum = "loggingAnalytics"
	TargetDetailsResponseKindMonitoring       TargetDetailsResponseKindEnum = "monitoring"
	TargetDetailsResponseKindNotifications    TargetDetailsResponseKindEnum = "notifications"
	TargetDetailsResponseKindObjectstorage    TargetDetailsResponseKindEnum = "objectStorage"
	TargetDetailsResponseKindStreaming        TargetDetailsResponseKindEnum = "streaming"
)

var mappingTargetDetailsResponseKindEnum = map[string]TargetDetailsResponseKindEnum{
	"functions":        TargetDetailsResponseKindFunctions,
	"loggingAnalytics": TargetDetailsResponseKindLogginganalytics,
	"monitoring":       TargetDetailsResponseKindMonitoring,
	"notifications":    TargetDetailsResponseKindNotifications,
	"objectStorage":    TargetDetailsResponseKindObjectstorage,
	"streaming":        TargetDetailsResponseKindStreaming,
}

var mappingTargetDetailsResponseKindEnumLowerCase = map[string]TargetDetailsResponseKindEnum{
	"functions":        TargetDetailsResponseKindFunctions,
	"logginganalytics": TargetDetailsResponseKindLogginganalytics,
	"monitoring":       TargetDetailsResponseKindMonitoring,
	"notifications":    TargetDetailsResponseKindNotifications,
	"objectstorage":    TargetDetailsResponseKindObjectstorage,
	"streaming":        TargetDetailsResponseKindStreaming,
}

// GetTargetDetailsResponseKindEnumValues Enumerates the set of values for TargetDetailsResponseKindEnum
func GetTargetDetailsResponseKindEnumValues() []TargetDetailsResponseKindEnum {
	values := make([]TargetDetailsResponseKindEnum, 0)
	for _, v := range mappingTargetDetailsResponseKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDetailsResponseKindEnumStringValues Enumerates the set of values in String for TargetDetailsResponseKindEnum
func GetTargetDetailsResponseKindEnumStringValues() []string {
	return []string{
		"functions",
		"loggingAnalytics",
		"monitoring",
		"notifications",
		"objectStorage",
		"streaming",
	}
}

// GetMappingTargetDetailsResponseKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDetailsResponseKindEnum(val string) (TargetDetailsResponseKindEnum, bool) {
	enum, ok := mappingTargetDetailsResponseKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
