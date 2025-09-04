// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"strings"
)

// SubscriptionTypeEnum Enum with underlying type: string
type SubscriptionTypeEnum string

// Set of constants representing the allowable values for SubscriptionTypeEnum
const (
	SubscriptionTypeOracledbatazure  SubscriptionTypeEnum = "ORACLEDBATAZURE"
	SubscriptionTypeOracledbatgoogle SubscriptionTypeEnum = "ORACLEDBATGOOGLE"
	SubscriptionTypeOracledbataws    SubscriptionTypeEnum = "ORACLEDBATAWS"
)

var mappingSubscriptionTypeEnum = map[string]SubscriptionTypeEnum{
	"ORACLEDBATAZURE":  SubscriptionTypeOracledbatazure,
	"ORACLEDBATGOOGLE": SubscriptionTypeOracledbatgoogle,
	"ORACLEDBATAWS":    SubscriptionTypeOracledbataws,
}

var mappingSubscriptionTypeEnumLowerCase = map[string]SubscriptionTypeEnum{
	"oracledbatazure":  SubscriptionTypeOracledbatazure,
	"oracledbatgoogle": SubscriptionTypeOracledbatgoogle,
	"oracledbataws":    SubscriptionTypeOracledbataws,
}

// GetSubscriptionTypeEnumValues Enumerates the set of values for SubscriptionTypeEnum
func GetSubscriptionTypeEnumValues() []SubscriptionTypeEnum {
	values := make([]SubscriptionTypeEnum, 0)
	for _, v := range mappingSubscriptionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionTypeEnumStringValues Enumerates the set of values in String for SubscriptionTypeEnum
func GetSubscriptionTypeEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingSubscriptionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionTypeEnum(val string) (SubscriptionTypeEnum, bool) {
	enum, ok := mappingSubscriptionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
