// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// BusinessInsightsCategoryEnum Enum with underlying type: string
type BusinessInsightsCategoryEnum string

// Set of constants representing the allowable values for BusinessInsightsCategoryEnum
const (
	BusinessInsightsCategoryApplicationPerformanceAndAvailability BusinessInsightsCategoryEnum = "APPLICATION_PERFORMANCE_AND_AVAILABILITY"
	BusinessInsightsCategoryErrorDetectionAndTroubleshooting      BusinessInsightsCategoryEnum = "ERROR_DETECTION_AND_TROUBLESHOOTING"
	BusinessInsightsCategorySecurityAndCompliance                 BusinessInsightsCategoryEnum = "SECURITY_AND_COMPLIANCE"
	BusinessInsightsCategoryCapacityPlanningAndUsageAnalytics     BusinessInsightsCategoryEnum = "CAPACITY_PLANNING_AND_USAGE_ANALYTICS"
	BusinessInsightsCategoryIntegrationAndDependencyMonitoring    BusinessInsightsCategoryEnum = "INTEGRATION_AND_DEPENDENCY_MONITORING"
)

var mappingBusinessInsightsCategoryEnum = map[string]BusinessInsightsCategoryEnum{
	"APPLICATION_PERFORMANCE_AND_AVAILABILITY": BusinessInsightsCategoryApplicationPerformanceAndAvailability,
	"ERROR_DETECTION_AND_TROUBLESHOOTING":      BusinessInsightsCategoryErrorDetectionAndTroubleshooting,
	"SECURITY_AND_COMPLIANCE":                  BusinessInsightsCategorySecurityAndCompliance,
	"CAPACITY_PLANNING_AND_USAGE_ANALYTICS":    BusinessInsightsCategoryCapacityPlanningAndUsageAnalytics,
	"INTEGRATION_AND_DEPENDENCY_MONITORING":    BusinessInsightsCategoryIntegrationAndDependencyMonitoring,
}

var mappingBusinessInsightsCategoryEnumLowerCase = map[string]BusinessInsightsCategoryEnum{
	"application_performance_and_availability": BusinessInsightsCategoryApplicationPerformanceAndAvailability,
	"error_detection_and_troubleshooting":      BusinessInsightsCategoryErrorDetectionAndTroubleshooting,
	"security_and_compliance":                  BusinessInsightsCategorySecurityAndCompliance,
	"capacity_planning_and_usage_analytics":    BusinessInsightsCategoryCapacityPlanningAndUsageAnalytics,
	"integration_and_dependency_monitoring":    BusinessInsightsCategoryIntegrationAndDependencyMonitoring,
}

// GetBusinessInsightsCategoryEnumValues Enumerates the set of values for BusinessInsightsCategoryEnum
func GetBusinessInsightsCategoryEnumValues() []BusinessInsightsCategoryEnum {
	values := make([]BusinessInsightsCategoryEnum, 0)
	for _, v := range mappingBusinessInsightsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetBusinessInsightsCategoryEnumStringValues Enumerates the set of values in String for BusinessInsightsCategoryEnum
func GetBusinessInsightsCategoryEnumStringValues() []string {
	return []string{
		"APPLICATION_PERFORMANCE_AND_AVAILABILITY",
		"ERROR_DETECTION_AND_TROUBLESHOOTING",
		"SECURITY_AND_COMPLIANCE",
		"CAPACITY_PLANNING_AND_USAGE_ANALYTICS",
		"INTEGRATION_AND_DEPENDENCY_MONITORING",
	}
}

// GetMappingBusinessInsightsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBusinessInsightsCategoryEnum(val string) (BusinessInsightsCategoryEnum, bool) {
	enum, ok := mappingBusinessInsightsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
