// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateGpuFleetShapeDetails Shape of the GPU fleet. Describes hardware resources of each node in the fleet.
type CreateGpuFleetShapeDetails interface {
}

type creategpufleetshapedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *creategpufleetshapedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreategpufleetshapedetails creategpufleetshapedetails
	s := struct {
		Model Unmarshalercreategpufleetshapedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *creategpufleetshapedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FIXED_GPU_FLEET_SHAPE":
		mm := CreateFixedGpuFleetShapeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateGpuFleetShapeDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m creategpufleetshapedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m creategpufleetshapedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateGpuFleetShapeDetailsTypeEnum Enum with underlying type: string
type CreateGpuFleetShapeDetailsTypeEnum string

// Set of constants representing the allowable values for CreateGpuFleetShapeDetailsTypeEnum
const (
	CreateGpuFleetShapeDetailsTypeFixedGpuFleetShape CreateGpuFleetShapeDetailsTypeEnum = "FIXED_GPU_FLEET_SHAPE"
)

var mappingCreateGpuFleetShapeDetailsTypeEnum = map[string]CreateGpuFleetShapeDetailsTypeEnum{
	"FIXED_GPU_FLEET_SHAPE": CreateGpuFleetShapeDetailsTypeFixedGpuFleetShape,
}

var mappingCreateGpuFleetShapeDetailsTypeEnumLowerCase = map[string]CreateGpuFleetShapeDetailsTypeEnum{
	"fixed_gpu_fleet_shape": CreateGpuFleetShapeDetailsTypeFixedGpuFleetShape,
}

// GetCreateGpuFleetShapeDetailsTypeEnumValues Enumerates the set of values for CreateGpuFleetShapeDetailsTypeEnum
func GetCreateGpuFleetShapeDetailsTypeEnumValues() []CreateGpuFleetShapeDetailsTypeEnum {
	values := make([]CreateGpuFleetShapeDetailsTypeEnum, 0)
	for _, v := range mappingCreateGpuFleetShapeDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateGpuFleetShapeDetailsTypeEnumStringValues Enumerates the set of values in String for CreateGpuFleetShapeDetailsTypeEnum
func GetCreateGpuFleetShapeDetailsTypeEnumStringValues() []string {
	return []string{
		"FIXED_GPU_FLEET_SHAPE",
	}
}

// GetMappingCreateGpuFleetShapeDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateGpuFleetShapeDetailsTypeEnum(val string) (CreateGpuFleetShapeDetailsTypeEnum, bool) {
	enum, ok := mappingCreateGpuFleetShapeDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
