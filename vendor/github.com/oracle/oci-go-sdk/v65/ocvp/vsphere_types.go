// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// VsphereTypesEnum Enum with underlying type: string
type VsphereTypesEnum string

// Set of constants representing the allowable values for VsphereTypesEnum
const (
	VsphereTypesManagement VsphereTypesEnum = "MANAGEMENT"
	VsphereTypesWorkload   VsphereTypesEnum = "WORKLOAD"
)

var mappingVsphereTypesEnum = map[string]VsphereTypesEnum{
	"MANAGEMENT": VsphereTypesManagement,
	"WORKLOAD":   VsphereTypesWorkload,
}

var mappingVsphereTypesEnumLowerCase = map[string]VsphereTypesEnum{
	"management": VsphereTypesManagement,
	"workload":   VsphereTypesWorkload,
}

// GetVsphereTypesEnumValues Enumerates the set of values for VsphereTypesEnum
func GetVsphereTypesEnumValues() []VsphereTypesEnum {
	values := make([]VsphereTypesEnum, 0)
	for _, v := range mappingVsphereTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetVsphereTypesEnumStringValues Enumerates the set of values in String for VsphereTypesEnum
func GetVsphereTypesEnumStringValues() []string {
	return []string{
		"MANAGEMENT",
		"WORKLOAD",
	}
}

// GetMappingVsphereTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVsphereTypesEnum(val string) (VsphereTypesEnum, bool) {
	enum, ok := mappingVsphereTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
