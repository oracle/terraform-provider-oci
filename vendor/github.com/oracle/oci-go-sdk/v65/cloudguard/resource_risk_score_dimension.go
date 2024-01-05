// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ResourceRiskScoreDimensionEnum Enum with underlying type: string
type ResourceRiskScoreDimensionEnum string

// Set of constants representing the allowable values for ResourceRiskScoreDimensionEnum
const (
	ResourceRiskScoreDimensionProblemId         ResourceRiskScoreDimensionEnum = "PROBLEM_ID"
	ResourceRiskScoreDimensionResourceProfileId ResourceRiskScoreDimensionEnum = "RESOURCE_PROFILE_ID"
)

var mappingResourceRiskScoreDimensionEnum = map[string]ResourceRiskScoreDimensionEnum{
	"PROBLEM_ID":          ResourceRiskScoreDimensionProblemId,
	"RESOURCE_PROFILE_ID": ResourceRiskScoreDimensionResourceProfileId,
}

var mappingResourceRiskScoreDimensionEnumLowerCase = map[string]ResourceRiskScoreDimensionEnum{
	"problem_id":          ResourceRiskScoreDimensionProblemId,
	"resource_profile_id": ResourceRiskScoreDimensionResourceProfileId,
}

// GetResourceRiskScoreDimensionEnumValues Enumerates the set of values for ResourceRiskScoreDimensionEnum
func GetResourceRiskScoreDimensionEnumValues() []ResourceRiskScoreDimensionEnum {
	values := make([]ResourceRiskScoreDimensionEnum, 0)
	for _, v := range mappingResourceRiskScoreDimensionEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceRiskScoreDimensionEnumStringValues Enumerates the set of values in String for ResourceRiskScoreDimensionEnum
func GetResourceRiskScoreDimensionEnumStringValues() []string {
	return []string{
		"PROBLEM_ID",
		"RESOURCE_PROFILE_ID",
	}
}

// GetMappingResourceRiskScoreDimensionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceRiskScoreDimensionEnum(val string) (ResourceRiskScoreDimensionEnum, bool) {
	enum, ok := mappingResourceRiskScoreDimensionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
