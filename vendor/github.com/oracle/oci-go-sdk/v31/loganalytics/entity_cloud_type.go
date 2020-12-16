// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// EntityCloudTypeEnum Enum with underlying type: string
type EntityCloudTypeEnum string

// Set of constants representing the allowable values for EntityCloudTypeEnum
const (
	EntityCloudTypeCloud    EntityCloudTypeEnum = "CLOUD"
	EntityCloudTypeNonCloud EntityCloudTypeEnum = "NON_CLOUD"
)

var mappingEntityCloudType = map[string]EntityCloudTypeEnum{
	"CLOUD":     EntityCloudTypeCloud,
	"NON_CLOUD": EntityCloudTypeNonCloud,
}

// GetEntityCloudTypeEnumValues Enumerates the set of values for EntityCloudTypeEnum
func GetEntityCloudTypeEnumValues() []EntityCloudTypeEnum {
	values := make([]EntityCloudTypeEnum, 0)
	for _, v := range mappingEntityCloudType {
		values = append(values, v)
	}
	return values
}
