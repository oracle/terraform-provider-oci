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

// Fleet Fleet configuration of the batch context.
type Fleet interface {
}

type fleet struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fleet) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfleet fleet
	s := struct {
		Model Unmarshalerfleet
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fleet) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SERVICE_MANAGED_FLEET":
		mm := ServiceManagedFleet{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Fleet: %s.", m.Type)
		return *m, nil
	}
}

func (m fleet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fleet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetTypeEnum Enum with underlying type: string
type FleetTypeEnum string

// Set of constants representing the allowable values for FleetTypeEnum
const (
	FleetTypeServiceManagedFleet FleetTypeEnum = "SERVICE_MANAGED_FLEET"
)

var mappingFleetTypeEnum = map[string]FleetTypeEnum{
	"SERVICE_MANAGED_FLEET": FleetTypeServiceManagedFleet,
}

var mappingFleetTypeEnumLowerCase = map[string]FleetTypeEnum{
	"service_managed_fleet": FleetTypeServiceManagedFleet,
}

// GetFleetTypeEnumValues Enumerates the set of values for FleetTypeEnum
func GetFleetTypeEnumValues() []FleetTypeEnum {
	values := make([]FleetTypeEnum, 0)
	for _, v := range mappingFleetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetTypeEnumStringValues Enumerates the set of values in String for FleetTypeEnum
func GetFleetTypeEnumStringValues() []string {
	return []string{
		"SERVICE_MANAGED_FLEET",
	}
}

// GetMappingFleetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetTypeEnum(val string) (FleetTypeEnum, bool) {
	enum, ok := mappingFleetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
