// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"strings"
)

// DrgTypeEnum Enum with underlying type: string
type DrgTypeEnum string

// Set of constants representing the allowable values for DrgTypeEnum
const (
	DrgTypeRegional DrgTypeEnum = "REGIONAL"
	DrgTypeGlobal   DrgTypeEnum = "GLOBAL"
)

var mappingDrgTypeEnum = map[string]DrgTypeEnum{
	"REGIONAL": DrgTypeRegional,
	"GLOBAL":   DrgTypeGlobal,
}

var mappingDrgTypeEnumLowerCase = map[string]DrgTypeEnum{
	"regional": DrgTypeRegional,
	"global":   DrgTypeGlobal,
}

// GetDrgTypeEnumValues Enumerates the set of values for DrgTypeEnum
func GetDrgTypeEnumValues() []DrgTypeEnum {
	values := make([]DrgTypeEnum, 0)
	for _, v := range mappingDrgTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgTypeEnumStringValues Enumerates the set of values in String for DrgTypeEnum
func GetDrgTypeEnumStringValues() []string {
	return []string{
		"REGIONAL",
		"GLOBAL",
	}
}

// GetMappingDrgTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgTypeEnum(val string) (DrgTypeEnum, bool) {
	enum, ok := mappingDrgTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
