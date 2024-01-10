// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"strings"
)

// OperatorActionCategoriesEnum Enum with underlying type: string
type OperatorActionCategoriesEnum string

// Set of constants representing the allowable values for OperatorActionCategoriesEnum
const (
	OperatorActionCategoriesServicediagnostics OperatorActionCategoriesEnum = "SERVICEDIAGNOSTICS"
	OperatorActionCategoriesServicemanagement  OperatorActionCategoriesEnum = "SERVICEMANAGEMENT"
	OperatorActionCategoriesSystemmanagement   OperatorActionCategoriesEnum = "SYSTEMMANAGEMENT"
	OperatorActionCategoriesFulladministration OperatorActionCategoriesEnum = "FULLADMINISTRATION"
	OperatorActionCategoriesCellmanagement     OperatorActionCategoriesEnum = "CELLMANAGEMENT"
)

var mappingOperatorActionCategoriesEnum = map[string]OperatorActionCategoriesEnum{
	"SERVICEDIAGNOSTICS": OperatorActionCategoriesServicediagnostics,
	"SERVICEMANAGEMENT":  OperatorActionCategoriesServicemanagement,
	"SYSTEMMANAGEMENT":   OperatorActionCategoriesSystemmanagement,
	"FULLADMINISTRATION": OperatorActionCategoriesFulladministration,
	"CELLMANAGEMENT":     OperatorActionCategoriesCellmanagement,
}

var mappingOperatorActionCategoriesEnumLowerCase = map[string]OperatorActionCategoriesEnum{
	"servicediagnostics": OperatorActionCategoriesServicediagnostics,
	"servicemanagement":  OperatorActionCategoriesServicemanagement,
	"systemmanagement":   OperatorActionCategoriesSystemmanagement,
	"fulladministration": OperatorActionCategoriesFulladministration,
	"cellmanagement":     OperatorActionCategoriesCellmanagement,
}

// GetOperatorActionCategoriesEnumValues Enumerates the set of values for OperatorActionCategoriesEnum
func GetOperatorActionCategoriesEnumValues() []OperatorActionCategoriesEnum {
	values := make([]OperatorActionCategoriesEnum, 0)
	for _, v := range mappingOperatorActionCategoriesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorActionCategoriesEnumStringValues Enumerates the set of values in String for OperatorActionCategoriesEnum
func GetOperatorActionCategoriesEnumStringValues() []string {
	return []string{
		"SERVICEDIAGNOSTICS",
		"SERVICEMANAGEMENT",
		"SYSTEMMANAGEMENT",
		"FULLADMINISTRATION",
		"CELLMANAGEMENT",
	}
}

// GetMappingOperatorActionCategoriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorActionCategoriesEnum(val string) (OperatorActionCategoriesEnum, bool) {
	enum, ok := mappingOperatorActionCategoriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
