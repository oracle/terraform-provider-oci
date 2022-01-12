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

// OptimizerEnrollmentStatusEnum Enum with underlying type: string
type OptimizerEnrollmentStatusEnum string

// Set of constants representing the allowable values for OptimizerEnrollmentStatusEnum
const (
	OptimizerEnrollmentStatusActive   OptimizerEnrollmentStatusEnum = "ACTIVE"
	OptimizerEnrollmentStatusInactive OptimizerEnrollmentStatusEnum = "INACTIVE"
)

var mappingOptimizerEnrollmentStatus = map[string]OptimizerEnrollmentStatusEnum{
	"ACTIVE":   OptimizerEnrollmentStatusActive,
	"INACTIVE": OptimizerEnrollmentStatusInactive,
}

// GetOptimizerEnrollmentStatusEnumValues Enumerates the set of values for OptimizerEnrollmentStatusEnum
func GetOptimizerEnrollmentStatusEnumValues() []OptimizerEnrollmentStatusEnum {
	values := make([]OptimizerEnrollmentStatusEnum, 0)
	for _, v := range mappingOptimizerEnrollmentStatus {
		values = append(values, v)
	}
	return values
}
