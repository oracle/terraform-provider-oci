// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// OptimizerEnrollmentStatusEnum Enum with underlying type: string
type OptimizerEnrollmentStatusEnum string

// Set of constants representing the allowable values for OptimizerEnrollmentStatusEnum
const (
	OptimizerEnrollmentStatusActive   OptimizerEnrollmentStatusEnum = "ACTIVE"
	OptimizerEnrollmentStatusInactive OptimizerEnrollmentStatusEnum = "INACTIVE"
)

var mappingOptimizerEnrollmentStatusEnum = map[string]OptimizerEnrollmentStatusEnum{
	"ACTIVE":   OptimizerEnrollmentStatusActive,
	"INACTIVE": OptimizerEnrollmentStatusInactive,
}

// GetOptimizerEnrollmentStatusEnumValues Enumerates the set of values for OptimizerEnrollmentStatusEnum
func GetOptimizerEnrollmentStatusEnumValues() []OptimizerEnrollmentStatusEnum {
	values := make([]OptimizerEnrollmentStatusEnum, 0)
	for _, v := range mappingOptimizerEnrollmentStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerEnrollmentStatusEnumStringValues Enumerates the set of values in String for OptimizerEnrollmentStatusEnum
func GetOptimizerEnrollmentStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingOptimizerEnrollmentStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerEnrollmentStatusEnum(val string) (OptimizerEnrollmentStatusEnum, bool) {
	mappingOptimizerEnrollmentStatusEnumIgnoreCase := make(map[string]OptimizerEnrollmentStatusEnum)
	for k, v := range mappingOptimizerEnrollmentStatusEnum {
		mappingOptimizerEnrollmentStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOptimizerEnrollmentStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
