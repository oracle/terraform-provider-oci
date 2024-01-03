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

// RecommendationLifecycleDetailEnum Enum with underlying type: string
type RecommendationLifecycleDetailEnum string

// Set of constants representing the allowable values for RecommendationLifecycleDetailEnum
const (
	RecommendationLifecycleDetailOpen      RecommendationLifecycleDetailEnum = "OPEN"
	RecommendationLifecycleDetailResolved  RecommendationLifecycleDetailEnum = "RESOLVED"
	RecommendationLifecycleDetailDismissed RecommendationLifecycleDetailEnum = "DISMISSED"
)

var mappingRecommendationLifecycleDetailEnum = map[string]RecommendationLifecycleDetailEnum{
	"OPEN":      RecommendationLifecycleDetailOpen,
	"RESOLVED":  RecommendationLifecycleDetailResolved,
	"DISMISSED": RecommendationLifecycleDetailDismissed,
}

var mappingRecommendationLifecycleDetailEnumLowerCase = map[string]RecommendationLifecycleDetailEnum{
	"open":      RecommendationLifecycleDetailOpen,
	"resolved":  RecommendationLifecycleDetailResolved,
	"dismissed": RecommendationLifecycleDetailDismissed,
}

// GetRecommendationLifecycleDetailEnumValues Enumerates the set of values for RecommendationLifecycleDetailEnum
func GetRecommendationLifecycleDetailEnumValues() []RecommendationLifecycleDetailEnum {
	values := make([]RecommendationLifecycleDetailEnum, 0)
	for _, v := range mappingRecommendationLifecycleDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationLifecycleDetailEnumStringValues Enumerates the set of values in String for RecommendationLifecycleDetailEnum
func GetRecommendationLifecycleDetailEnumStringValues() []string {
	return []string{
		"OPEN",
		"RESOLVED",
		"DISMISSED",
	}
}

// GetMappingRecommendationLifecycleDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationLifecycleDetailEnum(val string) (RecommendationLifecycleDetailEnum, bool) {
	enum, ok := mappingRecommendationLifecycleDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
