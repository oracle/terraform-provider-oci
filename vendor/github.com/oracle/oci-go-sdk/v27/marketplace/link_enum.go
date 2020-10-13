// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// LinkEnumEnum Enum with underlying type: string
type LinkEnumEnum string

// Set of constants representing the allowable values for LinkEnumEnum
const (
	LinkEnumSelf      LinkEnumEnum = "SELF"
	LinkEnumCanonical LinkEnumEnum = "CANONICAL"
	LinkEnumNext      LinkEnumEnum = "NEXT"
	LinkEnumTemplate  LinkEnumEnum = "TEMPLATE"
	LinkEnumPrev      LinkEnumEnum = "PREV"
)

var mappingLinkEnum = map[string]LinkEnumEnum{
	"SELF":      LinkEnumSelf,
	"CANONICAL": LinkEnumCanonical,
	"NEXT":      LinkEnumNext,
	"TEMPLATE":  LinkEnumTemplate,
	"PREV":      LinkEnumPrev,
}

// GetLinkEnumEnumValues Enumerates the set of values for LinkEnumEnum
func GetLinkEnumEnumValues() []LinkEnumEnum {
	values := make([]LinkEnumEnum, 0)
	for _, v := range mappingLinkEnum {
		values = append(values, v)
	}
	return values
}
