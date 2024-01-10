// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringSourceNamespaceDetails Discriminator for namespaces in the compartment-specific list.
type MonitoringSourceNamespaceDetails interface {
}

type monitoringsourcenamespacedetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *monitoringsourcenamespacedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitoringsourcenamespacedetails monitoringsourcenamespacedetails
	s := struct {
		Model Unmarshalermonitoringsourcenamespacedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitoringsourcenamespacedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "selected":
		mm := MonitoringSourceSelectedNamespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MonitoringSourceNamespaceDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m monitoringsourcenamespacedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m monitoringsourcenamespacedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoringSourceNamespaceDetailsKindEnum Enum with underlying type: string
type MonitoringSourceNamespaceDetailsKindEnum string

// Set of constants representing the allowable values for MonitoringSourceNamespaceDetailsKindEnum
const (
	MonitoringSourceNamespaceDetailsKindSelected MonitoringSourceNamespaceDetailsKindEnum = "selected"
)

var mappingMonitoringSourceNamespaceDetailsKindEnum = map[string]MonitoringSourceNamespaceDetailsKindEnum{
	"selected": MonitoringSourceNamespaceDetailsKindSelected,
}

var mappingMonitoringSourceNamespaceDetailsKindEnumLowerCase = map[string]MonitoringSourceNamespaceDetailsKindEnum{
	"selected": MonitoringSourceNamespaceDetailsKindSelected,
}

// GetMonitoringSourceNamespaceDetailsKindEnumValues Enumerates the set of values for MonitoringSourceNamespaceDetailsKindEnum
func GetMonitoringSourceNamespaceDetailsKindEnumValues() []MonitoringSourceNamespaceDetailsKindEnum {
	values := make([]MonitoringSourceNamespaceDetailsKindEnum, 0)
	for _, v := range mappingMonitoringSourceNamespaceDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoringSourceNamespaceDetailsKindEnumStringValues Enumerates the set of values in String for MonitoringSourceNamespaceDetailsKindEnum
func GetMonitoringSourceNamespaceDetailsKindEnumStringValues() []string {
	return []string{
		"selected",
	}
}

// GetMappingMonitoringSourceNamespaceDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoringSourceNamespaceDetailsKindEnum(val string) (MonitoringSourceNamespaceDetailsKindEnum, bool) {
	enum, ok := mappingMonitoringSourceNamespaceDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
