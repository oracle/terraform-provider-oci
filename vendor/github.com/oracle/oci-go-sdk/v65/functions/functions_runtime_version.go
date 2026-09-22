// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionsRuntimeVersion This represents a version of a FunctionsRuntime. Each new functional update to a FunctionsRuntime from the Functions team will change
// the image that will result in the creation of new FunctionsRuntimeVersion resource creation. This is a sub-resource of a FunctionsRuntime.
type FunctionsRuntimeVersion struct {

	// The OCID of the FunctionsRuntimeVersion that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the FunctionsRuntime this resource version belongs to.
	FunctionsRuntimeId *string `mandatory:"true" json:"functionsRuntimeId"`

	// The display name of the FunctionsRuntimeVersion.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The version of the operating system of the FunctionsRuntime. This is the OS version that the FunctionsRuntime provides for execution of customer payloads.
	OsVersion *string `mandatory:"true" json:"osVersion"`

	// The version of the programming language of the FunctionsRuntime. This is the language version that the FunctionsRuntime provides for execution of customer payloads.
	LanguageVersion *string `mandatory:"true" json:"languageVersion"`

	// The list of supported architectures for the FunctionsRuntimeVersion.
	SupportedArchitectures []FunctionsRuntimeVersionSupportedArchitecturesEnum `mandatory:"true" json:"supportedArchitectures"`

	// Details of the change in the FunctionsRuntimeVersion of the FunctionsRuntime.
	Metadata *string `mandatory:"true" json:"metadata"`

	// The current state of the FunctionsRuntimeVersion resource.
	// - `ACTIVE`: The resource is currently active and operational.
	// - `INACTIVE`: The resource is currently inactive and not operational.
	LifecycleState FunctionsRuntimeVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the FunctionsRuntimeVersion was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the FunctionsRuntimeVersion was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FunctionsRuntimeVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionsRuntimeVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.SupportedArchitectures {
		if _, ok := GetMappingFunctionsRuntimeVersionSupportedArchitecturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedArchitectures: %s. Supported values are: %s.", val, strings.Join(GetFunctionsRuntimeVersionSupportedArchitecturesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingFunctionsRuntimeVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFunctionsRuntimeVersionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionsRuntimeVersionSupportedArchitecturesEnum Enum with underlying type: string
type FunctionsRuntimeVersionSupportedArchitecturesEnum string

// Set of constants representing the allowable values for FunctionsRuntimeVersionSupportedArchitecturesEnum
const (
	FunctionsRuntimeVersionSupportedArchitecturesArm FunctionsRuntimeVersionSupportedArchitecturesEnum = "ARM"
	FunctionsRuntimeVersionSupportedArchitecturesX86 FunctionsRuntimeVersionSupportedArchitecturesEnum = "X86"
)

var mappingFunctionsRuntimeVersionSupportedArchitecturesEnum = map[string]FunctionsRuntimeVersionSupportedArchitecturesEnum{
	"ARM": FunctionsRuntimeVersionSupportedArchitecturesArm,
	"X86": FunctionsRuntimeVersionSupportedArchitecturesX86,
}

var mappingFunctionsRuntimeVersionSupportedArchitecturesEnumLowerCase = map[string]FunctionsRuntimeVersionSupportedArchitecturesEnum{
	"arm": FunctionsRuntimeVersionSupportedArchitecturesArm,
	"x86": FunctionsRuntimeVersionSupportedArchitecturesX86,
}

// GetFunctionsRuntimeVersionSupportedArchitecturesEnumValues Enumerates the set of values for FunctionsRuntimeVersionSupportedArchitecturesEnum
func GetFunctionsRuntimeVersionSupportedArchitecturesEnumValues() []FunctionsRuntimeVersionSupportedArchitecturesEnum {
	values := make([]FunctionsRuntimeVersionSupportedArchitecturesEnum, 0)
	for _, v := range mappingFunctionsRuntimeVersionSupportedArchitecturesEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionsRuntimeVersionSupportedArchitecturesEnumStringValues Enumerates the set of values in String for FunctionsRuntimeVersionSupportedArchitecturesEnum
func GetFunctionsRuntimeVersionSupportedArchitecturesEnumStringValues() []string {
	return []string{
		"ARM",
		"X86",
	}
}

// GetMappingFunctionsRuntimeVersionSupportedArchitecturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionsRuntimeVersionSupportedArchitecturesEnum(val string) (FunctionsRuntimeVersionSupportedArchitecturesEnum, bool) {
	enum, ok := mappingFunctionsRuntimeVersionSupportedArchitecturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FunctionsRuntimeVersionLifecycleStateEnum Enum with underlying type: string
type FunctionsRuntimeVersionLifecycleStateEnum string

// Set of constants representing the allowable values for FunctionsRuntimeVersionLifecycleStateEnum
const (
	FunctionsRuntimeVersionLifecycleStateActive   FunctionsRuntimeVersionLifecycleStateEnum = "ACTIVE"
	FunctionsRuntimeVersionLifecycleStateInactive FunctionsRuntimeVersionLifecycleStateEnum = "INACTIVE"
)

var mappingFunctionsRuntimeVersionLifecycleStateEnum = map[string]FunctionsRuntimeVersionLifecycleStateEnum{
	"ACTIVE":   FunctionsRuntimeVersionLifecycleStateActive,
	"INACTIVE": FunctionsRuntimeVersionLifecycleStateInactive,
}

var mappingFunctionsRuntimeVersionLifecycleStateEnumLowerCase = map[string]FunctionsRuntimeVersionLifecycleStateEnum{
	"active":   FunctionsRuntimeVersionLifecycleStateActive,
	"inactive": FunctionsRuntimeVersionLifecycleStateInactive,
}

// GetFunctionsRuntimeVersionLifecycleStateEnumValues Enumerates the set of values for FunctionsRuntimeVersionLifecycleStateEnum
func GetFunctionsRuntimeVersionLifecycleStateEnumValues() []FunctionsRuntimeVersionLifecycleStateEnum {
	values := make([]FunctionsRuntimeVersionLifecycleStateEnum, 0)
	for _, v := range mappingFunctionsRuntimeVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionsRuntimeVersionLifecycleStateEnumStringValues Enumerates the set of values in String for FunctionsRuntimeVersionLifecycleStateEnum
func GetFunctionsRuntimeVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingFunctionsRuntimeVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionsRuntimeVersionLifecycleStateEnum(val string) (FunctionsRuntimeVersionLifecycleStateEnum, bool) {
	enum, ok := mappingFunctionsRuntimeVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
