// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"strings"
)

// NamespaceEnum Enum with underlying type: string
type NamespaceEnum string

// Set of constants representing the allowable values for NamespaceEnum
const (
	NamespaceCompute NamespaceEnum = "COMPUTE"
)

var mappingNamespaceEnum = map[string]NamespaceEnum{
	"COMPUTE": NamespaceCompute,
}

var mappingNamespaceEnumLowerCase = map[string]NamespaceEnum{
	"compute": NamespaceCompute,
}

// GetNamespaceEnumValues Enumerates the set of values for NamespaceEnum
func GetNamespaceEnumValues() []NamespaceEnum {
	values := make([]NamespaceEnum, 0)
	for _, v := range mappingNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetNamespaceEnumStringValues Enumerates the set of values in String for NamespaceEnum
func GetNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamespaceEnum(val string) (NamespaceEnum, bool) {
	enum, ok := mappingNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
