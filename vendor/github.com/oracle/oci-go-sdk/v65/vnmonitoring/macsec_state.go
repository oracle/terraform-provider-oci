// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"strings"
)

// MacsecStateEnum Enum with underlying type: string
type MacsecStateEnum string

// Set of constants representing the allowable values for MacsecStateEnum
const (
	MacsecStateEnabled  MacsecStateEnum = "ENABLED"
	MacsecStateDisabled MacsecStateEnum = "DISABLED"
)

var mappingMacsecStateEnum = map[string]MacsecStateEnum{
	"ENABLED":  MacsecStateEnabled,
	"DISABLED": MacsecStateDisabled,
}

var mappingMacsecStateEnumLowerCase = map[string]MacsecStateEnum{
	"enabled":  MacsecStateEnabled,
	"disabled": MacsecStateDisabled,
}

// GetMacsecStateEnumValues Enumerates the set of values for MacsecStateEnum
func GetMacsecStateEnumValues() []MacsecStateEnum {
	values := make([]MacsecStateEnum, 0)
	for _, v := range mappingMacsecStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMacsecStateEnumStringValues Enumerates the set of values in String for MacsecStateEnum
func GetMacsecStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingMacsecStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacsecStateEnum(val string) (MacsecStateEnum, bool) {
	enum, ok := mappingMacsecStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
