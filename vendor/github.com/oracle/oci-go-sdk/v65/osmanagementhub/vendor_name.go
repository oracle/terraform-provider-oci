// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// VendorNameEnum Enum with underlying type: string
type VendorNameEnum string

// Set of constants representing the allowable values for VendorNameEnum
const (
	VendorNameOracle    VendorNameEnum = "ORACLE"
	VendorNameMicrosoft VendorNameEnum = "MICROSOFT"
)

var mappingVendorNameEnum = map[string]VendorNameEnum{
	"ORACLE":    VendorNameOracle,
	"MICROSOFT": VendorNameMicrosoft,
}

var mappingVendorNameEnumLowerCase = map[string]VendorNameEnum{
	"oracle":    VendorNameOracle,
	"microsoft": VendorNameMicrosoft,
}

// GetVendorNameEnumValues Enumerates the set of values for VendorNameEnum
func GetVendorNameEnumValues() []VendorNameEnum {
	values := make([]VendorNameEnum, 0)
	for _, v := range mappingVendorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetVendorNameEnumStringValues Enumerates the set of values in String for VendorNameEnum
func GetVendorNameEnumStringValues() []string {
	return []string{
		"ORACLE",
		"MICROSOFT",
	}
}

// GetMappingVendorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVendorNameEnum(val string) (VendorNameEnum, bool) {
	enum, ok := mappingVendorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
