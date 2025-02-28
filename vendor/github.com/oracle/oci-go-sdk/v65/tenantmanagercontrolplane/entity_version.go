// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// EntityVersionEnum Enum with underlying type: string
type EntityVersionEnum string

// Set of constants representing the allowable values for EntityVersionEnum
const (
	EntityVersionV1 EntityVersionEnum = "V1"
	EntityVersionV2 EntityVersionEnum = "V2"
)

var mappingEntityVersionEnum = map[string]EntityVersionEnum{
	"V1": EntityVersionV1,
	"V2": EntityVersionV2,
}

var mappingEntityVersionEnumLowerCase = map[string]EntityVersionEnum{
	"v1": EntityVersionV1,
	"v2": EntityVersionV2,
}

// GetEntityVersionEnumValues Enumerates the set of values for EntityVersionEnum
func GetEntityVersionEnumValues() []EntityVersionEnum {
	values := make([]EntityVersionEnum, 0)
	for _, v := range mappingEntityVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityVersionEnumStringValues Enumerates the set of values in String for EntityVersionEnum
func GetEntityVersionEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
	}
}

// GetMappingEntityVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityVersionEnum(val string) (EntityVersionEnum, bool) {
	enum, ok := mappingEntityVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
