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

// FunctionsRuntimeVersionSummary Summary of the FunctionsRuntimeVersion.
type FunctionsRuntimeVersionSummary struct {

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
	SupportedArchitectures []FunctionsRuntimeVersionSummarySupportedArchitecturesEnum `mandatory:"true" json:"supportedArchitectures"`

	// The current state of the FunctionsRuntimeVersion resource.
	LifecycleState FunctionsRuntimeVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Details of the change in the FunctionsRuntimeVersion of the FunctionsRuntime.
	Metadata *string `mandatory:"true" json:"metadata"`

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

func (m FunctionsRuntimeVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionsRuntimeVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.SupportedArchitectures {
		if _, ok := GetMappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedArchitectures: %s. Supported values are: %s.", val, strings.Join(GetFunctionsRuntimeVersionSummarySupportedArchitecturesEnumStringValues(), ",")))
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

// FunctionsRuntimeVersionSummarySupportedArchitecturesEnum Enum with underlying type: string
type FunctionsRuntimeVersionSummarySupportedArchitecturesEnum string

// Set of constants representing the allowable values for FunctionsRuntimeVersionSummarySupportedArchitecturesEnum
const (
	FunctionsRuntimeVersionSummarySupportedArchitecturesArm FunctionsRuntimeVersionSummarySupportedArchitecturesEnum = "ARM"
	FunctionsRuntimeVersionSummarySupportedArchitecturesX86 FunctionsRuntimeVersionSummarySupportedArchitecturesEnum = "X86"
)

var mappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnum = map[string]FunctionsRuntimeVersionSummarySupportedArchitecturesEnum{
	"ARM": FunctionsRuntimeVersionSummarySupportedArchitecturesArm,
	"X86": FunctionsRuntimeVersionSummarySupportedArchitecturesX86,
}

var mappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnumLowerCase = map[string]FunctionsRuntimeVersionSummarySupportedArchitecturesEnum{
	"arm": FunctionsRuntimeVersionSummarySupportedArchitecturesArm,
	"x86": FunctionsRuntimeVersionSummarySupportedArchitecturesX86,
}

// GetFunctionsRuntimeVersionSummarySupportedArchitecturesEnumValues Enumerates the set of values for FunctionsRuntimeVersionSummarySupportedArchitecturesEnum
func GetFunctionsRuntimeVersionSummarySupportedArchitecturesEnumValues() []FunctionsRuntimeVersionSummarySupportedArchitecturesEnum {
	values := make([]FunctionsRuntimeVersionSummarySupportedArchitecturesEnum, 0)
	for _, v := range mappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionsRuntimeVersionSummarySupportedArchitecturesEnumStringValues Enumerates the set of values in String for FunctionsRuntimeVersionSummarySupportedArchitecturesEnum
func GetFunctionsRuntimeVersionSummarySupportedArchitecturesEnumStringValues() []string {
	return []string{
		"ARM",
		"X86",
	}
}

// GetMappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnum(val string) (FunctionsRuntimeVersionSummarySupportedArchitecturesEnum, bool) {
	enum, ok := mappingFunctionsRuntimeVersionSummarySupportedArchitecturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
