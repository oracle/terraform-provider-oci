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

// HcxModesEnum Enum with underlying type: string
type HcxModesEnum string

// Set of constants representing the allowable values for HcxModesEnum
const (
	HcxModesDisabled   HcxModesEnum = "DISABLED"
	HcxModesAdvanced   HcxModesEnum = "ADVANCED"
	HcxModesEnterprise HcxModesEnum = "ENTERPRISE"
)

var mappingHcxModesEnum = map[string]HcxModesEnum{
	"DISABLED":   HcxModesDisabled,
	"ADVANCED":   HcxModesAdvanced,
	"ENTERPRISE": HcxModesEnterprise,
}

var mappingHcxModesEnumLowerCase = map[string]HcxModesEnum{
	"disabled":   HcxModesDisabled,
	"advanced":   HcxModesAdvanced,
	"enterprise": HcxModesEnterprise,
}

// GetHcxModesEnumValues Enumerates the set of values for HcxModesEnum
func GetHcxModesEnumValues() []HcxModesEnum {
	values := make([]HcxModesEnum, 0)
	for _, v := range mappingHcxModesEnum {
		values = append(values, v)
	}
	return values
}

// GetHcxModesEnumStringValues Enumerates the set of values in String for HcxModesEnum
func GetHcxModesEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ADVANCED",
		"ENTERPRISE",
	}
}

// GetMappingHcxModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHcxModesEnum(val string) (HcxModesEnum, bool) {
	enum, ok := mappingHcxModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
