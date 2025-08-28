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

// FleetDetails Fleet Type
type FleetDetails interface {
}

type fleetdetails struct {
	JsonData  []byte
	FleetType string `json:"fleetType"`
}

// UnmarshalJSON unmarshals json
func (m *fleetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfleetdetails fleetdetails
	s := struct {
		Model Unmarshalerfleetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FleetType = s.Model.FleetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fleetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FleetType {
	case "GROUP":
		mm := GroupFleetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRODUCT":
		mm := ProductFleetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC":
		mm := GenericFleetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENVIRONMENT":
		mm := EnvironmentFleetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FleetDetails: %s.", m.FleetType)
		return *m, nil
	}
}

func (m fleetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fleetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetDetailsFleetTypeEnum Enum with underlying type: string
type FleetDetailsFleetTypeEnum string

// Set of constants representing the allowable values for FleetDetailsFleetTypeEnum
const (
	FleetDetailsFleetTypeGeneric     FleetDetailsFleetTypeEnum = "GENERIC"
	FleetDetailsFleetTypeProduct     FleetDetailsFleetTypeEnum = "PRODUCT"
	FleetDetailsFleetTypeEnvironment FleetDetailsFleetTypeEnum = "ENVIRONMENT"
	FleetDetailsFleetTypeGroup       FleetDetailsFleetTypeEnum = "GROUP"
)

var mappingFleetDetailsFleetTypeEnum = map[string]FleetDetailsFleetTypeEnum{
	"GENERIC":     FleetDetailsFleetTypeGeneric,
	"PRODUCT":     FleetDetailsFleetTypeProduct,
	"ENVIRONMENT": FleetDetailsFleetTypeEnvironment,
	"GROUP":       FleetDetailsFleetTypeGroup,
}

var mappingFleetDetailsFleetTypeEnumLowerCase = map[string]FleetDetailsFleetTypeEnum{
	"generic":     FleetDetailsFleetTypeGeneric,
	"product":     FleetDetailsFleetTypeProduct,
	"environment": FleetDetailsFleetTypeEnvironment,
	"group":       FleetDetailsFleetTypeGroup,
}

// GetFleetDetailsFleetTypeEnumValues Enumerates the set of values for FleetDetailsFleetTypeEnum
func GetFleetDetailsFleetTypeEnumValues() []FleetDetailsFleetTypeEnum {
	values := make([]FleetDetailsFleetTypeEnum, 0)
	for _, v := range mappingFleetDetailsFleetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetDetailsFleetTypeEnumStringValues Enumerates the set of values in String for FleetDetailsFleetTypeEnum
func GetFleetDetailsFleetTypeEnumStringValues() []string {
	return []string{
		"GENERIC",
		"PRODUCT",
		"ENVIRONMENT",
		"GROUP",
	}
}

// GetMappingFleetDetailsFleetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetDetailsFleetTypeEnum(val string) (FleetDetailsFleetTypeEnum, bool) {
	enum, ok := mappingFleetDetailsFleetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
