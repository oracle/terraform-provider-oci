// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
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

var mappingInstallTypesEnumLowerCase = map[string]InstallTypesEnum{
	"agent":   InstallTypesAgent,
	"gateway": InstallTypesGateway,
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
	enum, ok := mappingInstallTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
