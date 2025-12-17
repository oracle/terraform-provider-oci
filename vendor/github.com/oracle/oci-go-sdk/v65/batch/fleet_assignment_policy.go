// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetAssignmentPolicy A fleet assignment policy provides instructions to the system as to how a task should be assigned to a fleet,
// given it's minimum hardware requirements.
type FleetAssignmentPolicy interface {
}

type fleetassignmentpolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fleetassignmentpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfleetassignmentpolicy fleetassignmentpolicy
	s := struct {
		Model Unmarshalerfleetassignmentpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fleetassignmentpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FLEX_FIT":
		mm := FlexFitFleetAssignmentPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BEST_FIT":
		mm := BestFitFleetAssignmentPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FleetAssignmentPolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m fleetassignmentpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fleetassignmentpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetAssignmentPolicyTypeEnum Enum with underlying type: string
type FleetAssignmentPolicyTypeEnum string

// Set of constants representing the allowable values for FleetAssignmentPolicyTypeEnum
const (
	FleetAssignmentPolicyTypeBestFit FleetAssignmentPolicyTypeEnum = "BEST_FIT"
	FleetAssignmentPolicyTypeFlexFit FleetAssignmentPolicyTypeEnum = "FLEX_FIT"
)

var mappingFleetAssignmentPolicyTypeEnum = map[string]FleetAssignmentPolicyTypeEnum{
	"BEST_FIT": FleetAssignmentPolicyTypeBestFit,
	"FLEX_FIT": FleetAssignmentPolicyTypeFlexFit,
}

var mappingFleetAssignmentPolicyTypeEnumLowerCase = map[string]FleetAssignmentPolicyTypeEnum{
	"best_fit": FleetAssignmentPolicyTypeBestFit,
	"flex_fit": FleetAssignmentPolicyTypeFlexFit,
}

// GetFleetAssignmentPolicyTypeEnumValues Enumerates the set of values for FleetAssignmentPolicyTypeEnum
func GetFleetAssignmentPolicyTypeEnumValues() []FleetAssignmentPolicyTypeEnum {
	values := make([]FleetAssignmentPolicyTypeEnum, 0)
	for _, v := range mappingFleetAssignmentPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetAssignmentPolicyTypeEnumStringValues Enumerates the set of values in String for FleetAssignmentPolicyTypeEnum
func GetFleetAssignmentPolicyTypeEnumStringValues() []string {
	return []string{
		"BEST_FIT",
		"FLEX_FIT",
	}
}

// GetMappingFleetAssignmentPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetAssignmentPolicyTypeEnum(val string) (FleetAssignmentPolicyTypeEnum, bool) {
	enum, ok := mappingFleetAssignmentPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
