// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

// LifecycleStatesEnum Enum with underlying type: string
type LifecycleStatesEnum string

// Set of constants representing the allowable values for LifecycleStatesEnum
const (
	LifecycleStatesActive LifecycleStatesEnum = "ACTIVE"
)

var mappingLifecycleStates = map[string]LifecycleStatesEnum{
	"ACTIVE": LifecycleStatesActive,
}

// GetLifecycleStatesEnumValues Enumerates the set of values for LifecycleStatesEnum
func GetLifecycleStatesEnumValues() []LifecycleStatesEnum {
	values := make([]LifecycleStatesEnum, 0)
	for _, v := range mappingLifecycleStates {
		values = append(values, v)
	}
	return values
}
