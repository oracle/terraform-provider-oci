// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// ManagedInstanceTypeEnum Enum with underlying type: string
type ManagedInstanceTypeEnum string

// Set of constants representing the allowable values for ManagedInstanceTypeEnum
const (
	ManagedInstanceTypeOracleManagementAgent ManagedInstanceTypeEnum = "ORACLE_MANAGEMENT_AGENT"
	ManagedInstanceTypeOracleCloudAgent      ManagedInstanceTypeEnum = "ORACLE_CLOUD_AGENT"
)

var mappingManagedInstanceTypeEnum = map[string]ManagedInstanceTypeEnum{
	"ORACLE_MANAGEMENT_AGENT": ManagedInstanceTypeOracleManagementAgent,
	"ORACLE_CLOUD_AGENT":      ManagedInstanceTypeOracleCloudAgent,
}

var mappingManagedInstanceTypeEnumLowerCase = map[string]ManagedInstanceTypeEnum{
	"oracle_management_agent": ManagedInstanceTypeOracleManagementAgent,
	"oracle_cloud_agent":      ManagedInstanceTypeOracleCloudAgent,
}

// GetManagedInstanceTypeEnumValues Enumerates the set of values for ManagedInstanceTypeEnum
func GetManagedInstanceTypeEnumValues() []ManagedInstanceTypeEnum {
	values := make([]ManagedInstanceTypeEnum, 0)
	for _, v := range mappingManagedInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedInstanceTypeEnumStringValues Enumerates the set of values in String for ManagedInstanceTypeEnum
func GetManagedInstanceTypeEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGEMENT_AGENT",
		"ORACLE_CLOUD_AGENT",
	}
}

// GetMappingManagedInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedInstanceTypeEnum(val string) (ManagedInstanceTypeEnum, bool) {
	enum, ok := mappingManagedInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
