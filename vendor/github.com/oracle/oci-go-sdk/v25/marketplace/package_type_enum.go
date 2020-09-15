// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// PackageTypeEnumEnum Enum with underlying type: string
type PackageTypeEnumEnum string

// Set of constants representing the allowable values for PackageTypeEnumEnum
const (
	PackageTypeEnumOrchestration PackageTypeEnumEnum = "ORCHESTRATION"
	PackageTypeEnumImage         PackageTypeEnumEnum = "IMAGE"
)

var mappingPackageTypeEnum = map[string]PackageTypeEnumEnum{
	"ORCHESTRATION": PackageTypeEnumOrchestration,
	"IMAGE":         PackageTypeEnumImage,
}

// GetPackageTypeEnumEnumValues Enumerates the set of values for PackageTypeEnumEnum
func GetPackageTypeEnumEnumValues() []PackageTypeEnumEnum {
	values := make([]PackageTypeEnumEnum, 0)
	for _, v := range mappingPackageTypeEnum {
		values = append(values, v)
	}
	return values
}
