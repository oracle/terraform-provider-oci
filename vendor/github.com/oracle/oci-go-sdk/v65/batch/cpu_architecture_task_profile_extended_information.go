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

// CpuArchitectureTaskProfileExtendedInformation Extended information about CPU architecture for the task profile.
type CpuArchitectureTaskProfileExtendedInformation struct {

	// Type of CPU architecture.
	Architecture CpuArchitectureTaskProfileExtendedInformationArchitectureEnum `mandatory:"true" json:"architecture"`
}

func (m CpuArchitectureTaskProfileExtendedInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CpuArchitectureTaskProfileExtendedInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetCpuArchitectureTaskProfileExtendedInformationArchitectureEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CpuArchitectureTaskProfileExtendedInformation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCpuArchitectureTaskProfileExtendedInformation CpuArchitectureTaskProfileExtendedInformation
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCpuArchitectureTaskProfileExtendedInformation
	}{
		"CPU_ARCHITECTURE_TASK_PROFILE_EXTENDED_INFORMATION",
		(MarshalTypeCpuArchitectureTaskProfileExtendedInformation)(m),
	}

	return json.Marshal(&s)
}

// CpuArchitectureTaskProfileExtendedInformationArchitectureEnum Enum with underlying type: string
type CpuArchitectureTaskProfileExtendedInformationArchitectureEnum string

// Set of constants representing the allowable values for CpuArchitectureTaskProfileExtendedInformationArchitectureEnum
const (
	CpuArchitectureTaskProfileExtendedInformationArchitectureX86 CpuArchitectureTaskProfileExtendedInformationArchitectureEnum = "GENERIC_X86"
	CpuArchitectureTaskProfileExtendedInformationArchitectureArm CpuArchitectureTaskProfileExtendedInformationArchitectureEnum = "GENERIC_ARM"
)

var mappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnum = map[string]CpuArchitectureTaskProfileExtendedInformationArchitectureEnum{
	"GENERIC_X86": CpuArchitectureTaskProfileExtendedInformationArchitectureX86,
	"GENERIC_ARM": CpuArchitectureTaskProfileExtendedInformationArchitectureArm,
}

var mappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnumLowerCase = map[string]CpuArchitectureTaskProfileExtendedInformationArchitectureEnum{
	"generic_x86": CpuArchitectureTaskProfileExtendedInformationArchitectureX86,
	"generic_arm": CpuArchitectureTaskProfileExtendedInformationArchitectureArm,
}

// GetCpuArchitectureTaskProfileExtendedInformationArchitectureEnumValues Enumerates the set of values for CpuArchitectureTaskProfileExtendedInformationArchitectureEnum
func GetCpuArchitectureTaskProfileExtendedInformationArchitectureEnumValues() []CpuArchitectureTaskProfileExtendedInformationArchitectureEnum {
	values := make([]CpuArchitectureTaskProfileExtendedInformationArchitectureEnum, 0)
	for _, v := range mappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetCpuArchitectureTaskProfileExtendedInformationArchitectureEnumStringValues Enumerates the set of values in String for CpuArchitectureTaskProfileExtendedInformationArchitectureEnum
func GetCpuArchitectureTaskProfileExtendedInformationArchitectureEnumStringValues() []string {
	return []string{
		"GENERIC_X86",
		"GENERIC_ARM",
	}
}

// GetMappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnum(val string) (CpuArchitectureTaskProfileExtendedInformationArchitectureEnum, bool) {
	enum, ok := mappingCpuArchitectureTaskProfileExtendedInformationArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
