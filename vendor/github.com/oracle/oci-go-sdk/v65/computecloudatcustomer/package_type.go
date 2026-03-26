// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"strings"
)

// PackageTypeEnum Enum with underlying type: string
type PackageTypeEnum string

// Set of constants representing the allowable values for PackageTypeEnum
const (
	PackageTypeOrchestration PackageTypeEnum = "ORCHESTRATION"
	PackageTypeImage         PackageTypeEnum = "IMAGE"
	PackageTypeContainer     PackageTypeEnum = "CONTAINER"
	PackageTypeKubernetes    PackageTypeEnum = "KUBERNETES"
	PackageTypeSaas          PackageTypeEnum = "SAAS"
)

var mappingPackageTypeEnum = map[string]PackageTypeEnum{
	"ORCHESTRATION": PackageTypeOrchestration,
	"IMAGE":         PackageTypeImage,
	"CONTAINER":     PackageTypeContainer,
	"KUBERNETES":    PackageTypeKubernetes,
	"SAAS":          PackageTypeSaas,
}

var mappingPackageTypeEnumLowerCase = map[string]PackageTypeEnum{
	"orchestration": PackageTypeOrchestration,
	"image":         PackageTypeImage,
	"container":     PackageTypeContainer,
	"kubernetes":    PackageTypeKubernetes,
	"saas":          PackageTypeSaas,
}

// GetPackageTypeEnumValues Enumerates the set of values for PackageTypeEnum
func GetPackageTypeEnumValues() []PackageTypeEnum {
	values := make([]PackageTypeEnum, 0)
	for _, v := range mappingPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageTypeEnumStringValues Enumerates the set of values in String for PackageTypeEnum
func GetPackageTypeEnumStringValues() []string {
	return []string{
		"ORCHESTRATION",
		"IMAGE",
		"CONTAINER",
		"KUBERNETES",
		"SAAS",
	}
}

// GetMappingPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnum(val string) (PackageTypeEnum, bool) {
	enum, ok := mappingPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
