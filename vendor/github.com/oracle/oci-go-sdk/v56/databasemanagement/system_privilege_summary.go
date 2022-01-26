// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SystemPrivilegeSummary Summary of SystemPrivileges.
type SystemPrivilegeSummary struct {

	// The name of a system privilege
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the grant was with the ADMIN option (YES) or not (NO)
	AdminOption SystemPrivilegeSummaryAdminOptionEnum `mandatory:"false" json:"adminOption,omitempty"`

	// Indicates how the grant was made. Possible values:
	// YES if the role was granted commonly (CONTAINER=ALL was used)
	// NO if the role was granted locally (CONTAINER=ALL was not used)
	Common SystemPrivilegeSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the role grant was inherited from another container (YES) or not (NO)
	Inherited SystemPrivilegeSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m SystemPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// SystemPrivilegeSummaryAdminOptionEnum Enum with underlying type: string
type SystemPrivilegeSummaryAdminOptionEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryAdminOptionEnum
const (
	SystemPrivilegeSummaryAdminOptionYes SystemPrivilegeSummaryAdminOptionEnum = "YES"
	SystemPrivilegeSummaryAdminOptionNo  SystemPrivilegeSummaryAdminOptionEnum = "NO"
)

var mappingSystemPrivilegeSummaryAdminOption = map[string]SystemPrivilegeSummaryAdminOptionEnum{
	"YES": SystemPrivilegeSummaryAdminOptionYes,
	"NO":  SystemPrivilegeSummaryAdminOptionNo,
}

// GetSystemPrivilegeSummaryAdminOptionEnumValues Enumerates the set of values for SystemPrivilegeSummaryAdminOptionEnum
func GetSystemPrivilegeSummaryAdminOptionEnumValues() []SystemPrivilegeSummaryAdminOptionEnum {
	values := make([]SystemPrivilegeSummaryAdminOptionEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryAdminOption {
		values = append(values, v)
	}
	return values
}

// SystemPrivilegeSummaryCommonEnum Enum with underlying type: string
type SystemPrivilegeSummaryCommonEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryCommonEnum
const (
	SystemPrivilegeSummaryCommonYes SystemPrivilegeSummaryCommonEnum = "YES"
	SystemPrivilegeSummaryCommonNo  SystemPrivilegeSummaryCommonEnum = "NO"
)

var mappingSystemPrivilegeSummaryCommon = map[string]SystemPrivilegeSummaryCommonEnum{
	"YES": SystemPrivilegeSummaryCommonYes,
	"NO":  SystemPrivilegeSummaryCommonNo,
}

// GetSystemPrivilegeSummaryCommonEnumValues Enumerates the set of values for SystemPrivilegeSummaryCommonEnum
func GetSystemPrivilegeSummaryCommonEnumValues() []SystemPrivilegeSummaryCommonEnum {
	values := make([]SystemPrivilegeSummaryCommonEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryCommon {
		values = append(values, v)
	}
	return values
}

// SystemPrivilegeSummaryInheritedEnum Enum with underlying type: string
type SystemPrivilegeSummaryInheritedEnum string

// Set of constants representing the allowable values for SystemPrivilegeSummaryInheritedEnum
const (
	SystemPrivilegeSummaryInheritedYes SystemPrivilegeSummaryInheritedEnum = "YES"
	SystemPrivilegeSummaryInheritedNo  SystemPrivilegeSummaryInheritedEnum = "NO"
)

var mappingSystemPrivilegeSummaryInherited = map[string]SystemPrivilegeSummaryInheritedEnum{
	"YES": SystemPrivilegeSummaryInheritedYes,
	"NO":  SystemPrivilegeSummaryInheritedNo,
}

// GetSystemPrivilegeSummaryInheritedEnumValues Enumerates the set of values for SystemPrivilegeSummaryInheritedEnum
func GetSystemPrivilegeSummaryInheritedEnumValues() []SystemPrivilegeSummaryInheritedEnum {
	values := make([]SystemPrivilegeSummaryInheritedEnum, 0)
	for _, v := range mappingSystemPrivilegeSummaryInherited {
		values = append(values, v)
	}
	return values
}
