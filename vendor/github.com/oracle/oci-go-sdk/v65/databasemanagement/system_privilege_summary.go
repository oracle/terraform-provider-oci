// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SystemPrivilegeSummary A Summary of system privileges.
type SystemPrivilegeSummary struct {

	// The name of a system privilege.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the system privilege is granted with the ADMIN option (YES) or not (NO).
	AdminOption SystemPrivilegeSummaryAdminOptionEnum `mandatory:"false" json:"adminOption,omitempty"`

	// Indicates how the system privilege was granted. Possible values:
	// YES if the system privilege is granted commonly (CONTAINER=ALL is used)
	// NO if the system privilege is granted locally (CONTAINER=ALL is not used)
	Common SystemPrivilegeSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the granted system privilege is inherited from another container (YES) or not (NO).
	Inherited SystemPrivilegeSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m SystemPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SystemPrivilegeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSystemPrivilegeSummaryAdminOptionEnum(string(m.AdminOption)); !ok && m.AdminOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdminOption: %s. Supported values are: %s.", m.AdminOption, strings.Join(GetSystemPrivilegeSummaryAdminOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSystemPrivilegeSummaryCommonEnum(string(m.Common)); !ok && m.Common != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Common: %s. Supported values are: %s.", m.Common, strings.Join(GetSystemPrivilegeSummaryCommonEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSystemPrivilegeSummaryInheritedEnum(string(m.Inherited)); !ok && m.Inherited != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inherited: %s. Supported values are: %s.", m.Inherited, strings.Join(GetSystemPrivilegeSummaryInheritedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SystemPrivilegeSummaryAdminOptionEnum Enum with underlying type: string
type SystemPrivilegeSummaryAdminOptionEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryAdminOptionEnum
const (
	SystemPrivilegeSummaryAdminOptionYes SystemPrivilegeSummaryAdminOptionEnum = "YES"
	SystemPrivilegeSummaryAdminOptionNo  SystemPrivilegeSummaryAdminOptionEnum = "NO"
)

var mappingSystemPrivilegeSummaryAdminOptionEnum = map[string]SystemPrivilegeSummaryAdminOptionEnum{
	"YES": SystemPrivilegeSummaryAdminOptionYes,
	"NO":  SystemPrivilegeSummaryAdminOptionNo,
}

var mappingSystemPrivilegeSummaryAdminOptionEnumLowerCase = map[string]SystemPrivilegeSummaryAdminOptionEnum{
	"yes": SystemPrivilegeSummaryAdminOptionYes,
	"no":  SystemPrivilegeSummaryAdminOptionNo,
}

// GetSystemPrivilegeSummaryAdminOptionEnumValues Enumerates the set of values for SystemPrivilegeSummaryAdminOptionEnum
func GetSystemPrivilegeSummaryAdminOptionEnumValues() []SystemPrivilegeSummaryAdminOptionEnum {
	values := make([]SystemPrivilegeSummaryAdminOptionEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryAdminOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSystemPrivilegeSummaryAdminOptionEnumStringValues Enumerates the set of values in String for SystemPrivilegeSummaryAdminOptionEnum
func GetSystemPrivilegeSummaryAdminOptionEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSystemPrivilegeSummaryAdminOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSystemPrivilegeSummaryAdminOptionEnum(val string) (SystemPrivilegeSummaryAdminOptionEnum, bool) {
	enum, ok := mappingSystemPrivilegeSummaryAdminOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SystemPrivilegeSummaryCommonEnum Enum with underlying type: string
type SystemPrivilegeSummaryCommonEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryCommonEnum
const (
	SystemPrivilegeSummaryCommonYes SystemPrivilegeSummaryCommonEnum = "YES"
	SystemPrivilegeSummaryCommonNo  SystemPrivilegeSummaryCommonEnum = "NO"
)

var mappingSystemPrivilegeSummaryCommonEnum = map[string]SystemPrivilegeSummaryCommonEnum{
	"YES": SystemPrivilegeSummaryCommonYes,
	"NO":  SystemPrivilegeSummaryCommonNo,
}

var mappingSystemPrivilegeSummaryCommonEnumLowerCase = map[string]SystemPrivilegeSummaryCommonEnum{
	"yes": SystemPrivilegeSummaryCommonYes,
	"no":  SystemPrivilegeSummaryCommonNo,
}

// GetSystemPrivilegeSummaryCommonEnumValues Enumerates the set of values for SystemPrivilegeSummaryCommonEnum
func GetSystemPrivilegeSummaryCommonEnumValues() []SystemPrivilegeSummaryCommonEnum {
	values := make([]SystemPrivilegeSummaryCommonEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryCommonEnum {
		values = append(values, v)
	}
	return values
}

// GetSystemPrivilegeSummaryCommonEnumStringValues Enumerates the set of values in String for SystemPrivilegeSummaryCommonEnum
func GetSystemPrivilegeSummaryCommonEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSystemPrivilegeSummaryCommonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSystemPrivilegeSummaryCommonEnum(val string) (SystemPrivilegeSummaryCommonEnum, bool) {
	enum, ok := mappingSystemPrivilegeSummaryCommonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SystemPrivilegeSummaryInheritedEnum Enum with underlying type: string
type SystemPrivilegeSummaryInheritedEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryInheritedEnum
const (
	SystemPrivilegeSummaryInheritedYes SystemPrivilegeSummaryInheritedEnum = "YES"
	SystemPrivilegeSummaryInheritedNo  SystemPrivilegeSummaryInheritedEnum = "NO"
)

var mappingSystemPrivilegeSummaryInheritedEnum = map[string]SystemPrivilegeSummaryInheritedEnum{
	"YES": SystemPrivilegeSummaryInheritedYes,
	"NO":  SystemPrivilegeSummaryInheritedNo,
}

var mappingSystemPrivilegeSummaryInheritedEnumLowerCase = map[string]SystemPrivilegeSummaryInheritedEnum{
	"yes": SystemPrivilegeSummaryInheritedYes,
	"no":  SystemPrivilegeSummaryInheritedNo,
}

// GetSystemPrivilegeSummaryInheritedEnumValues Enumerates the set of values for SystemPrivilegeSummaryInheritedEnum
func GetSystemPrivilegeSummaryInheritedEnumValues() []SystemPrivilegeSummaryInheritedEnum {
	values := make([]SystemPrivilegeSummaryInheritedEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryInheritedEnum {
		values = append(values, v)
	}
	return values
}

// GetSystemPrivilegeSummaryInheritedEnumStringValues Enumerates the set of values in String for SystemPrivilegeSummaryInheritedEnum
func GetSystemPrivilegeSummaryInheritedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSystemPrivilegeSummaryInheritedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSystemPrivilegeSummaryInheritedEnum(val string) (SystemPrivilegeSummaryInheritedEnum, bool) {
	enum, ok := mappingSystemPrivilegeSummaryInheritedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
