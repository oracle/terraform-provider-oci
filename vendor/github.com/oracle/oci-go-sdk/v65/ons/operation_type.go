// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreatePhoneApplication OperationTypeEnum = "CREATE_PHONE_APPLICATION"
	OperationTypeUpdatePhoneApplication OperationTypeEnum = "UPDATE_PHONE_APPLICATION"
	OperationTypeDeletePhoneApplication OperationTypeEnum = "DELETE_PHONE_APPLICATION"
	OperationTypeMovePhoneApplication   OperationTypeEnum = "MOVE_PHONE_APPLICATION"
	OperationTypeCreatePhoneNumber      OperationTypeEnum = "CREATE_PHONE_NUMBER"
	OperationTypeDeletePhoneNumber      OperationTypeEnum = "DELETE_PHONE_NUMBER"
	OperationTypeUpdatePhoneNumber      OperationTypeEnum = "UPDATE_PHONE_NUMBER"
	OperationTypeStartLogging           OperationTypeEnum = "START_LOGGING"
	OperationTypeUpdateLogging          OperationTypeEnum = "UPDATE_LOGGING"
	OperationTypeStopLogging            OperationTypeEnum = "STOP_LOGGING"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PHONE_APPLICATION": OperationTypeCreatePhoneApplication,
	"UPDATE_PHONE_APPLICATION": OperationTypeUpdatePhoneApplication,
	"DELETE_PHONE_APPLICATION": OperationTypeDeletePhoneApplication,
	"MOVE_PHONE_APPLICATION":   OperationTypeMovePhoneApplication,
	"CREATE_PHONE_NUMBER":      OperationTypeCreatePhoneNumber,
	"DELETE_PHONE_NUMBER":      OperationTypeDeletePhoneNumber,
	"UPDATE_PHONE_NUMBER":      OperationTypeUpdatePhoneNumber,
	"START_LOGGING":            OperationTypeStartLogging,
	"UPDATE_LOGGING":           OperationTypeUpdateLogging,
	"STOP_LOGGING":             OperationTypeStopLogging,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_phone_application": OperationTypeCreatePhoneApplication,
	"update_phone_application": OperationTypeUpdatePhoneApplication,
	"delete_phone_application": OperationTypeDeletePhoneApplication,
	"move_phone_application":   OperationTypeMovePhoneApplication,
	"create_phone_number":      OperationTypeCreatePhoneNumber,
	"delete_phone_number":      OperationTypeDeletePhoneNumber,
	"update_phone_number":      OperationTypeUpdatePhoneNumber,
	"start_logging":            OperationTypeStartLogging,
	"update_logging":           OperationTypeUpdateLogging,
	"stop_logging":             OperationTypeStopLogging,
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
		"CREATE_PHONE_APPLICATION",
		"UPDATE_PHONE_APPLICATION",
		"DELETE_PHONE_APPLICATION",
		"MOVE_PHONE_APPLICATION",
		"CREATE_PHONE_NUMBER",
		"DELETE_PHONE_NUMBER",
		"UPDATE_PHONE_NUMBER",
		"START_LOGGING",
		"UPDATE_LOGGING",
		"STOP_LOGGING",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
