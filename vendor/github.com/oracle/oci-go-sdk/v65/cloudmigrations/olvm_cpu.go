// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmCpu CPU attributes in OLVM
type OlvmCpu struct {

	// CPU architecture
	Architecture OlvmCpuArchitectureEnum `mandatory:"false" json:"architecture,omitempty"`

	// List of cores of this CPU
	Core []OlvmCore `mandatory:"false" json:"core"`

	CpuTune *OlvmCpuTune `mandatory:"false" json:"cpuTune"`

	// Level of this CPU
	Level *int `mandatory:"false" json:"level"`

	// CPU mode
	Mode OlvmCpuModeEnum `mandatory:"false" json:"mode,omitempty"`

	// A human-readable name in plain text.
	Name *string `mandatory:"false" json:"name"`

	// Speed of this CPU
	Speed *float32 `mandatory:"false" json:"speed"`

	CpuTopology *OlvmCpuTopology `mandatory:"false" json:"cpuTopology"`

	// CPU type
	Type *string `mandatory:"false" json:"type"`
}

func (m OlvmCpu) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmCpu) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmCpuArchitectureEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetOlvmCpuArchitectureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmCpuModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetOlvmCpuModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmCpuArchitectureEnum Enum with underlying type: string
type OlvmCpuArchitectureEnum string

// Set of constants representing the allowable values for OlvmCpuArchitectureEnum
const (
	OlvmCpuArchitectureAarch64   OlvmCpuArchitectureEnum = "AARCH64"
	OlvmCpuArchitecturePpc64     OlvmCpuArchitectureEnum = "PPC64"
	OlvmCpuArchitectureS390x     OlvmCpuArchitectureEnum = "S390X"
	OlvmCpuArchitectureUndefined OlvmCpuArchitectureEnum = "UNDEFINED"
	OlvmCpuArchitectureX8664     OlvmCpuArchitectureEnum = "X86_64"
)

var mappingOlvmCpuArchitectureEnum = map[string]OlvmCpuArchitectureEnum{
	"AARCH64":   OlvmCpuArchitectureAarch64,
	"PPC64":     OlvmCpuArchitecturePpc64,
	"S390X":     OlvmCpuArchitectureS390x,
	"UNDEFINED": OlvmCpuArchitectureUndefined,
	"X86_64":    OlvmCpuArchitectureX8664,
}

var mappingOlvmCpuArchitectureEnumLowerCase = map[string]OlvmCpuArchitectureEnum{
	"aarch64":   OlvmCpuArchitectureAarch64,
	"ppc64":     OlvmCpuArchitecturePpc64,
	"s390x":     OlvmCpuArchitectureS390x,
	"undefined": OlvmCpuArchitectureUndefined,
	"x86_64":    OlvmCpuArchitectureX8664,
}

// GetOlvmCpuArchitectureEnumValues Enumerates the set of values for OlvmCpuArchitectureEnum
func GetOlvmCpuArchitectureEnumValues() []OlvmCpuArchitectureEnum {
	values := make([]OlvmCpuArchitectureEnum, 0)
	for _, v := range mappingOlvmCpuArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmCpuArchitectureEnumStringValues Enumerates the set of values in String for OlvmCpuArchitectureEnum
func GetOlvmCpuArchitectureEnumStringValues() []string {
	return []string{
		"AARCH64",
		"PPC64",
		"S390X",
		"UNDEFINED",
		"X86_64",
	}
}

// GetMappingOlvmCpuArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmCpuArchitectureEnum(val string) (OlvmCpuArchitectureEnum, bool) {
	enum, ok := mappingOlvmCpuArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmCpuModeEnum Enum with underlying type: string
type OlvmCpuModeEnum string

// Set of constants representing the allowable values for OlvmCpuModeEnum
const (
	OlvmCpuModeCustom          OlvmCpuModeEnum = "CUSTOM"
	OlvmCpuModeHostModel       OlvmCpuModeEnum = "HOST_MODEL"
	OlvmCpuModeHostPassthrough OlvmCpuModeEnum = "HOST_PASSTHROUGH"
)

var mappingOlvmCpuModeEnum = map[string]OlvmCpuModeEnum{
	"CUSTOM":           OlvmCpuModeCustom,
	"HOST_MODEL":       OlvmCpuModeHostModel,
	"HOST_PASSTHROUGH": OlvmCpuModeHostPassthrough,
}

var mappingOlvmCpuModeEnumLowerCase = map[string]OlvmCpuModeEnum{
	"custom":           OlvmCpuModeCustom,
	"host_model":       OlvmCpuModeHostModel,
	"host_passthrough": OlvmCpuModeHostPassthrough,
}

// GetOlvmCpuModeEnumValues Enumerates the set of values for OlvmCpuModeEnum
func GetOlvmCpuModeEnumValues() []OlvmCpuModeEnum {
	values := make([]OlvmCpuModeEnum, 0)
	for _, v := range mappingOlvmCpuModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmCpuModeEnumStringValues Enumerates the set of values in String for OlvmCpuModeEnum
func GetOlvmCpuModeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"HOST_MODEL",
		"HOST_PASSTHROUGH",
	}
}

// GetMappingOlvmCpuModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmCpuModeEnum(val string) (OlvmCpuModeEnum, bool) {
	enum, ok := mappingOlvmCpuModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
