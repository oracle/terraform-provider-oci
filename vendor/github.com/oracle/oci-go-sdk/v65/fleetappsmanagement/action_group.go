// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActionGroup Action Group.
type ActionGroup interface {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	GetDisplayName() *string
}

type actiongroup struct {
	JsonData    []byte
	DisplayName *string `mandatory:"false" json:"displayName"`
	Kind        string  `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *actiongroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleractiongroup actiongroup
	s := struct {
		Model Unmarshaleractiongroup
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *actiongroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "FLEET_USING_RUNBOOK":
		mm := FleetBasedActionGroup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ActionGroup: %s.", m.Kind)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m actiongroup) GetDisplayName() *string {
	return m.DisplayName
}

func (m actiongroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m actiongroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionGroupKindEnum Enum with underlying type: string
type ActionGroupKindEnum string

// Set of constants representing the allowable values for ActionGroupKindEnum
const (
	ActionGroupKindFleetUsingRunbook ActionGroupKindEnum = "FLEET_USING_RUNBOOK"
)

var mappingActionGroupKindEnum = map[string]ActionGroupKindEnum{
	"FLEET_USING_RUNBOOK": ActionGroupKindFleetUsingRunbook,
}

var mappingActionGroupKindEnumLowerCase = map[string]ActionGroupKindEnum{
	"fleet_using_runbook": ActionGroupKindFleetUsingRunbook,
}

// GetActionGroupKindEnumValues Enumerates the set of values for ActionGroupKindEnum
func GetActionGroupKindEnumValues() []ActionGroupKindEnum {
	values := make([]ActionGroupKindEnum, 0)
	for _, v := range mappingActionGroupKindEnum {
		values = append(values, v)
	}
	return values
}

// GetActionGroupKindEnumStringValues Enumerates the set of values in String for ActionGroupKindEnum
func GetActionGroupKindEnumStringValues() []string {
	return []string{
		"FLEET_USING_RUNBOOK",
	}
}

// GetMappingActionGroupKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionGroupKindEnum(val string) (ActionGroupKindEnum, bool) {
	enum, ok := mappingActionGroupKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
