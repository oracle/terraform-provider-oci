// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateZprPolicy        OperationTypeEnum = "CREATE_ZPR_POLICY"
	OperationTypeUpdateZprPolicy        OperationTypeEnum = "UPDATE_ZPR_POLICY"
	OperationTypeDeleteZprPolicy        OperationTypeEnum = "DELETE_ZPR_POLICY"
	OperationTypeCreateZprConfiguration OperationTypeEnum = "CREATE_ZPR_CONFIGURATION"
	OperationTypeUpdateZprConfiguration OperationTypeEnum = "UPDATE_ZPR_CONFIGURATION"
	OperationTypeDeleteZprConfiguration OperationTypeEnum = "DELETE_ZPR_CONFIGURATION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_ZPR_POLICY":        OperationTypeCreateZprPolicy,
	"UPDATE_ZPR_POLICY":        OperationTypeUpdateZprPolicy,
	"DELETE_ZPR_POLICY":        OperationTypeDeleteZprPolicy,
	"CREATE_ZPR_CONFIGURATION": OperationTypeCreateZprConfiguration,
	"UPDATE_ZPR_CONFIGURATION": OperationTypeUpdateZprConfiguration,
	"DELETE_ZPR_CONFIGURATION": OperationTypeDeleteZprConfiguration,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_zpr_policy":        OperationTypeCreateZprPolicy,
	"update_zpr_policy":        OperationTypeUpdateZprPolicy,
	"delete_zpr_policy":        OperationTypeDeleteZprPolicy,
	"create_zpr_configuration": OperationTypeCreateZprConfiguration,
	"update_zpr_configuration": OperationTypeUpdateZprConfiguration,
	"delete_zpr_configuration": OperationTypeDeleteZprConfiguration,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_ZPR_POLICY",
		"UPDATE_ZPR_POLICY",
		"DELETE_ZPR_POLICY",
		"CREATE_ZPR_CONFIGURATION",
		"UPDATE_ZPR_CONFIGURATION",
		"DELETE_ZPR_CONFIGURATION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
