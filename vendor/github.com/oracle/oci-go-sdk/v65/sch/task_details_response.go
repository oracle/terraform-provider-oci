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

// TaskDetailsResponse An object that represents a task within the flow defined by the connector.
// An example task is a filter for error logs.
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type TaskDetailsResponse interface {
	GetPrivateEndpointMetadata() *PrivateEndpointMetadata
}

type taskdetailsresponse struct {
	JsonData                []byte
	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`
	Kind                    string                   `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *taskdetailsresponse) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskdetailsresponse taskdetailsresponse
	s := struct {
		Model Unmarshalertaskdetailsresponse
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
func (m *taskdetailsresponse) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "logRule":
		mm := LogRuleTaskDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "function":
		mm := FunctionTaskDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TaskDetailsResponse: %s.", m.Kind)
		return *m, nil
	}
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m taskdetailsresponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m taskdetailsresponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m taskdetailsresponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskDetailsResponseKindEnum Enum with underlying type: string
type TaskDetailsResponseKindEnum string

// Set of constants representing the allowable values for TaskDetailsResponseKindEnum
const (
	TaskDetailsResponseKindFunction TaskDetailsResponseKindEnum = "function"
	TaskDetailsResponseKindLogrule  TaskDetailsResponseKindEnum = "logRule"
)

var mappingTaskDetailsResponseKindEnum = map[string]TaskDetailsResponseKindEnum{
	"function": TaskDetailsResponseKindFunction,
	"logRule":  TaskDetailsResponseKindLogrule,
}

var mappingTaskDetailsResponseKindEnumLowerCase = map[string]TaskDetailsResponseKindEnum{
	"function": TaskDetailsResponseKindFunction,
	"logrule":  TaskDetailsResponseKindLogrule,
}

// GetTaskDetailsResponseKindEnumValues Enumerates the set of values for TaskDetailsResponseKindEnum
func GetTaskDetailsResponseKindEnumValues() []TaskDetailsResponseKindEnum {
	values := make([]TaskDetailsResponseKindEnum, 0)
	for _, v := range mappingTaskDetailsResponseKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskDetailsResponseKindEnumStringValues Enumerates the set of values in String for TaskDetailsResponseKindEnum
func GetTaskDetailsResponseKindEnumStringValues() []string {
	return []string{
		"function",
		"logRule",
	}
}

// GetMappingTaskDetailsResponseKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskDetailsResponseKindEnum(val string) (TaskDetailsResponseKindEnum, bool) {
	enum, ok := mappingTaskDetailsResponseKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
