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

// RoleSummary Summary of each role.
type RoleSummary struct {

	// The name of a granted role
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the grant was with the ADMIN OPTION (YES) or not (NO)
	AdminOption RoleSummaryAdminOptionEnum `mandatory:"false" json:"adminOption,omitempty"`

	// Indicates whether the grant was with the DELEGATE OPTION (YES) or not (NO)
	DelegateOption RoleSummaryDelegateOptionEnum `mandatory:"false" json:"delegateOption,omitempty"`

	// Indicates whether the role is designated as a DEFAULT ROLE for the user (YES) or not (NO)
	DefaultRole RoleSummaryDefaultRoleEnum `mandatory:"false" json:"defaultRole,omitempty"`

	// Indicates how the grant was made. Possible values:
	// YES if the role was granted commonly (CONTAINER=ALL was used)
	// NO if the role was granted locally (CONTAINER=ALL was not used)
	Common RoleSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the role grant was inherited from another container (YES) or not (NO)
	Inherited RoleSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m RoleSummary) String() string {
	return common.PointerString(m)
}

// RoleSummaryAdminOptionEnum Enum with underlying type: string
type RoleSummaryAdminOptionEnum string

// Set of constants representing the allowable values for RoleSummaryAdminOptionEnum
const (
	RoleSummaryAdminOptionYes RoleSummaryAdminOptionEnum = "YES"
	RoleSummaryAdminOptionNo  RoleSummaryAdminOptionEnum = "NO"
)

var mappingRoleSummaryAdminOption = map[string]RoleSummaryAdminOptionEnum{
	"YES": RoleSummaryAdminOptionYes,
	"NO":  RoleSummaryAdminOptionNo,
}

// GetRoleSummaryAdminOptionEnumValues Enumerates the set of values for RoleSummaryAdminOptionEnum
func GetRoleSummaryAdminOptionEnumValues() []RoleSummaryAdminOptionEnum {
	values := make([]RoleSummaryAdminOptionEnum, 0)
	for _, v := range mappingRoleSummaryAdminOption {
		values = append(values, v)
	}
	return values
}

// RoleSummaryDelegateOptionEnum Enum with underlying type: string
type RoleSummaryDelegateOptionEnum string

// Set of constants representing the allowable values for RoleSummaryDelegateOptionEnum
const (
	RoleSummaryDelegateOptionYes RoleSummaryDelegateOptionEnum = "YES"
	RoleSummaryDelegateOptionNo  RoleSummaryDelegateOptionEnum = "NO"
)

var mappingRoleSummaryDelegateOption = map[string]RoleSummaryDelegateOptionEnum{
	"YES": RoleSummaryDelegateOptionYes,
	"NO":  RoleSummaryDelegateOptionNo,
}

// GetRoleSummaryDelegateOptionEnumValues Enumerates the set of values for RoleSummaryDelegateOptionEnum
func GetRoleSummaryDelegateOptionEnumValues() []RoleSummaryDelegateOptionEnum {
	values := make([]RoleSummaryDelegateOptionEnum, 0)
	for _, v := range mappingRoleSummaryDelegateOption {
		values = append(values, v)
	}
	return values
}

// RoleSummaryDefaultRoleEnum Enum with underlying type: string
type RoleSummaryDefaultRoleEnum string

// Set of constants representing the allowable values for RoleSummaryDefaultRoleEnum
const (
	RoleSummaryDefaultRoleYes RoleSummaryDefaultRoleEnum = "YES"
	RoleSummaryDefaultRoleNo  RoleSummaryDefaultRoleEnum = "NO"
)

var mappingRoleSummaryDefaultRole = map[string]RoleSummaryDefaultRoleEnum{
	"YES": RoleSummaryDefaultRoleYes,
	"NO":  RoleSummaryDefaultRoleNo,
}

// GetRoleSummaryDefaultRoleEnumValues Enumerates the set of values for RoleSummaryDefaultRoleEnum
func GetRoleSummaryDefaultRoleEnumValues() []RoleSummaryDefaultRoleEnum {
	values := make([]RoleSummaryDefaultRoleEnum, 0)
	for _, v := range mappingRoleSummaryDefaultRole {
		values = append(values, v)
	}
	return values
}

// RoleSummaryCommonEnum Enum with underlying type: string
type RoleSummaryCommonEnum string

// Set of constants representing the allowable values for RoleSummaryCommonEnum
const (
	RoleSummaryCommonYes RoleSummaryCommonEnum = "YES"
	RoleSummaryCommonNo  RoleSummaryCommonEnum = "NO"
)

var mappingRoleSummaryCommon = map[string]RoleSummaryCommonEnum{
	"YES": RoleSummaryCommonYes,
	"NO":  RoleSummaryCommonNo,
}

// GetRoleSummaryCommonEnumValues Enumerates the set of values for RoleSummaryCommonEnum
func GetRoleSummaryCommonEnumValues() []RoleSummaryCommonEnum {
	values := make([]RoleSummaryCommonEnum, 0)
	for _, v := range mappingRoleSummaryCommon {
		values = append(values, v)
	}
	return values
}

// RoleSummaryInheritedEnum Enum with underlying type: string
type RoleSummaryInheritedEnum string

// Set of constants representing the allowable values for RoleSummaryInheritedEnum
const (
	RoleSummaryInheritedYes RoleSummaryInheritedEnum = "YES"
	RoleSummaryInheritedNo  RoleSummaryInheritedEnum = "NO"
)

var mappingRoleSummaryInherited = map[string]RoleSummaryInheritedEnum{
	"YES": RoleSummaryInheritedYes,
	"NO":  RoleSummaryInheritedNo,
}

// GetRoleSummaryInheritedEnumValues Enumerates the set of values for RoleSummaryInheritedEnum
func GetRoleSummaryInheritedEnumValues() []RoleSummaryInheritedEnum {
	values := make([]RoleSummaryInheritedEnum, 0)
	for _, v := range mappingRoleSummaryInherited {
		values = append(values, v)
	}
	return values
}
