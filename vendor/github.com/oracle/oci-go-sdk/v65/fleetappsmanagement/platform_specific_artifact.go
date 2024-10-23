// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PlatformSpecificArtifact Patch artifact metadata Details which is platform specific.
type PlatformSpecificArtifact struct {
	Content ContentDetails `mandatory:"true" json:"content"`

	// The OS type the patch is applicable for.
	OsType PlatformSpecificArtifactOsTypeEnum `mandatory:"true" json:"osType"`

	// System architecture.
	Architecture PlatformSpecificArtifactArchitectureEnum `mandatory:"true" json:"architecture"`
}

func (m PlatformSpecificArtifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlatformSpecificArtifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlatformSpecificArtifactOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetPlatformSpecificArtifactOsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPlatformSpecificArtifactArchitectureEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetPlatformSpecificArtifactArchitectureEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PlatformSpecificArtifact) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Content      contentdetails                           `json:"content"`
		OsType       PlatformSpecificArtifactOsTypeEnum       `json:"osType"`
		Architecture PlatformSpecificArtifactArchitectureEnum `json:"architecture"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(ContentDetails)
	} else {
		m.Content = nil
	}

	m.OsType = model.OsType

	m.Architecture = model.Architecture

	return
}

// PlatformSpecificArtifactOsTypeEnum Enum with underlying type: string
type PlatformSpecificArtifactOsTypeEnum string

// Set of constants representing the allowable values for PlatformSpecificArtifactOsTypeEnum
const (
	PlatformSpecificArtifactOsTypeWindows PlatformSpecificArtifactOsTypeEnum = "WINDOWS"
	PlatformSpecificArtifactOsTypeLinux   PlatformSpecificArtifactOsTypeEnum = "LINUX"
)

var mappingPlatformSpecificArtifactOsTypeEnum = map[string]PlatformSpecificArtifactOsTypeEnum{
	"WINDOWS": PlatformSpecificArtifactOsTypeWindows,
	"LINUX":   PlatformSpecificArtifactOsTypeLinux,
}

var mappingPlatformSpecificArtifactOsTypeEnumLowerCase = map[string]PlatformSpecificArtifactOsTypeEnum{
	"windows": PlatformSpecificArtifactOsTypeWindows,
	"linux":   PlatformSpecificArtifactOsTypeLinux,
}

// GetPlatformSpecificArtifactOsTypeEnumValues Enumerates the set of values for PlatformSpecificArtifactOsTypeEnum
func GetPlatformSpecificArtifactOsTypeEnumValues() []PlatformSpecificArtifactOsTypeEnum {
	values := make([]PlatformSpecificArtifactOsTypeEnum, 0)
	for _, v := range mappingPlatformSpecificArtifactOsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformSpecificArtifactOsTypeEnumStringValues Enumerates the set of values in String for PlatformSpecificArtifactOsTypeEnum
func GetPlatformSpecificArtifactOsTypeEnumStringValues() []string {
	return []string{
		"WINDOWS",
		"LINUX",
	}
}

// GetMappingPlatformSpecificArtifactOsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformSpecificArtifactOsTypeEnum(val string) (PlatformSpecificArtifactOsTypeEnum, bool) {
	enum, ok := mappingPlatformSpecificArtifactOsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PlatformSpecificArtifactArchitectureEnum Enum with underlying type: string
type PlatformSpecificArtifactArchitectureEnum string

// Set of constants representing the allowable values for PlatformSpecificArtifactArchitectureEnum
const (
	PlatformSpecificArtifactArchitectureArm64 PlatformSpecificArtifactArchitectureEnum = "ARM_64"
	PlatformSpecificArtifactArchitectureX64   PlatformSpecificArtifactArchitectureEnum = "X64"
)

var mappingPlatformSpecificArtifactArchitectureEnum = map[string]PlatformSpecificArtifactArchitectureEnum{
	"ARM_64": PlatformSpecificArtifactArchitectureArm64,
	"X64":    PlatformSpecificArtifactArchitectureX64,
}

var mappingPlatformSpecificArtifactArchitectureEnumLowerCase = map[string]PlatformSpecificArtifactArchitectureEnum{
	"arm_64": PlatformSpecificArtifactArchitectureArm64,
	"x64":    PlatformSpecificArtifactArchitectureX64,
}

// GetPlatformSpecificArtifactArchitectureEnumValues Enumerates the set of values for PlatformSpecificArtifactArchitectureEnum
func GetPlatformSpecificArtifactArchitectureEnumValues() []PlatformSpecificArtifactArchitectureEnum {
	values := make([]PlatformSpecificArtifactArchitectureEnum, 0)
	for _, v := range mappingPlatformSpecificArtifactArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformSpecificArtifactArchitectureEnumStringValues Enumerates the set of values in String for PlatformSpecificArtifactArchitectureEnum
func GetPlatformSpecificArtifactArchitectureEnumStringValues() []string {
	return []string{
		"ARM_64",
		"X64",
	}
}

// GetMappingPlatformSpecificArtifactArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformSpecificArtifactArchitectureEnum(val string) (PlatformSpecificArtifactArchitectureEnum, bool) {
	enum, ok := mappingPlatformSpecificArtifactArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
