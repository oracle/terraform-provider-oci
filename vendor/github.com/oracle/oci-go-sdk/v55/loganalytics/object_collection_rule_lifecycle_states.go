// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// ObjectCollectionRuleLifecycleStatesEnum Enum with underlying type: string
type ObjectCollectionRuleLifecycleStatesEnum string

// Set of constants representing the allowable values for ObjectCollectionRuleLifecycleStatesEnum
const (
	ObjectCollectionRuleLifecycleStatesActive   ObjectCollectionRuleLifecycleStatesEnum = "ACTIVE"
	ObjectCollectionRuleLifecycleStatesDeleted  ObjectCollectionRuleLifecycleStatesEnum = "DELETED"
	ObjectCollectionRuleLifecycleStatesInactive ObjectCollectionRuleLifecycleStatesEnum = "INACTIVE"
)

var mappingObjectCollectionRuleLifecycleStates = map[string]ObjectCollectionRuleLifecycleStatesEnum{
	"ACTIVE":   ObjectCollectionRuleLifecycleStatesActive,
	"DELETED":  ObjectCollectionRuleLifecycleStatesDeleted,
	"INACTIVE": ObjectCollectionRuleLifecycleStatesInactive,
}

// GetObjectCollectionRuleLifecycleStatesEnumValues Enumerates the set of values for ObjectCollectionRuleLifecycleStatesEnum
func GetObjectCollectionRuleLifecycleStatesEnumValues() []ObjectCollectionRuleLifecycleStatesEnum {
	values := make([]ObjectCollectionRuleLifecycleStatesEnum, 0)
	for _, v := range mappingObjectCollectionRuleLifecycleStates {
		values = append(values, v)
	}
	return values
}
