// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// FeatureSetEnum Enum with underlying type: string
type FeatureSetEnum string

// Set of constants representing the allowable values for FeatureSetEnum
const (
	FeatureSetSelfServiceAnalytics FeatureSetEnum = "SELF_SERVICE_ANALYTICS"
	FeatureSetEnterpriseAnalytics  FeatureSetEnum = "ENTERPRISE_ANALYTICS"
)

var mappingFeatureSet = map[string]FeatureSetEnum{
	"SELF_SERVICE_ANALYTICS": FeatureSetSelfServiceAnalytics,
	"ENTERPRISE_ANALYTICS":   FeatureSetEnterpriseAnalytics,
}

// GetFeatureSetEnumValues Enumerates the set of values for FeatureSetEnum
func GetFeatureSetEnumValues() []FeatureSetEnum {
	values := make([]FeatureSetEnum, 0)
	for _, v := range mappingFeatureSet {
		values = append(values, v)
	}
	return values
}
