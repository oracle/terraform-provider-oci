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

// CreateCpuArchitectureTaskProfileExtendedInformationDetails Extended information about CPU architecture for the task profile.
type CreateCpuArchitectureTaskProfileExtendedInformationDetails struct {

	// Type of CPU architecture.
	Architecture CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum `mandatory:"true" json:"architecture"`
}

func (m CreateCpuArchitectureTaskProfileExtendedInformationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCpuArchitectureTaskProfileExtendedInformationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCpuArchitectureTaskProfileExtendedInformationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCpuArchitectureTaskProfileExtendedInformationDetails CreateCpuArchitectureTaskProfileExtendedInformationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateCpuArchitectureTaskProfileExtendedInformationDetails
	}{
		"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION",
		(MarshalTypeCreateCpuArchitectureTaskProfileExtendedInformationDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum Enum with underlying type: string
type CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum string

// Set of constants representing the allowable values for CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum
const (
	CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureX86 CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum = "GENERIC_X86"
	CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureArm CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum = "GENERIC_ARM"
)

var mappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum = map[string]CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum{
	"GENERIC_X86": CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureX86,
	"GENERIC_ARM": CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureArm,
}

var mappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumLowerCase = map[string]CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum{
	"generic_x86": CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureX86,
	"generic_arm": CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureArm,
}

// GetCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumValues Enumerates the set of values for CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum
func GetCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumValues() []CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum {
	values := make([]CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum, 0)
	for _, v := range mappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumStringValues Enumerates the set of values in String for CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum
func GetCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumStringValues() []string {
	return []string{
		"GENERIC_X86",
		"GENERIC_ARM",
	}
}

// GetMappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum(val string) (CreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnum, bool) {
	enum, ok := mappingCreateCpuArchitectureTaskProfileExtendedInformationDetailsArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
