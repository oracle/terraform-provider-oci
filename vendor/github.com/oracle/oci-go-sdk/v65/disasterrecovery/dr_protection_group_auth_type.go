// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrProtectionGroupAuthTypeEnum Enum with underlying type: string
type DrProtectionGroupAuthTypeEnum string

// Set of constants representing the allowable values for DrProtectionGroupAuthTypeEnum
const (
	DrProtectionGroupAuthTypeObo               DrProtectionGroupAuthTypeEnum = "OBO"
	DrProtectionGroupAuthTypeResourcePrincipal DrProtectionGroupAuthTypeEnum = "RESOURCE_PRINCIPAL"
)

var mappingDrProtectionGroupAuthTypeEnum = map[string]DrProtectionGroupAuthTypeEnum{
	"OBO":                DrProtectionGroupAuthTypeObo,
	"RESOURCE_PRINCIPAL": DrProtectionGroupAuthTypeResourcePrincipal,
}

var mappingDrProtectionGroupAuthTypeEnumLowerCase = map[string]DrProtectionGroupAuthTypeEnum{
	"obo":                DrProtectionGroupAuthTypeObo,
	"resource_principal": DrProtectionGroupAuthTypeResourcePrincipal,
}

// GetDrProtectionGroupAuthTypeEnumValues Enumerates the set of values for DrProtectionGroupAuthTypeEnum
func GetDrProtectionGroupAuthTypeEnumValues() []DrProtectionGroupAuthTypeEnum {
	values := make([]DrProtectionGroupAuthTypeEnum, 0)
	for _, v := range mappingDrProtectionGroupAuthTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrProtectionGroupAuthTypeEnumStringValues Enumerates the set of values in String for DrProtectionGroupAuthTypeEnum
func GetDrProtectionGroupAuthTypeEnumStringValues() []string {
	return []string{
		"OBO",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingDrProtectionGroupAuthTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupAuthTypeEnum(val string) (DrProtectionGroupAuthTypeEnum, bool) {
	enum, ok := mappingDrProtectionGroupAuthTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
