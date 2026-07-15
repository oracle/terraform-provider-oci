// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// InfrastructureTypeEnum Enum with underlying type: string
type InfrastructureTypeEnum string

// Set of constants representing the allowable values for InfrastructureTypeEnum
const (
	InfrastructureTypeCloud           InfrastructureTypeEnum = "CLOUD"
	InfrastructureTypeCloudAtCustomer InfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingInfrastructureTypeEnum = map[string]InfrastructureTypeEnum{
	"CLOUD":             InfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": InfrastructureTypeCloudAtCustomer,
}

var mappingInfrastructureTypeEnumLowerCase = map[string]InfrastructureTypeEnum{
	"cloud":             InfrastructureTypeCloud,
	"cloud_at_customer": InfrastructureTypeCloudAtCustomer,
}

// GetInfrastructureTypeEnumValues Enumerates the set of values for InfrastructureTypeEnum
func GetInfrastructureTypeEnumValues() []InfrastructureTypeEnum {
	values := make([]InfrastructureTypeEnum, 0)
	for _, v := range mappingInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureTypeEnumStringValues Enumerates the set of values in String for InfrastructureTypeEnum
func GetInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureTypeEnum(val string) (InfrastructureTypeEnum, bool) {
	enum, ok := mappingInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
