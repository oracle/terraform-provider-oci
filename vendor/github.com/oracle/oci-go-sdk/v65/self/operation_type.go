// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateSubscription OperationTypeEnum = "CREATE_SUBSCRIPTION"
	OperationTypeUpdateSubscription OperationTypeEnum = "UPDATE_SUBSCRIPTION"
	OperationTypeDeleteSubscription OperationTypeEnum = "DELETE_SUBSCRIPTION"
	OperationTypeMoveSubscription   OperationTypeEnum = "MOVE_SUBSCRIPTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_SUBSCRIPTION": OperationTypeCreateSubscription,
	"UPDATE_SUBSCRIPTION": OperationTypeUpdateSubscription,
	"DELETE_SUBSCRIPTION": OperationTypeDeleteSubscription,
	"MOVE_SUBSCRIPTION":   OperationTypeMoveSubscription,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_subscription": OperationTypeCreateSubscription,
	"update_subscription": OperationTypeUpdateSubscription,
	"delete_subscription": OperationTypeDeleteSubscription,
	"move_subscription":   OperationTypeMoveSubscription,
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
		"CREATE_SUBSCRIPTION",
		"UPDATE_SUBSCRIPTION",
		"DELETE_SUBSCRIPTION",
		"MOVE_SUBSCRIPTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
