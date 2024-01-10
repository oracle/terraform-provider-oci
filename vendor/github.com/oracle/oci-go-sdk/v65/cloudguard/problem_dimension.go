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

// ProblemDimensionEnum Enum with underlying type: string
type ProblemDimensionEnum string

// Set of constants representing the allowable values for ProblemDimensionEnum
const (
	ProblemDimensionResourceType  ProblemDimensionEnum = "RESOURCE_TYPE"
	ProblemDimensionRegion        ProblemDimensionEnum = "REGION"
	ProblemDimensionCompartmentId ProblemDimensionEnum = "COMPARTMENT_ID"
	ProblemDimensionRiskLevel     ProblemDimensionEnum = "RISK_LEVEL"
)

var mappingProblemDimensionEnum = map[string]ProblemDimensionEnum{
	"RESOURCE_TYPE":  ProblemDimensionResourceType,
	"REGION":         ProblemDimensionRegion,
	"COMPARTMENT_ID": ProblemDimensionCompartmentId,
	"RISK_LEVEL":     ProblemDimensionRiskLevel,
}

var mappingProblemDimensionEnumLowerCase = map[string]ProblemDimensionEnum{
	"resource_type":  ProblemDimensionResourceType,
	"region":         ProblemDimensionRegion,
	"compartment_id": ProblemDimensionCompartmentId,
	"risk_level":     ProblemDimensionRiskLevel,
}

// GetProblemDimensionEnumValues Enumerates the set of values for ProblemDimensionEnum
func GetProblemDimensionEnumValues() []ProblemDimensionEnum {
	values := make([]ProblemDimensionEnum, 0)
	for _, v := range mappingProblemDimensionEnum {
		values = append(values, v)
	}
	return values
}

// GetProblemDimensionEnumStringValues Enumerates the set of values in String for ProblemDimensionEnum
func GetProblemDimensionEnumStringValues() []string {
	return []string{
		"RESOURCE_TYPE",
		"REGION",
		"COMPARTMENT_ID",
		"RISK_LEVEL",
	}
}

// GetMappingProblemDimensionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProblemDimensionEnum(val string) (ProblemDimensionEnum, bool) {
	enum, ok := mappingProblemDimensionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
