// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// ObjectCollectionRuleCollectionTypesEnum Enum with underlying type: string
type ObjectCollectionRuleCollectionTypesEnum string

// Set of constants representing the allowable values for ObjectCollectionRuleCollectionTypesEnum
const (
	ObjectCollectionRuleCollectionTypesLive         ObjectCollectionRuleCollectionTypesEnum = "LIVE"
	ObjectCollectionRuleCollectionTypesHistoric     ObjectCollectionRuleCollectionTypesEnum = "HISTORIC"
	ObjectCollectionRuleCollectionTypesHistoricLive ObjectCollectionRuleCollectionTypesEnum = "HISTORIC_LIVE"
)

var mappingObjectCollectionRuleCollectionTypes = map[string]ObjectCollectionRuleCollectionTypesEnum{
	"LIVE":          ObjectCollectionRuleCollectionTypesLive,
	"HISTORIC":      ObjectCollectionRuleCollectionTypesHistoric,
	"HISTORIC_LIVE": ObjectCollectionRuleCollectionTypesHistoricLive,
}

// GetObjectCollectionRuleCollectionTypesEnumValues Enumerates the set of values for ObjectCollectionRuleCollectionTypesEnum
func GetObjectCollectionRuleCollectionTypesEnumValues() []ObjectCollectionRuleCollectionTypesEnum {
	values := make([]ObjectCollectionRuleCollectionTypesEnum, 0)
	for _, v := range mappingObjectCollectionRuleCollectionTypes {
		values = append(values, v)
	}
	return values
}
