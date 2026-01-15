// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MemberReferenceTypeEnum Enum with underlying type: string
type MemberReferenceTypeEnum string

// Set of constants representing the allowable values for MemberReferenceTypeEnum
const (
	MemberReferenceTypeResourceInstance MemberReferenceTypeEnum = "RESOURCE_INSTANCE"
	MemberReferenceTypeResourceType     MemberReferenceTypeEnum = "RESOURCE_TYPE"
	MemberReferenceTypeResourceGroup    MemberReferenceTypeEnum = "RESOURCE_GROUP"
)

var mappingMemberReferenceTypeEnum = map[string]MemberReferenceTypeEnum{
	"RESOURCE_INSTANCE": MemberReferenceTypeResourceInstance,
	"RESOURCE_TYPE":     MemberReferenceTypeResourceType,
	"RESOURCE_GROUP":    MemberReferenceTypeResourceGroup,
}

var mappingMemberReferenceTypeEnumLowerCase = map[string]MemberReferenceTypeEnum{
	"resource_instance": MemberReferenceTypeResourceInstance,
	"resource_type":     MemberReferenceTypeResourceType,
	"resource_group":    MemberReferenceTypeResourceGroup,
}

// GetMemberReferenceTypeEnumValues Enumerates the set of values for MemberReferenceTypeEnum
func GetMemberReferenceTypeEnumValues() []MemberReferenceTypeEnum {
	values := make([]MemberReferenceTypeEnum, 0)
	for _, v := range mappingMemberReferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMemberReferenceTypeEnumStringValues Enumerates the set of values in String for MemberReferenceTypeEnum
func GetMemberReferenceTypeEnumStringValues() []string {
	return []string{
		"RESOURCE_INSTANCE",
		"RESOURCE_TYPE",
		"RESOURCE_GROUP",
	}
}

// GetMappingMemberReferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMemberReferenceTypeEnum(val string) (MemberReferenceTypeEnum, bool) {
	enum, ok := mappingMemberReferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
