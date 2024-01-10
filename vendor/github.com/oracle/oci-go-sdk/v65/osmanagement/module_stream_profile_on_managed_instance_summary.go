// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModuleStreamProfileOnManagedInstanceSummary Summary information pertaining to a module stream profile on a managed instance
type ModuleStreamProfileOnManagedInstanceSummary struct {

	// The name of the module that contains the stream profile
	ModuleName *string `mandatory:"true" json:"moduleName"`

	// The name of the stream that contains the profile
	StreamName *string `mandatory:"true" json:"streamName"`

	// The name of the profile
	ProfileName *string `mandatory:"true" json:"profileName"`

	// The status of the profile.
	// A profile with the "INSTALLED" status indicates that the profile has been
	// installed.
	// A profile with the "AVAILABLE" status indicates that the profile is
	// not installed, but can be.
	Status ModuleStreamProfileOnManagedInstanceSummaryStatusEnum `mandatory:"true" json:"status"`

	// The date and time of the last status change for this profile, as
	// described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`
}

func (m ModuleStreamProfileOnManagedInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModuleStreamProfileOnManagedInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModuleStreamProfileOnManagedInstanceSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetModuleStreamProfileOnManagedInstanceSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModuleStreamProfileOnManagedInstanceSummaryStatusEnum Enum with underlying type: string
type ModuleStreamProfileOnManagedInstanceSummaryStatusEnum string

// Set of constants representing the allowable values for ModuleStreamProfileOnManagedInstanceSummaryStatusEnum
const (
	ModuleStreamProfileOnManagedInstanceSummaryStatusInstalled ModuleStreamProfileOnManagedInstanceSummaryStatusEnum = "INSTALLED"
	ModuleStreamProfileOnManagedInstanceSummaryStatusAvailable ModuleStreamProfileOnManagedInstanceSummaryStatusEnum = "AVAILABLE"
)

var mappingModuleStreamProfileOnManagedInstanceSummaryStatusEnum = map[string]ModuleStreamProfileOnManagedInstanceSummaryStatusEnum{
	"INSTALLED": ModuleStreamProfileOnManagedInstanceSummaryStatusInstalled,
	"AVAILABLE": ModuleStreamProfileOnManagedInstanceSummaryStatusAvailable,
}

var mappingModuleStreamProfileOnManagedInstanceSummaryStatusEnumLowerCase = map[string]ModuleStreamProfileOnManagedInstanceSummaryStatusEnum{
	"installed": ModuleStreamProfileOnManagedInstanceSummaryStatusInstalled,
	"available": ModuleStreamProfileOnManagedInstanceSummaryStatusAvailable,
}

// GetModuleStreamProfileOnManagedInstanceSummaryStatusEnumValues Enumerates the set of values for ModuleStreamProfileOnManagedInstanceSummaryStatusEnum
func GetModuleStreamProfileOnManagedInstanceSummaryStatusEnumValues() []ModuleStreamProfileOnManagedInstanceSummaryStatusEnum {
	values := make([]ModuleStreamProfileOnManagedInstanceSummaryStatusEnum, 0)
	for _, v := range mappingModuleStreamProfileOnManagedInstanceSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetModuleStreamProfileOnManagedInstanceSummaryStatusEnumStringValues Enumerates the set of values in String for ModuleStreamProfileOnManagedInstanceSummaryStatusEnum
func GetModuleStreamProfileOnManagedInstanceSummaryStatusEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"AVAILABLE",
	}
}

// GetMappingModuleStreamProfileOnManagedInstanceSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModuleStreamProfileOnManagedInstanceSummaryStatusEnum(val string) (ModuleStreamProfileOnManagedInstanceSummaryStatusEnum, bool) {
	enum, ok := mappingModuleStreamProfileOnManagedInstanceSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
