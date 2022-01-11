// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// RecommendationLifecycleDetailEnum Enum with underlying type: string
type RecommendationLifecycleDetailEnum string

// Set of constants representing the allowable values for RecommendationLifecycleDetailEnum
const (
	RecommendationLifecycleDetailOpen      RecommendationLifecycleDetailEnum = "OPEN"
	RecommendationLifecycleDetailResolved  RecommendationLifecycleDetailEnum = "RESOLVED"
	RecommendationLifecycleDetailDismissed RecommendationLifecycleDetailEnum = "DISMISSED"
)

var mappingRecommendationLifecycleDetail = map[string]RecommendationLifecycleDetailEnum{
	"OPEN":      RecommendationLifecycleDetailOpen,
	"RESOLVED":  RecommendationLifecycleDetailResolved,
	"DISMISSED": RecommendationLifecycleDetailDismissed,
}

// GetRecommendationLifecycleDetailEnumValues Enumerates the set of values for RecommendationLifecycleDetailEnum
func GetRecommendationLifecycleDetailEnumValues() []RecommendationLifecycleDetailEnum {
	values := make([]RecommendationLifecycleDetailEnum, 0)
	for _, v := range mappingRecommendationLifecycleDetail {
		values = append(values, v)
	}
	return values
}
