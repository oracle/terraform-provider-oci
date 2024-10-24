// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ModelSettingCustomerNotificationTypeEnum Enum with underlying type: string
type ModelSettingCustomerNotificationTypeEnum string

// Set of constants representing the allowable values for ModelSettingCustomerNotificationTypeEnum
const (
	ModelSettingCustomerNotificationTypeNone      ModelSettingCustomerNotificationTypeEnum = "NONE"
	ModelSettingCustomerNotificationTypeAll       ModelSettingCustomerNotificationTypeEnum = "ALL"
	ModelSettingCustomerNotificationTypeOnFailure ModelSettingCustomerNotificationTypeEnum = "ON_FAILURE"
	ModelSettingCustomerNotificationTypeOnSuccess ModelSettingCustomerNotificationTypeEnum = "ON_SUCCESS"
)

var mappingModelSettingCustomerNotificationTypeEnum = map[string]ModelSettingCustomerNotificationTypeEnum{
	"NONE":       ModelSettingCustomerNotificationTypeNone,
	"ALL":        ModelSettingCustomerNotificationTypeAll,
	"ON_FAILURE": ModelSettingCustomerNotificationTypeOnFailure,
	"ON_SUCCESS": ModelSettingCustomerNotificationTypeOnSuccess,
}

var mappingModelSettingCustomerNotificationTypeEnumLowerCase = map[string]ModelSettingCustomerNotificationTypeEnum{
	"none":       ModelSettingCustomerNotificationTypeNone,
	"all":        ModelSettingCustomerNotificationTypeAll,
	"on_failure": ModelSettingCustomerNotificationTypeOnFailure,
	"on_success": ModelSettingCustomerNotificationTypeOnSuccess,
}

// GetModelSettingCustomerNotificationTypeEnumValues Enumerates the set of values for ModelSettingCustomerNotificationTypeEnum
func GetModelSettingCustomerNotificationTypeEnumValues() []ModelSettingCustomerNotificationTypeEnum {
	values := make([]ModelSettingCustomerNotificationTypeEnum, 0)
	for _, v := range mappingModelSettingCustomerNotificationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelSettingCustomerNotificationTypeEnumStringValues Enumerates the set of values in String for ModelSettingCustomerNotificationTypeEnum
func GetModelSettingCustomerNotificationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"ALL",
		"ON_FAILURE",
		"ON_SUCCESS",
	}
}

// GetMappingModelSettingCustomerNotificationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelSettingCustomerNotificationTypeEnum(val string) (ModelSettingCustomerNotificationTypeEnum, bool) {
	enum, ok := mappingModelSettingCustomerNotificationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
