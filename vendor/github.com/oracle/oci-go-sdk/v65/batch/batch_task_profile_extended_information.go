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

// BatchTaskProfileExtendedInformation Extended information for the task profile.
type BatchTaskProfileExtendedInformation interface {
}

type batchtaskprofileextendedinformation struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *batchtaskprofileextendedinformation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbatchtaskprofileextendedinformation batchtaskprofileextendedinformation
	s := struct {
		Model Unmarshalerbatchtaskprofileextendedinformation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *batchtaskprofileextendedinformation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := CpuArchitectureTaskProfileExtendedInformation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := CpuShapeTaskProfileExtendedInformation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := GpuShapeTaskProfileExtendedInformation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BatchTaskProfileExtendedInformation: %s.", m.Type)
		return *m, nil
	}
}

func (m batchtaskprofileextendedinformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m batchtaskprofileextendedinformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchTaskProfileExtendedInformationTypeEnum Enum with underlying type: string
type BatchTaskProfileExtendedInformationTypeEnum string

// Set of constants representing the allowable values for BatchTaskProfileExtendedInformationTypeEnum
const (
	BatchTaskProfileExtendedInformationTypeCpuArchitectureTaskProfileExtendedInformation BatchTaskProfileExtendedInformationTypeEnum = "CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION"
	BatchTaskProfileExtendedInformationTypeCpuShapeTaskProfileExtendedInformation        BatchTaskProfileExtendedInformationTypeEnum = "CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION"
	BatchTaskProfileExtendedInformationTypeGpuShapeTaskProfileExtendedInformation        BatchTaskProfileExtendedInformationTypeEnum = "GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION"
)

var mappingBatchTaskProfileExtendedInformationTypeEnum = map[string]BatchTaskProfileExtendedInformationTypeEnum{
	"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION": BatchTaskProfileExtendedInformationTypeCpuArchitectureTaskProfileExtendedInformation,
	"CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":        BatchTaskProfileExtendedInformationTypeCpuShapeTaskProfileExtendedInformation,
	"GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":        BatchTaskProfileExtendedInformationTypeGpuShapeTaskProfileExtendedInformation,
}

var mappingBatchTaskProfileExtendedInformationTypeEnumLowerCase = map[string]BatchTaskProfileExtendedInformationTypeEnum{
	"cpu_architecture_task_profile_extended_information": BatchTaskProfileExtendedInformationTypeCpuArchitectureTaskProfileExtendedInformation,
	"cpu_shape_task_profile_extended_information":        BatchTaskProfileExtendedInformationTypeCpuShapeTaskProfileExtendedInformation,
	"gpu_shape_task_profile_extended_information":        BatchTaskProfileExtendedInformationTypeGpuShapeTaskProfileExtendedInformation,
}

// GetBatchTaskProfileExtendedInformationTypeEnumValues Enumerates the set of values for BatchTaskProfileExtendedInformationTypeEnum
func GetBatchTaskProfileExtendedInformationTypeEnumValues() []BatchTaskProfileExtendedInformationTypeEnum {
	values := make([]BatchTaskProfileExtendedInformationTypeEnum, 0)
	for _, v := range mappingBatchTaskProfileExtendedInformationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskProfileExtendedInformationTypeEnumStringValues Enumerates the set of values in String for BatchTaskProfileExtendedInformationTypeEnum
func GetBatchTaskProfileExtendedInformationTypeEnumStringValues() []string {
	return []string{
		"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION",
		"CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION",
		"GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION",
	}
}

// GetMappingBatchTaskProfileExtendedInformationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskProfileExtendedInformationTypeEnum(val string) (BatchTaskProfileExtendedInformationTypeEnum, bool) {
	enum, ok := mappingBatchTaskProfileExtendedInformationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
