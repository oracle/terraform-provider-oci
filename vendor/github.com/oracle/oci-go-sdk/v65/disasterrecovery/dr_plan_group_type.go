// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanGroupTypeEnum Enum with underlying type: string
type DrPlanGroupTypeEnum string

// Set of constants representing the allowable values for DrPlanGroupTypeEnum
const (
	DrPlanGroupTypeUserDefined      DrPlanGroupTypeEnum = "USER_DEFINED"
	DrPlanGroupTypeBuiltIn          DrPlanGroupTypeEnum = "BUILT_IN"
	DrPlanGroupTypeBuiltInPrecheck  DrPlanGroupTypeEnum = "BUILT_IN_PRECHECK"
	DrPlanGroupTypeUserDefinedPause DrPlanGroupTypeEnum = "USER_DEFINED_PAUSE"
)

var mappingDrPlanGroupTypeEnum = map[string]DrPlanGroupTypeEnum{
	"USER_DEFINED":       DrPlanGroupTypeUserDefined,
	"BUILT_IN":           DrPlanGroupTypeBuiltIn,
	"BUILT_IN_PRECHECK":  DrPlanGroupTypeBuiltInPrecheck,
	"USER_DEFINED_PAUSE": DrPlanGroupTypeUserDefinedPause,
}

var mappingDrPlanGroupTypeEnumLowerCase = map[string]DrPlanGroupTypeEnum{
	"user_defined":       DrPlanGroupTypeUserDefined,
	"built_in":           DrPlanGroupTypeBuiltIn,
	"built_in_precheck":  DrPlanGroupTypeBuiltInPrecheck,
	"user_defined_pause": DrPlanGroupTypeUserDefinedPause,
}

// GetDrPlanGroupTypeEnumValues Enumerates the set of values for DrPlanGroupTypeEnum
func GetDrPlanGroupTypeEnumValues() []DrPlanGroupTypeEnum {
	values := make([]DrPlanGroupTypeEnum, 0)
	for _, v := range mappingDrPlanGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanGroupTypeEnumStringValues Enumerates the set of values in String for DrPlanGroupTypeEnum
func GetDrPlanGroupTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"BUILT_IN",
		"BUILT_IN_PRECHECK",
		"USER_DEFINED_PAUSE",
	}
}

// GetMappingDrPlanGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanGroupTypeEnum(val string) (DrPlanGroupTypeEnum, bool) {
	enum, ok := mappingDrPlanGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
