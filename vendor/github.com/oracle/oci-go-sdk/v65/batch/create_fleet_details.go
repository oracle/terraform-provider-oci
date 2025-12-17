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

// CreateFleetDetails Fleet configuration of the batch context.
type CreateFleetDetails interface {
}

type createfleetdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createfleetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatefleetdetails createfleetdetails
	s := struct {
		Model Unmarshalercreatefleetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createfleetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SERVICE_MANAGED_FLEET":
		mm := CreateServiceManagedFleetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateFleetDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createfleetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createfleetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateFleetDetailsTypeEnum Enum with underlying type: string
type CreateFleetDetailsTypeEnum string

// Set of constants representing the allowable values for CreateFleetDetailsTypeEnum
const (
	CreateFleetDetailsTypeServiceManagedFleet CreateFleetDetailsTypeEnum = "SERVICE_MANAGED_FLEET"
)

var mappingCreateFleetDetailsTypeEnum = map[string]CreateFleetDetailsTypeEnum{
	"SERVICE_MANAGED_FLEET": CreateFleetDetailsTypeServiceManagedFleet,
}

var mappingCreateFleetDetailsTypeEnumLowerCase = map[string]CreateFleetDetailsTypeEnum{
	"service_managed_fleet": CreateFleetDetailsTypeServiceManagedFleet,
}

// GetCreateFleetDetailsTypeEnumValues Enumerates the set of values for CreateFleetDetailsTypeEnum
func GetCreateFleetDetailsTypeEnumValues() []CreateFleetDetailsTypeEnum {
	values := make([]CreateFleetDetailsTypeEnum, 0)
	for _, v := range mappingCreateFleetDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFleetDetailsTypeEnumStringValues Enumerates the set of values in String for CreateFleetDetailsTypeEnum
func GetCreateFleetDetailsTypeEnumStringValues() []string {
	return []string{
		"SERVICE_MANAGED_FLEET",
	}
}

// GetMappingCreateFleetDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFleetDetailsTypeEnum(val string) (CreateFleetDetailsTypeEnum, bool) {
	enum, ok := mappingCreateFleetDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
