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

// CreateBatchTaskProfileExtendedInformationDetails Extended information for the task profile.
type CreateBatchTaskProfileExtendedInformationDetails interface {
}

type createbatchtaskprofileextendedinformationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createbatchtaskprofileextendedinformationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebatchtaskprofileextendedinformationdetails createbatchtaskprofileextendedinformationdetails
	s := struct {
		Model Unmarshalercreatebatchtaskprofileextendedinformationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbatchtaskprofileextendedinformationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := CreateGpuShapeTaskProfileExtendedInformationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := CreateCpuShapeTaskProfileExtendedInformationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION":
		mm := CreateCpuArchitectureTaskProfileExtendedInformationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateBatchTaskProfileExtendedInformationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createbatchtaskprofileextendedinformationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createbatchtaskprofileextendedinformationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBatchTaskProfileExtendedInformationDetailsTypeEnum Enum with underlying type: string
type CreateBatchTaskProfileExtendedInformationDetailsTypeEnum string

// Set of constants representing the allowable values for CreateBatchTaskProfileExtendedInformationDetailsTypeEnum
const (
	CreateBatchTaskProfileExtendedInformationDetailsTypeCpuArchitectureTaskProfileExtendedInformation CreateBatchTaskProfileExtendedInformationDetailsTypeEnum = "CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION"
	CreateBatchTaskProfileExtendedInformationDetailsTypeCpuShapeTaskProfileExtendedInformation        CreateBatchTaskProfileExtendedInformationDetailsTypeEnum = "CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION"
	CreateBatchTaskProfileExtendedInformationDetailsTypeGpuShapeTaskProfileExtendedInformation        CreateBatchTaskProfileExtendedInformationDetailsTypeEnum = "GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION"
)

var mappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnum = map[string]CreateBatchTaskProfileExtendedInformationDetailsTypeEnum{
	"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION": CreateBatchTaskProfileExtendedInformationDetailsTypeCpuArchitectureTaskProfileExtendedInformation,
	"CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":        CreateBatchTaskProfileExtendedInformationDetailsTypeCpuShapeTaskProfileExtendedInformation,
	"GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION":        CreateBatchTaskProfileExtendedInformationDetailsTypeGpuShapeTaskProfileExtendedInformation,
}

var mappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnumLowerCase = map[string]CreateBatchTaskProfileExtendedInformationDetailsTypeEnum{
	"cpu_architecture_task_profile_extended_information": CreateBatchTaskProfileExtendedInformationDetailsTypeCpuArchitectureTaskProfileExtendedInformation,
	"cpu_shape_task_profile_extended_information":        CreateBatchTaskProfileExtendedInformationDetailsTypeCpuShapeTaskProfileExtendedInformation,
	"gpu_shape_task_profile_extended_information":        CreateBatchTaskProfileExtendedInformationDetailsTypeGpuShapeTaskProfileExtendedInformation,
}

// GetCreateBatchTaskProfileExtendedInformationDetailsTypeEnumValues Enumerates the set of values for CreateBatchTaskProfileExtendedInformationDetailsTypeEnum
func GetCreateBatchTaskProfileExtendedInformationDetailsTypeEnumValues() []CreateBatchTaskProfileExtendedInformationDetailsTypeEnum {
	values := make([]CreateBatchTaskProfileExtendedInformationDetailsTypeEnum, 0)
	for _, v := range mappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBatchTaskProfileExtendedInformationDetailsTypeEnumStringValues Enumerates the set of values in String for CreateBatchTaskProfileExtendedInformationDetailsTypeEnum
func GetCreateBatchTaskProfileExtendedInformationDetailsTypeEnumStringValues() []string {
	return []string{
		"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION",
		"CPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION",
		"GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION",
	}
}

// GetMappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnum(val string) (CreateBatchTaskProfileExtendedInformationDetailsTypeEnum, bool) {
	enum, ok := mappingCreateBatchTaskProfileExtendedInformationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
