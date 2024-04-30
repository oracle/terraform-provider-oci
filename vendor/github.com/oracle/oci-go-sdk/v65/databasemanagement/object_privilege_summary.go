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

// ObjectPrivilegeSummary A summary of object privileges.
type ObjectPrivilegeSummary struct {

	// The name of the privilege on the object.
	Name *string `mandatory:"false" json:"name"`

	// The type of object.
	SchemaType *string `mandatory:"false" json:"schemaType"`

	// The owner of the object.
	Owner *string `mandatory:"false" json:"owner"`

	// The name of the user who granted the object privilege.
	Grantor *string `mandatory:"false" json:"grantor"`

	// Indicates whether the privilege is granted with the HIERARCHY OPTION (YES) or not (NO).
	Hierarchy ObjectPrivilegeSummaryHierarchyEnum `mandatory:"false" json:"hierarchy,omitempty"`

	// The name of the object. The object can be any object, including tables, packages, indexes, sequences, and so on.
	Object *string `mandatory:"false" json:"object"`

	// Indicates whether the privilege is granted with the GRANT OPTION (YES) or not (NO).
	GrantOption ObjectPrivilegeSummaryGrantOptionEnum `mandatory:"false" json:"grantOption,omitempty"`

	// Indicates how the object privilege was granted. Possible values:
	// YES if the role is granted commonly (CONTAINER=ALL is used)
	// NO if the role is granted locally (CONTAINER=ALL is not used)
	Common ObjectPrivilegeSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the granted privilege is inherited from another container (YES) or not (NO).
	Inherited ObjectPrivilegeSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m ObjectPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectPrivilegeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectPrivilegeSummaryHierarchyEnum(string(m.Hierarchy)); !ok && m.Hierarchy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Hierarchy: %s. Supported values are: %s.", m.Hierarchy, strings.Join(GetObjectPrivilegeSummaryHierarchyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingObjectPrivilegeSummaryGrantOptionEnum(string(m.GrantOption)); !ok && m.GrantOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantOption: %s. Supported values are: %s.", m.GrantOption, strings.Join(GetObjectPrivilegeSummaryGrantOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingObjectPrivilegeSummaryCommonEnum(string(m.Common)); !ok && m.Common != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Common: %s. Supported values are: %s.", m.Common, strings.Join(GetObjectPrivilegeSummaryCommonEnumStringValues(), ",")))
	}
	if _, ok := GetMappingObjectPrivilegeSummaryInheritedEnum(string(m.Inherited)); !ok && m.Inherited != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inherited: %s. Supported values are: %s.", m.Inherited, strings.Join(GetObjectPrivilegeSummaryInheritedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ObjectPrivilegeSummaryHierarchyEnum Enum with underlying type: string
type ObjectPrivilegeSummaryHierarchyEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryHierarchyEnum
const (
	ObjectPrivilegeSummaryHierarchyYes ObjectPrivilegeSummaryHierarchyEnum = "YES"
	ObjectPrivilegeSummaryHierarchyNo  ObjectPrivilegeSummaryHierarchyEnum = "NO"
)

var mappingObjectPrivilegeSummaryHierarchyEnum = map[string]ObjectPrivilegeSummaryHierarchyEnum{
	"YES": ObjectPrivilegeSummaryHierarchyYes,
	"NO":  ObjectPrivilegeSummaryHierarchyNo,
}

var mappingObjectPrivilegeSummaryHierarchyEnumLowerCase = map[string]ObjectPrivilegeSummaryHierarchyEnum{
	"yes": ObjectPrivilegeSummaryHierarchyYes,
	"no":  ObjectPrivilegeSummaryHierarchyNo,
}

// GetObjectPrivilegeSummaryHierarchyEnumValues Enumerates the set of values for ObjectPrivilegeSummaryHierarchyEnum
func GetObjectPrivilegeSummaryHierarchyEnumValues() []ObjectPrivilegeSummaryHierarchyEnum {
	values := make([]ObjectPrivilegeSummaryHierarchyEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryHierarchyEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectPrivilegeSummaryHierarchyEnumStringValues Enumerates the set of values in String for ObjectPrivilegeSummaryHierarchyEnum
func GetObjectPrivilegeSummaryHierarchyEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingObjectPrivilegeSummaryHierarchyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectPrivilegeSummaryHierarchyEnum(val string) (ObjectPrivilegeSummaryHierarchyEnum, bool) {
	enum, ok := mappingObjectPrivilegeSummaryHierarchyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ObjectPrivilegeSummaryGrantOptionEnum Enum with underlying type: string
type ObjectPrivilegeSummaryGrantOptionEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryGrantOptionEnum
const (
	ObjectPrivilegeSummaryGrantOptionYes ObjectPrivilegeSummaryGrantOptionEnum = "YES"
	ObjectPrivilegeSummaryGrantOptionNo  ObjectPrivilegeSummaryGrantOptionEnum = "NO"
)

var mappingObjectPrivilegeSummaryGrantOptionEnum = map[string]ObjectPrivilegeSummaryGrantOptionEnum{
	"YES": ObjectPrivilegeSummaryGrantOptionYes,
	"NO":  ObjectPrivilegeSummaryGrantOptionNo,
}

var mappingObjectPrivilegeSummaryGrantOptionEnumLowerCase = map[string]ObjectPrivilegeSummaryGrantOptionEnum{
	"yes": ObjectPrivilegeSummaryGrantOptionYes,
	"no":  ObjectPrivilegeSummaryGrantOptionNo,
}

// GetObjectPrivilegeSummaryGrantOptionEnumValues Enumerates the set of values for ObjectPrivilegeSummaryGrantOptionEnum
func GetObjectPrivilegeSummaryGrantOptionEnumValues() []ObjectPrivilegeSummaryGrantOptionEnum {
	values := make([]ObjectPrivilegeSummaryGrantOptionEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryGrantOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectPrivilegeSummaryGrantOptionEnumStringValues Enumerates the set of values in String for ObjectPrivilegeSummaryGrantOptionEnum
func GetObjectPrivilegeSummaryGrantOptionEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingObjectPrivilegeSummaryGrantOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectPrivilegeSummaryGrantOptionEnum(val string) (ObjectPrivilegeSummaryGrantOptionEnum, bool) {
	enum, ok := mappingObjectPrivilegeSummaryGrantOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ObjectPrivilegeSummaryCommonEnum Enum with underlying type: string
type ObjectPrivilegeSummaryCommonEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryCommonEnum
const (
	ObjectPrivilegeSummaryCommonYes ObjectPrivilegeSummaryCommonEnum = "YES"
	ObjectPrivilegeSummaryCommonNo  ObjectPrivilegeSummaryCommonEnum = "NO"
)

var mappingObjectPrivilegeSummaryCommonEnum = map[string]ObjectPrivilegeSummaryCommonEnum{
	"YES": ObjectPrivilegeSummaryCommonYes,
	"NO":  ObjectPrivilegeSummaryCommonNo,
}

var mappingObjectPrivilegeSummaryCommonEnumLowerCase = map[string]ObjectPrivilegeSummaryCommonEnum{
	"yes": ObjectPrivilegeSummaryCommonYes,
	"no":  ObjectPrivilegeSummaryCommonNo,
}

// GetObjectPrivilegeSummaryCommonEnumValues Enumerates the set of values for ObjectPrivilegeSummaryCommonEnum
func GetObjectPrivilegeSummaryCommonEnumValues() []ObjectPrivilegeSummaryCommonEnum {
	values := make([]ObjectPrivilegeSummaryCommonEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryCommonEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectPrivilegeSummaryCommonEnumStringValues Enumerates the set of values in String for ObjectPrivilegeSummaryCommonEnum
func GetObjectPrivilegeSummaryCommonEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingObjectPrivilegeSummaryCommonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectPrivilegeSummaryCommonEnum(val string) (ObjectPrivilegeSummaryCommonEnum, bool) {
	enum, ok := mappingObjectPrivilegeSummaryCommonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ObjectPrivilegeSummaryInheritedEnum Enum with underlying type: string
type ObjectPrivilegeSummaryInheritedEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryInheritedEnum
const (
	ObjectPrivilegeSummaryInheritedYes ObjectPrivilegeSummaryInheritedEnum = "YES"
	ObjectPrivilegeSummaryInheritedNo  ObjectPrivilegeSummaryInheritedEnum = "NO"
)

var mappingObjectPrivilegeSummaryInheritedEnum = map[string]ObjectPrivilegeSummaryInheritedEnum{
	"YES": ObjectPrivilegeSummaryInheritedYes,
	"NO":  ObjectPrivilegeSummaryInheritedNo,
}

var mappingObjectPrivilegeSummaryInheritedEnumLowerCase = map[string]ObjectPrivilegeSummaryInheritedEnum{
	"yes": ObjectPrivilegeSummaryInheritedYes,
	"no":  ObjectPrivilegeSummaryInheritedNo,
}

// GetObjectPrivilegeSummaryInheritedEnumValues Enumerates the set of values for ObjectPrivilegeSummaryInheritedEnum
func GetObjectPrivilegeSummaryInheritedEnumValues() []ObjectPrivilegeSummaryInheritedEnum {
	values := make([]ObjectPrivilegeSummaryInheritedEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryInheritedEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectPrivilegeSummaryInheritedEnumStringValues Enumerates the set of values in String for ObjectPrivilegeSummaryInheritedEnum
func GetObjectPrivilegeSummaryInheritedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingObjectPrivilegeSummaryInheritedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectPrivilegeSummaryInheritedEnum(val string) (ObjectPrivilegeSummaryInheritedEnum, bool) {
	enum, ok := mappingObjectPrivilegeSummaryInheritedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
