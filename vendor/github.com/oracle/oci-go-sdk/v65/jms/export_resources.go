// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// ExportResourcesEnum Enum with underlying type: string
type ExportResourcesEnum string

// Set of constants representing the allowable values for ExportResourcesEnum
const (
	ExportResourcesManagedInstance                               ExportResourcesEnum = "MANAGED_INSTANCE"
	ExportResourcesManagedInstancePlusJavaRuntime                ExportResourcesEnum = "MANAGED_INSTANCE_PLUS_JAVA_RUNTIME"
	ExportResourcesManagedInstancePlusJavaRuntimePlusApplication ExportResourcesEnum = "MANAGED_INSTANCE_PLUS_JAVA_RUNTIME_PLUS_APPLICATION"
)

var mappingExportResourcesEnum = map[string]ExportResourcesEnum{
	"MANAGED_INSTANCE":                                    ExportResourcesManagedInstance,
	"MANAGED_INSTANCE_PLUS_JAVA_RUNTIME":                  ExportResourcesManagedInstancePlusJavaRuntime,
	"MANAGED_INSTANCE_PLUS_JAVA_RUNTIME_PLUS_APPLICATION": ExportResourcesManagedInstancePlusJavaRuntimePlusApplication,
}

var mappingExportResourcesEnumLowerCase = map[string]ExportResourcesEnum{
	"managed_instance":                                    ExportResourcesManagedInstance,
	"managed_instance_plus_java_runtime":                  ExportResourcesManagedInstancePlusJavaRuntime,
	"managed_instance_plus_java_runtime_plus_application": ExportResourcesManagedInstancePlusJavaRuntimePlusApplication,
}

// GetExportResourcesEnumValues Enumerates the set of values for ExportResourcesEnum
func GetExportResourcesEnumValues() []ExportResourcesEnum {
	values := make([]ExportResourcesEnum, 0)
	for _, v := range mappingExportResourcesEnum {
		values = append(values, v)
	}
	return values
}

// GetExportResourcesEnumStringValues Enumerates the set of values in String for ExportResourcesEnum
func GetExportResourcesEnumStringValues() []string {
	return []string{
		"MANAGED_INSTANCE",
		"MANAGED_INSTANCE_PLUS_JAVA_RUNTIME",
		"MANAGED_INSTANCE_PLUS_JAVA_RUNTIME_PLUS_APPLICATION",
	}
}

// GetMappingExportResourcesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportResourcesEnum(val string) (ExportResourcesEnum, bool) {
	enum, ok := mappingExportResourcesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
