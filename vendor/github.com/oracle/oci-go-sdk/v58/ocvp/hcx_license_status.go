// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// HcxLicenseStatusEnum Enum with underlying type: string
type HcxLicenseStatusEnum string

// Set of constants representing the allowable values for HcxLicenseStatusEnum
const (
	HcxLicenseStatusAvailable   HcxLicenseStatusEnum = "AVAILABLE"
	HcxLicenseStatusConsumed    HcxLicenseStatusEnum = "CONSUMED"
	HcxLicenseStatusDeactivated HcxLicenseStatusEnum = "DEACTIVATED"
	HcxLicenseStatusDeleted     HcxLicenseStatusEnum = "DELETED"
)

var mappingHcxLicenseStatusEnum = map[string]HcxLicenseStatusEnum{
	"AVAILABLE":   HcxLicenseStatusAvailable,
	"CONSUMED":    HcxLicenseStatusConsumed,
	"DEACTIVATED": HcxLicenseStatusDeactivated,
	"DELETED":     HcxLicenseStatusDeleted,
}

// GetHcxLicenseStatusEnumValues Enumerates the set of values for HcxLicenseStatusEnum
func GetHcxLicenseStatusEnumValues() []HcxLicenseStatusEnum {
	values := make([]HcxLicenseStatusEnum, 0)
	for _, v := range mappingHcxLicenseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHcxLicenseStatusEnumStringValues Enumerates the set of values in String for HcxLicenseStatusEnum
func GetHcxLicenseStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"CONSUMED",
		"DEACTIVATED",
		"DELETED",
	}
}

// GetMappingHcxLicenseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHcxLicenseStatusEnum(val string) (HcxLicenseStatusEnum, bool) {
	mappingHcxLicenseStatusEnumIgnoreCase := make(map[string]HcxLicenseStatusEnum)
	for k, v := range mappingHcxLicenseStatusEnum {
		mappingHcxLicenseStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHcxLicenseStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
