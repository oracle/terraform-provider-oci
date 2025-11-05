// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// ServiceEnum Enum with underlying type: string
type ServiceEnum string

// Set of constants representing the allowable values for ServiceEnum
const (
	ServiceAdb           ServiceEnum = "ADB"
	ServiceGgcs          ServiceEnum = "GGCS"
	ServiceObjectstorage ServiceEnum = "OBJECTSTORAGE"
	ServiceGenai         ServiceEnum = "GENAI"
	ServiceDataflow      ServiceEnum = "DATAFLOW"
)

var mappingServiceEnum = map[string]ServiceEnum{
	"ADB":           ServiceAdb,
	"GGCS":          ServiceGgcs,
	"OBJECTSTORAGE": ServiceObjectstorage,
	"GENAI":         ServiceGenai,
	"DATAFLOW":      ServiceDataflow,
}

var mappingServiceEnumLowerCase = map[string]ServiceEnum{
	"adb":           ServiceAdb,
	"ggcs":          ServiceGgcs,
	"objectstorage": ServiceObjectstorage,
	"genai":         ServiceGenai,
	"dataflow":      ServiceDataflow,
}

// GetServiceEnumValues Enumerates the set of values for ServiceEnum
func GetServiceEnumValues() []ServiceEnum {
	values := make([]ServiceEnum, 0)
	for _, v := range mappingServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceEnumStringValues Enumerates the set of values in String for ServiceEnum
func GetServiceEnumStringValues() []string {
	return []string{
		"ADB",
		"GGCS",
		"OBJECTSTORAGE",
		"GENAI",
		"DATAFLOW",
	}
}

// GetMappingServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceEnum(val string) (ServiceEnum, bool) {
	enum, ok := mappingServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
