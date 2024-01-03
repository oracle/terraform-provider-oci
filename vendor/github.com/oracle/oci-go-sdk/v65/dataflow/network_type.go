// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// NetworkTypeEnum Enum with underlying type: string
type NetworkTypeEnum string

// Set of constants representing the allowable values for NetworkTypeEnum
const (
	NetworkTypeVcn          NetworkTypeEnum = "VCN"
	NetworkTypeSecureAccess NetworkTypeEnum = "SECURE_ACCESS"
)

var mappingNetworkTypeEnum = map[string]NetworkTypeEnum{
	"VCN":           NetworkTypeVcn,
	"SECURE_ACCESS": NetworkTypeSecureAccess,
}

var mappingNetworkTypeEnumLowerCase = map[string]NetworkTypeEnum{
	"vcn":           NetworkTypeVcn,
	"secure_access": NetworkTypeSecureAccess,
}

// GetNetworkTypeEnumValues Enumerates the set of values for NetworkTypeEnum
func GetNetworkTypeEnumValues() []NetworkTypeEnum {
	values := make([]NetworkTypeEnum, 0)
	for _, v := range mappingNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkTypeEnumStringValues Enumerates the set of values in String for NetworkTypeEnum
func GetNetworkTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"SECURE_ACCESS",
	}
}

// GetMappingNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkTypeEnum(val string) (NetworkTypeEnum, bool) {
	enum, ok := mappingNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
