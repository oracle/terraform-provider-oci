// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

// DedicatedVantagePointStatusEnum Enum with underlying type: string
type DedicatedVantagePointStatusEnum string

// Set of constants representing the allowable values for DedicatedVantagePointStatusEnum
const (
	DedicatedVantagePointStatusEnabled  DedicatedVantagePointStatusEnum = "ENABLED"
	DedicatedVantagePointStatusDisabled DedicatedVantagePointStatusEnum = "DISABLED"
)

var mappingDedicatedVantagePointStatusEnum = map[string]DedicatedVantagePointStatusEnum{
	"ENABLED":  DedicatedVantagePointStatusEnabled,
	"DISABLED": DedicatedVantagePointStatusDisabled,
}

// GetDedicatedVantagePointStatusEnumValues Enumerates the set of values for DedicatedVantagePointStatusEnum
func GetDedicatedVantagePointStatusEnumValues() []DedicatedVantagePointStatusEnum {
	values := make([]DedicatedVantagePointStatusEnum, 0)
	for _, v := range mappingDedicatedVantagePointStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedVantagePointStatusEnumStringValues Enumerates the set of values in String for DedicatedVantagePointStatusEnum
func GetDedicatedVantagePointStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}
