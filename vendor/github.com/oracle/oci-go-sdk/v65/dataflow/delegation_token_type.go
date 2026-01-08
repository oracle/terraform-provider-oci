// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DelegationTokenTypeEnum Enum with underlying type: string
type DelegationTokenTypeEnum string

// Set of constants representing the allowable values for DelegationTokenTypeEnum
const (
	DelegationTokenTypeInstancePrincipal DelegationTokenTypeEnum = "INSTANCE_PRINCIPAL"
	DelegationTokenTypeResourcePrincipal DelegationTokenTypeEnum = "RESOURCE_PRINCIPAL"
)

var mappingDelegationTokenTypeEnum = map[string]DelegationTokenTypeEnum{
	"INSTANCE_PRINCIPAL": DelegationTokenTypeInstancePrincipal,
	"RESOURCE_PRINCIPAL": DelegationTokenTypeResourcePrincipal,
}

var mappingDelegationTokenTypeEnumLowerCase = map[string]DelegationTokenTypeEnum{
	"instance_principal": DelegationTokenTypeInstancePrincipal,
	"resource_principal": DelegationTokenTypeResourcePrincipal,
}

// GetDelegationTokenTypeEnumValues Enumerates the set of values for DelegationTokenTypeEnum
func GetDelegationTokenTypeEnumValues() []DelegationTokenTypeEnum {
	values := make([]DelegationTokenTypeEnum, 0)
	for _, v := range mappingDelegationTokenTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationTokenTypeEnumStringValues Enumerates the set of values in String for DelegationTokenTypeEnum
func GetDelegationTokenTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingDelegationTokenTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationTokenTypeEnum(val string) (DelegationTokenTypeEnum, bool) {
	enum, ok := mappingDelegationTokenTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
