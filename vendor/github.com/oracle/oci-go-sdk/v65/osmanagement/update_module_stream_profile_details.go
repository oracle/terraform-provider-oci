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

// UpdateModuleStreamProfileDetails Information detailing the state of a module stream profile
type UpdateModuleStreamProfileDetails struct {

	// The name of the profile of the parent stream
	ProfileName *string `mandatory:"true" json:"profileName"`

	// The status of the profile.
	// A profile with the "INSTALLED" status indicates that the
	// profile has been installed.
	// A profile with the "AVAILABLE" status indicates that the
	// profile is not installed, but can be.
	Status UpdateModuleStreamProfileDetailsStatusEnum `mandatory:"true" json:"status"`

	// The date and time of the last status change for this object, as
	// described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// Indicates if the module stream profile is the default
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m UpdateModuleStreamProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateModuleStreamProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateModuleStreamProfileDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateModuleStreamProfileDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateModuleStreamProfileDetailsStatusEnum Enum with underlying type: string
type UpdateModuleStreamProfileDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateModuleStreamProfileDetailsStatusEnum
const (
	UpdateModuleStreamProfileDetailsStatusInstalled UpdateModuleStreamProfileDetailsStatusEnum = "INSTALLED"
	UpdateModuleStreamProfileDetailsStatusAvailable UpdateModuleStreamProfileDetailsStatusEnum = "AVAILABLE"
)

var mappingUpdateModuleStreamProfileDetailsStatusEnum = map[string]UpdateModuleStreamProfileDetailsStatusEnum{
	"INSTALLED": UpdateModuleStreamProfileDetailsStatusInstalled,
	"AVAILABLE": UpdateModuleStreamProfileDetailsStatusAvailable,
}

var mappingUpdateModuleStreamProfileDetailsStatusEnumLowerCase = map[string]UpdateModuleStreamProfileDetailsStatusEnum{
	"installed": UpdateModuleStreamProfileDetailsStatusInstalled,
	"available": UpdateModuleStreamProfileDetailsStatusAvailable,
}

// GetUpdateModuleStreamProfileDetailsStatusEnumValues Enumerates the set of values for UpdateModuleStreamProfileDetailsStatusEnum
func GetUpdateModuleStreamProfileDetailsStatusEnumValues() []UpdateModuleStreamProfileDetailsStatusEnum {
	values := make([]UpdateModuleStreamProfileDetailsStatusEnum, 0)
	for _, v := range mappingUpdateModuleStreamProfileDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateModuleStreamProfileDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateModuleStreamProfileDetailsStatusEnum
func GetUpdateModuleStreamProfileDetailsStatusEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"AVAILABLE",
	}
}

// GetMappingUpdateModuleStreamProfileDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateModuleStreamProfileDetailsStatusEnum(val string) (UpdateModuleStreamProfileDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateModuleStreamProfileDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
