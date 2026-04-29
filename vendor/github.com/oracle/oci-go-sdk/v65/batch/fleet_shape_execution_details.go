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

// FleetShapeExecutionDetails Details about the shape which was used for the task execution.
type FleetShapeExecutionDetails interface {
}

type fleetshapeexecutiondetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fleetshapeexecutiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfleetshapeexecutiondetails fleetshapeexecutiondetails
	s := struct {
		Model Unmarshalerfleetshapeexecutiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fleetshapeexecutiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CPU_FLEET_SHAPE_EXECUTION_DETAILS":
		mm := CpuFleetShapeExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GPU_FLEET_SHAPE_EXECUTION_DETAILS":
		mm := GpuFleetShapeExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FleetShapeExecutionDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m fleetshapeexecutiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fleetshapeexecutiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetShapeExecutionDetailsTypeEnum Enum with underlying type: string
type FleetShapeExecutionDetailsTypeEnum string

// Set of constants representing the allowable values for FleetShapeExecutionDetailsTypeEnum
const (
	FleetShapeExecutionDetailsTypeCpuFleetShapeExecutionDetails FleetShapeExecutionDetailsTypeEnum = "CPU_FLEET_SHAPE_EXECUTION_DETAILS"
	FleetShapeExecutionDetailsTypeGpuFleetShapeExecutionDetails FleetShapeExecutionDetailsTypeEnum = "GPU_FLEET_SHAPE_EXECUTION_DETAILS"
)

var mappingFleetShapeExecutionDetailsTypeEnum = map[string]FleetShapeExecutionDetailsTypeEnum{
	"CPU_FLEET_SHAPE_EXECUTION_DETAILS": FleetShapeExecutionDetailsTypeCpuFleetShapeExecutionDetails,
	"GPU_FLEET_SHAPE_EXECUTION_DETAILS": FleetShapeExecutionDetailsTypeGpuFleetShapeExecutionDetails,
}

var mappingFleetShapeExecutionDetailsTypeEnumLowerCase = map[string]FleetShapeExecutionDetailsTypeEnum{
	"cpu_fleet_shape_execution_details": FleetShapeExecutionDetailsTypeCpuFleetShapeExecutionDetails,
	"gpu_fleet_shape_execution_details": FleetShapeExecutionDetailsTypeGpuFleetShapeExecutionDetails,
}

// GetFleetShapeExecutionDetailsTypeEnumValues Enumerates the set of values for FleetShapeExecutionDetailsTypeEnum
func GetFleetShapeExecutionDetailsTypeEnumValues() []FleetShapeExecutionDetailsTypeEnum {
	values := make([]FleetShapeExecutionDetailsTypeEnum, 0)
	for _, v := range mappingFleetShapeExecutionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetShapeExecutionDetailsTypeEnumStringValues Enumerates the set of values in String for FleetShapeExecutionDetailsTypeEnum
func GetFleetShapeExecutionDetailsTypeEnumStringValues() []string {
	return []string{
		"CPU_FLEET_SHAPE_EXECUTION_DETAILS",
		"GPU_FLEET_SHAPE_EXECUTION_DETAILS",
	}
}

// GetMappingFleetShapeExecutionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetShapeExecutionDetailsTypeEnum(val string) (FleetShapeExecutionDetailsTypeEnum, bool) {
	enum, ok := mappingFleetShapeExecutionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
