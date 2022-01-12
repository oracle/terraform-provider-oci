// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ProblemDimensionEnum Enum with underlying type: string
type ProblemDimensionEnum string

// Set of constants representing the allowable values for ProblemDimensionEnum
const (
	ProblemDimensionResourceType  ProblemDimensionEnum = "RESOURCE_TYPE"
	ProblemDimensionRegion        ProblemDimensionEnum = "REGION"
	ProblemDimensionCompartmentId ProblemDimensionEnum = "COMPARTMENT_ID"
	ProblemDimensionRiskLevel     ProblemDimensionEnum = "RISK_LEVEL"
)

var mappingProblemDimension = map[string]ProblemDimensionEnum{
	"RESOURCE_TYPE":  ProblemDimensionResourceType,
	"REGION":         ProblemDimensionRegion,
	"COMPARTMENT_ID": ProblemDimensionCompartmentId,
	"RISK_LEVEL":     ProblemDimensionRiskLevel,
}

// GetProblemDimensionEnumValues Enumerates the set of values for ProblemDimensionEnum
func GetProblemDimensionEnumValues() []ProblemDimensionEnum {
	values := make([]ProblemDimensionEnum, 0)
	for _, v := range mappingProblemDimension {
		values = append(values, v)
	}
	return values
}
