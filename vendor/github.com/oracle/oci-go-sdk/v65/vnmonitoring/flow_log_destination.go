// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlowLogDestination Where to store the flow logs.
type FlowLogDestination interface {
}

type flowlogdestination struct {
	JsonData        []byte
	DestinationType string `json:"destinationType"`
}

// UnmarshalJSON unmarshals json
func (m *flowlogdestination) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerflowlogdestination flowlogdestination
	s := struct {
		Model Unmarshalerflowlogdestination
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DestinationType = s.Model.DestinationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *flowlogdestination) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DestinationType {
	case "OBJECT_STORAGE":
		mm := FlowLogObjectStorageDestination{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FlowLogDestination: %s.", m.DestinationType)
		return *m, nil
	}
}

func (m flowlogdestination) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m flowlogdestination) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlowLogDestinationDestinationTypeEnum Enum with underlying type: string
type FlowLogDestinationDestinationTypeEnum string

// Set of constants representing the allowable values for FlowLogDestinationDestinationTypeEnum
const (
	FlowLogDestinationDestinationTypeObjectStorage FlowLogDestinationDestinationTypeEnum = "OBJECT_STORAGE"
)

var mappingFlowLogDestinationDestinationTypeEnum = map[string]FlowLogDestinationDestinationTypeEnum{
	"OBJECT_STORAGE": FlowLogDestinationDestinationTypeObjectStorage,
}

var mappingFlowLogDestinationDestinationTypeEnumLowerCase = map[string]FlowLogDestinationDestinationTypeEnum{
	"object_storage": FlowLogDestinationDestinationTypeObjectStorage,
}

// GetFlowLogDestinationDestinationTypeEnumValues Enumerates the set of values for FlowLogDestinationDestinationTypeEnum
func GetFlowLogDestinationDestinationTypeEnumValues() []FlowLogDestinationDestinationTypeEnum {
	values := make([]FlowLogDestinationDestinationTypeEnum, 0)
	for _, v := range mappingFlowLogDestinationDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFlowLogDestinationDestinationTypeEnumStringValues Enumerates the set of values in String for FlowLogDestinationDestinationTypeEnum
func GetFlowLogDestinationDestinationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingFlowLogDestinationDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlowLogDestinationDestinationTypeEnum(val string) (FlowLogDestinationDestinationTypeEnum, bool) {
	enum, ok := mappingFlowLogDestinationDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
