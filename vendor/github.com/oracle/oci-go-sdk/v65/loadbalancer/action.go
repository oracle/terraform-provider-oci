// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Action An entity that represents an action to apply for a routing rule.
type Action interface {
}

type action struct {
	JsonData []byte
	Name     string `json:"name"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Name {
	case "FORWARD_TO_BACKENDSET":
		mm := ForwardToBackendSet{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Action: %s.", m.Name)
		return *m, nil
	}
}

func (m action) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m action) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionNameEnum Enum with underlying type: string
type ActionNameEnum string

// Set of constants representing the allowable values for ActionNameEnum
const (
	ActionNameForwardToBackendset ActionNameEnum = "FORWARD_TO_BACKENDSET"
)

var mappingActionNameEnum = map[string]ActionNameEnum{
	"FORWARD_TO_BACKENDSET": ActionNameForwardToBackendset,
}

var mappingActionNameEnumLowerCase = map[string]ActionNameEnum{
	"forward_to_backendset": ActionNameForwardToBackendset,
}

// GetActionNameEnumValues Enumerates the set of values for ActionNameEnum
func GetActionNameEnumValues() []ActionNameEnum {
	values := make([]ActionNameEnum, 0)
	for _, v := range mappingActionNameEnum {
		values = append(values, v)
	}
	return values
}

// GetActionNameEnumStringValues Enumerates the set of values in String for ActionNameEnum
func GetActionNameEnumStringValues() []string {
	return []string{
		"FORWARD_TO_BACKENDSET",
	}
}

// GetMappingActionNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionNameEnum(val string) (ActionNameEnum, bool) {
	enum, ok := mappingActionNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
