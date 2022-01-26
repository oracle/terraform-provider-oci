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

// ObjectPrivilegeSummary Summary of objectPrivileges.
type ObjectPrivilegeSummary struct {

	// The name of the privilege on the object.
	Name *string `mandatory:"false" json:"name"`

	// The type of the object.
	SchemaType *string `mandatory:"false" json:"schemaType"`

	// The owner of the object.
	Owner *string `mandatory:"false" json:"owner"`

	// The name of the user who performed the grant
	Grantor *string `mandatory:"false" json:"grantor"`

	// Indicates whether the privilege was granted with the HIERARCHY OPTION (YES) or not (NO)
	Hierarchy ObjectPrivilegeSummaryHierarchyEnum `mandatory:"false" json:"hierarchy,omitempty"`

	// The name of the object. The object can be any object, including tables, packages, indexes, sequences, and so on.
	Object *string `mandatory:"false" json:"object"`

	// Indicates whether the privilege was granted with the GRANT OPTION (YES) or not (NO)
	GrantOption ObjectPrivilegeSummaryGrantOptionEnum `mandatory:"false" json:"grantOption,omitempty"`

	// Indicates how the grant was made. Possible values:
	// YES if the role was granted commonly (CONTAINER=ALL was used)
	// NO if the role was granted locally (CONTAINER=ALL was not used)
	Common ObjectPrivilegeSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the role grant was inherited from another container (YES) or not (NO)
	Inherited ObjectPrivilegeSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m ObjectPrivilegeSummary) String() string {
	return common.PointerString(m)
}

// ObjectPrivilegeSummaryHierarchyEnum Enum with underlying type: string
type ObjectPrivilegeSummaryHierarchyEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryHierarchyEnum
const (
	ObjectPrivilegeSummaryHierarchyYes ObjectPrivilegeSummaryHierarchyEnum = "YES"
	ObjectPrivilegeSummaryHierarchyNo  ObjectPrivilegeSummaryHierarchyEnum = "NO"
)

var mappingObjectPrivilegeSummaryHierarchy = map[string]ObjectPrivilegeSummaryHierarchyEnum{
	"YES": ObjectPrivilegeSummaryHierarchyYes,
	"NO":  ObjectPrivilegeSummaryHierarchyNo,
}

// GetObjectPrivilegeSummaryHierarchyEnumValues Enumerates the set of values for ObjectPrivilegeSummaryHierarchyEnum
func GetObjectPrivilegeSummaryHierarchyEnumValues() []ObjectPrivilegeSummaryHierarchyEnum {
	values := make([]ObjectPrivilegeSummaryHierarchyEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryHierarchy {
		values = append(values, v)
	}
	return values
}

// ObjectPrivilegeSummaryGrantOptionEnum Enum with underlying type: string
type ObjectPrivilegeSummaryGrantOptionEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryGrantOptionEnum
const (
	ObjectPrivilegeSummaryGrantOptionYes ObjectPrivilegeSummaryGrantOptionEnum = "YES"
	ObjectPrivilegeSummaryGrantOptionNo  ObjectPrivilegeSummaryGrantOptionEnum = "NO"
)

var mappingObjectPrivilegeSummaryGrantOption = map[string]ObjectPrivilegeSummaryGrantOptionEnum{
	"YES": ObjectPrivilegeSummaryGrantOptionYes,
	"NO":  ObjectPrivilegeSummaryGrantOptionNo,
}

// GetObjectPrivilegeSummaryGrantOptionEnumValues Enumerates the set of values for ObjectPrivilegeSummaryGrantOptionEnum
func GetObjectPrivilegeSummaryGrantOptionEnumValues() []ObjectPrivilegeSummaryGrantOptionEnum {
	values := make([]ObjectPrivilegeSummaryGrantOptionEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryGrantOption {
		values = append(values, v)
	}
	return values
}

// ObjectPrivilegeSummaryCommonEnum Enum with underlying type: string
type ObjectPrivilegeSummaryCommonEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryCommonEnum
const (
	ObjectPrivilegeSummaryCommonYes ObjectPrivilegeSummaryCommonEnum = "YES"
	ObjectPrivilegeSummaryCommonNo  ObjectPrivilegeSummaryCommonEnum = "NO"
)

var mappingObjectPrivilegeSummaryCommon = map[string]ObjectPrivilegeSummaryCommonEnum{
	"YES": ObjectPrivilegeSummaryCommonYes,
	"NO":  ObjectPrivilegeSummaryCommonNo,
}

// GetObjectPrivilegeSummaryCommonEnumValues Enumerates the set of values for ObjectPrivilegeSummaryCommonEnum
func GetObjectPrivilegeSummaryCommonEnumValues() []ObjectPrivilegeSummaryCommonEnum {
	values := make([]ObjectPrivilegeSummaryCommonEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryCommon {
		values = append(values, v)
	}
	return values
}

// ObjectPrivilegeSummaryInheritedEnum Enum with underlying type: string
type ObjectPrivilegeSummaryInheritedEnum string

// Set of constants representing the allowable values for ObjectPrivilegeSummaryInheritedEnum
const (
	ObjectPrivilegeSummaryInheritedYes ObjectPrivilegeSummaryInheritedEnum = "YES"
	ObjectPrivilegeSummaryInheritedNo  ObjectPrivilegeSummaryInheritedEnum = "NO"
)

var mappingObjectPrivilegeSummaryInherited = map[string]ObjectPrivilegeSummaryInheritedEnum{
	"YES": ObjectPrivilegeSummaryInheritedYes,
	"NO":  ObjectPrivilegeSummaryInheritedNo,
}

// GetObjectPrivilegeSummaryInheritedEnumValues Enumerates the set of values for ObjectPrivilegeSummaryInheritedEnum
func GetObjectPrivilegeSummaryInheritedEnumValues() []ObjectPrivilegeSummaryInheritedEnum {
	values := make([]ObjectPrivilegeSummaryInheritedEnum, 0)
	for _, v := range mappingObjectPrivilegeSummaryInherited {
		values = append(values, v)
	}
	return values
}
