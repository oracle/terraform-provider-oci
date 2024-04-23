// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// MetricNameEnum Enum with underlying type: string
type MetricNameEnum string

// Set of constants representing the allowable values for MetricNameEnum
const (
	MetricNameTotalInstanceCount                        MetricNameEnum = "TOTAL_INSTANCE_COUNT"
	MetricNameInstanceWithAvailableSecurityUpdatesCount MetricNameEnum = "INSTANCE_WITH_AVAILABLE_SECURITY_UPDATES_COUNT"
	MetricNameInstanceWithAvailableBugfixUpdatesCount   MetricNameEnum = "INSTANCE_WITH_AVAILABLE_BUGFIX_UPDATES_COUNT"
	MetricNameNormalInstanceCount                       MetricNameEnum = "NORMAL_INSTANCE_COUNT"
	MetricNameErrorInstanceCount                        MetricNameEnum = "ERROR_INSTANCE_COUNT"
	MetricNameWarningInstanceCount                      MetricNameEnum = "WARNING_INSTANCE_COUNT"
	MetricNameUnreachableInstanceCount                  MetricNameEnum = "UNREACHABLE_INSTANCE_COUNT"
	MetricNameRegistrationFailedInstanceCount           MetricNameEnum = "REGISTRATION_FAILED_INSTANCE_COUNT"
	MetricNameDeletingInstanceCount                     MetricNameEnum = "DELETING_INSTANCE_COUNT"
	MetricNameOnboardingInstanceCount                   MetricNameEnum = "ONBOARDING_INSTANCE_COUNT"
	MetricNameInstanceSecurityUpdatesCount              MetricNameEnum = "INSTANCE_SECURITY_UPDATES_COUNT"
	MetricNameInstanceBugfixUpdatesCount                MetricNameEnum = "INSTANCE_BUGFIX_UPDATES_COUNT"
	MetricNameInstanceSecurityAdvisoryCount             MetricNameEnum = "INSTANCE_SECURITY_ADVISORY_COUNT"
	MetricNameInstanceBugfixAdvisoryCount               MetricNameEnum = "INSTANCE_BUGFIX_ADVISORY_COUNT"
)

var mappingMetricNameEnum = map[string]MetricNameEnum{
	"TOTAL_INSTANCE_COUNT":                           MetricNameTotalInstanceCount,
	"INSTANCE_WITH_AVAILABLE_SECURITY_UPDATES_COUNT": MetricNameInstanceWithAvailableSecurityUpdatesCount,
	"INSTANCE_WITH_AVAILABLE_BUGFIX_UPDATES_COUNT":   MetricNameInstanceWithAvailableBugfixUpdatesCount,
	"NORMAL_INSTANCE_COUNT":                          MetricNameNormalInstanceCount,
	"ERROR_INSTANCE_COUNT":                           MetricNameErrorInstanceCount,
	"WARNING_INSTANCE_COUNT":                         MetricNameWarningInstanceCount,
	"UNREACHABLE_INSTANCE_COUNT":                     MetricNameUnreachableInstanceCount,
	"REGISTRATION_FAILED_INSTANCE_COUNT":             MetricNameRegistrationFailedInstanceCount,
	"DELETING_INSTANCE_COUNT":                        MetricNameDeletingInstanceCount,
	"ONBOARDING_INSTANCE_COUNT":                      MetricNameOnboardingInstanceCount,
	"INSTANCE_SECURITY_UPDATES_COUNT":                MetricNameInstanceSecurityUpdatesCount,
	"INSTANCE_BUGFIX_UPDATES_COUNT":                  MetricNameInstanceBugfixUpdatesCount,
	"INSTANCE_SECURITY_ADVISORY_COUNT":               MetricNameInstanceSecurityAdvisoryCount,
	"INSTANCE_BUGFIX_ADVISORY_COUNT":                 MetricNameInstanceBugfixAdvisoryCount,
}

var mappingMetricNameEnumLowerCase = map[string]MetricNameEnum{
	"total_instance_count":                           MetricNameTotalInstanceCount,
	"instance_with_available_security_updates_count": MetricNameInstanceWithAvailableSecurityUpdatesCount,
	"instance_with_available_bugfix_updates_count":   MetricNameInstanceWithAvailableBugfixUpdatesCount,
	"normal_instance_count":                          MetricNameNormalInstanceCount,
	"error_instance_count":                           MetricNameErrorInstanceCount,
	"warning_instance_count":                         MetricNameWarningInstanceCount,
	"unreachable_instance_count":                     MetricNameUnreachableInstanceCount,
	"registration_failed_instance_count":             MetricNameRegistrationFailedInstanceCount,
	"deleting_instance_count":                        MetricNameDeletingInstanceCount,
	"onboarding_instance_count":                      MetricNameOnboardingInstanceCount,
	"instance_security_updates_count":                MetricNameInstanceSecurityUpdatesCount,
	"instance_bugfix_updates_count":                  MetricNameInstanceBugfixUpdatesCount,
	"instance_security_advisory_count":               MetricNameInstanceSecurityAdvisoryCount,
	"instance_bugfix_advisory_count":                 MetricNameInstanceBugfixAdvisoryCount,
}

// GetMetricNameEnumValues Enumerates the set of values for MetricNameEnum
func GetMetricNameEnumValues() []MetricNameEnum {
	values := make([]MetricNameEnum, 0)
	for _, v := range mappingMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricNameEnumStringValues Enumerates the set of values in String for MetricNameEnum
func GetMetricNameEnumStringValues() []string {
	return []string{
		"TOTAL_INSTANCE_COUNT",
		"INSTANCE_WITH_AVAILABLE_SECURITY_UPDATES_COUNT",
		"INSTANCE_WITH_AVAILABLE_BUGFIX_UPDATES_COUNT",
		"NORMAL_INSTANCE_COUNT",
		"ERROR_INSTANCE_COUNT",
		"WARNING_INSTANCE_COUNT",
		"UNREACHABLE_INSTANCE_COUNT",
		"REGISTRATION_FAILED_INSTANCE_COUNT",
		"DELETING_INSTANCE_COUNT",
		"ONBOARDING_INSTANCE_COUNT",
		"INSTANCE_SECURITY_UPDATES_COUNT",
		"INSTANCE_BUGFIX_UPDATES_COUNT",
		"INSTANCE_SECURITY_ADVISORY_COUNT",
		"INSTANCE_BUGFIX_ADVISORY_COUNT",
	}
}

// GetMappingMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricNameEnum(val string) (MetricNameEnum, bool) {
	enum, ok := mappingMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
