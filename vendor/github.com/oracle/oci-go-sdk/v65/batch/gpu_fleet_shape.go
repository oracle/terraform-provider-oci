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

// GpuFleetShape Shape of the GPU fleet. Describes hardware resources of each node in the fleet.
type GpuFleetShape interface {
}

type gpufleetshape struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *gpufleetshape) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergpufleetshape gpufleetshape
	s := struct {
		Model Unmarshalergpufleetshape
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *gpufleetshape) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FIXED_GPU_FLEET_SHAPE":
		mm := FixedGpuFleetShape{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GpuFleetShape: %s.", m.Type)
		return *m, nil
	}
}

func (m gpufleetshape) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m gpufleetshape) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GpuFleetShapeTypeEnum Enum with underlying type: string
type GpuFleetShapeTypeEnum string

// Set of constants representing the allowable values for GpuFleetShapeTypeEnum
const (
	GpuFleetShapeTypeFixedGpuFleetShape GpuFleetShapeTypeEnum = "FIXED_GPU_FLEET_SHAPE"
)

var mappingGpuFleetShapeTypeEnum = map[string]GpuFleetShapeTypeEnum{
	"FIXED_GPU_FLEET_SHAPE": GpuFleetShapeTypeFixedGpuFleetShape,
}

var mappingGpuFleetShapeTypeEnumLowerCase = map[string]GpuFleetShapeTypeEnum{
	"fixed_gpu_fleet_shape": GpuFleetShapeTypeFixedGpuFleetShape,
}

// GetGpuFleetShapeTypeEnumValues Enumerates the set of values for GpuFleetShapeTypeEnum
func GetGpuFleetShapeTypeEnumValues() []GpuFleetShapeTypeEnum {
	values := make([]GpuFleetShapeTypeEnum, 0)
	for _, v := range mappingGpuFleetShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGpuFleetShapeTypeEnumStringValues Enumerates the set of values in String for GpuFleetShapeTypeEnum
func GetGpuFleetShapeTypeEnumStringValues() []string {
	return []string{
		"FIXED_GPU_FLEET_SHAPE",
	}
}

// GetMappingGpuFleetShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGpuFleetShapeTypeEnum(val string) (GpuFleetShapeTypeEnum, bool) {
	enum, ok := mappingGpuFleetShapeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
