// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// LifeCycleActionGroupTypeEnum Enum with underlying type: string
type LifeCycleActionGroupTypeEnum string

// Set of constants representing the allowable values for LifeCycleActionGroupTypeEnum
const (
	LifeCycleActionGroupTypeProduct     LifeCycleActionGroupTypeEnum = "PRODUCT"
	LifeCycleActionGroupTypeEnvironment LifeCycleActionGroupTypeEnum = "ENVIRONMENT"
)

var mappingLifeCycleActionGroupTypeEnum = map[string]LifeCycleActionGroupTypeEnum{
	"PRODUCT":     LifeCycleActionGroupTypeProduct,
	"ENVIRONMENT": LifeCycleActionGroupTypeEnvironment,
}

var mappingLifeCycleActionGroupTypeEnumLowerCase = map[string]LifeCycleActionGroupTypeEnum{
	"product":     LifeCycleActionGroupTypeProduct,
	"environment": LifeCycleActionGroupTypeEnvironment,
}

// GetLifeCycleActionGroupTypeEnumValues Enumerates the set of values for LifeCycleActionGroupTypeEnum
func GetLifeCycleActionGroupTypeEnumValues() []LifeCycleActionGroupTypeEnum {
	values := make([]LifeCycleActionGroupTypeEnum, 0)
	for _, v := range mappingLifeCycleActionGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLifeCycleActionGroupTypeEnumStringValues Enumerates the set of values in String for LifeCycleActionGroupTypeEnum
func GetLifeCycleActionGroupTypeEnumStringValues() []string {
	return []string{
		"PRODUCT",
		"ENVIRONMENT",
	}
}

// GetMappingLifeCycleActionGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifeCycleActionGroupTypeEnum(val string) (LifeCycleActionGroupTypeEnum, bool) {
	enum, ok := mappingLifeCycleActionGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
