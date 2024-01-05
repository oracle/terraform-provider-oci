// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// PackageTypeEnumEnum Enum with underlying type: string
type PackageTypeEnumEnum string

// Set of constants representing the allowable values for PackageTypeEnumEnum
const (
	PackageTypeEnumOrchestration PackageTypeEnumEnum = "ORCHESTRATION"
	PackageTypeEnumImage         PackageTypeEnumEnum = "IMAGE"
	PackageTypeEnumContainer     PackageTypeEnumEnum = "CONTAINER"
	PackageTypeEnumKubernetes    PackageTypeEnumEnum = "KUBERNETES"
)

var mappingPackageTypeEnumEnum = map[string]PackageTypeEnumEnum{
	"ORCHESTRATION": PackageTypeEnumOrchestration,
	"IMAGE":         PackageTypeEnumImage,
	"CONTAINER":     PackageTypeEnumContainer,
	"KUBERNETES":    PackageTypeEnumKubernetes,
}

var mappingPackageTypeEnumEnumLowerCase = map[string]PackageTypeEnumEnum{
	"orchestration": PackageTypeEnumOrchestration,
	"image":         PackageTypeEnumImage,
	"container":     PackageTypeEnumContainer,
	"kubernetes":    PackageTypeEnumKubernetes,
}

// GetPackageTypeEnumEnumValues Enumerates the set of values for PackageTypeEnumEnum
func GetPackageTypeEnumEnumValues() []PackageTypeEnumEnum {
	values := make([]PackageTypeEnumEnum, 0)
	for _, v := range mappingPackageTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageTypeEnumEnumStringValues Enumerates the set of values in String for PackageTypeEnumEnum
func GetPackageTypeEnumEnumStringValues() []string {
	return []string{
		"ORCHESTRATION",
		"IMAGE",
		"CONTAINER",
		"KUBERNETES",
	}
}

// GetMappingPackageTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnumEnum(val string) (PackageTypeEnumEnum, bool) {
	enum, ok := mappingPackageTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
