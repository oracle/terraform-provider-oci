// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// CreationSourceTypeEnum Enum with underlying type: string
type CreationSourceTypeEnum string

// Set of constants representing the allowable values for CreationSourceTypeEnum
const (
	CreationSourceTypeEmBridge            CreationSourceTypeEnum = "EM_BRIDGE"
	CreationSourceTypeServiceConnectorHub CreationSourceTypeEnum = "SERVICE_CONNECTOR_HUB"
	CreationSourceTypeNone                CreationSourceTypeEnum = "NONE"
)

var mappingCreationSourceType = map[string]CreationSourceTypeEnum{
	"EM_BRIDGE":             CreationSourceTypeEmBridge,
	"SERVICE_CONNECTOR_HUB": CreationSourceTypeServiceConnectorHub,
	"NONE":                  CreationSourceTypeNone,
}

// GetCreationSourceTypeEnumValues Enumerates the set of values for CreationSourceTypeEnum
func GetCreationSourceTypeEnumValues() []CreationSourceTypeEnum {
	values := make([]CreationSourceTypeEnum, 0)
	for _, v := range mappingCreationSourceType {
		values = append(values, v)
	}
	return values
}
