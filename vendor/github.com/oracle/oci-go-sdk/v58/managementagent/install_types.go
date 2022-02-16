// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// InstallTypesEnum Enum with underlying type: string
type InstallTypesEnum string

// Set of constants representing the allowable values for InstallTypesEnum
const (
	InstallTypesAgent   InstallTypesEnum = "AGENT"
	InstallTypesGateway InstallTypesEnum = "GATEWAY"
)

var mappingInstallTypesEnum = map[string]InstallTypesEnum{
	"AGENT":   InstallTypesAgent,
	"GATEWAY": InstallTypesGateway,
}

// GetInstallTypesEnumValues Enumerates the set of values for InstallTypesEnum
func GetInstallTypesEnumValues() []InstallTypesEnum {
	values := make([]InstallTypesEnum, 0)
	for _, v := range mappingInstallTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallTypesEnumStringValues Enumerates the set of values in String for InstallTypesEnum
func GetInstallTypesEnumStringValues() []string {
	return []string{
		"AGENT",
		"GATEWAY",
	}
}

// GetMappingInstallTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallTypesEnum(val string) (InstallTypesEnum, bool) {
	mappingInstallTypesEnumIgnoreCase := make(map[string]InstallTypesEnum)
	for k, v := range mappingInstallTypesEnum {
		mappingInstallTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingInstallTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
