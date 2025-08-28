// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OsPatchPackageSummary Summary of a package contained in a os patch.
type OsPatchPackageSummary struct {

	// The package's name.
	PackageName *string `mandatory:"true" json:"packageName"`

	// The action that current package will be executed on the cluster.
	UpdateType OsPatchPackageSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// Related CVEs of the package update.
	RelatedCVEs []string `mandatory:"true" json:"relatedCVEs"`

	// The target version of the package.
	TargetVersion *string `mandatory:"false" json:"targetVersion"`

	// Package type based on package installation manager.
	PackageType OsPatchPackageSummaryPackageTypeEnum `mandatory:"false" json:"packageType,omitempty"`
}

func (m OsPatchPackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OsPatchPackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOsPatchPackageSummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetOsPatchPackageSummaryUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOsPatchPackageSummaryPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetOsPatchPackageSummaryPackageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OsPatchPackageSummaryUpdateTypeEnum Enum with underlying type: string
type OsPatchPackageSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for OsPatchPackageSummaryUpdateTypeEnum
const (
	OsPatchPackageSummaryUpdateTypeInstall OsPatchPackageSummaryUpdateTypeEnum = "INSTALL"
	OsPatchPackageSummaryUpdateTypeRemove  OsPatchPackageSummaryUpdateTypeEnum = "REMOVE"
	OsPatchPackageSummaryUpdateTypeUpdate  OsPatchPackageSummaryUpdateTypeEnum = "UPDATE"
)

var mappingOsPatchPackageSummaryUpdateTypeEnum = map[string]OsPatchPackageSummaryUpdateTypeEnum{
	"INSTALL": OsPatchPackageSummaryUpdateTypeInstall,
	"REMOVE":  OsPatchPackageSummaryUpdateTypeRemove,
	"UPDATE":  OsPatchPackageSummaryUpdateTypeUpdate,
}

var mappingOsPatchPackageSummaryUpdateTypeEnumLowerCase = map[string]OsPatchPackageSummaryUpdateTypeEnum{
	"install": OsPatchPackageSummaryUpdateTypeInstall,
	"remove":  OsPatchPackageSummaryUpdateTypeRemove,
	"update":  OsPatchPackageSummaryUpdateTypeUpdate,
}

// GetOsPatchPackageSummaryUpdateTypeEnumValues Enumerates the set of values for OsPatchPackageSummaryUpdateTypeEnum
func GetOsPatchPackageSummaryUpdateTypeEnumValues() []OsPatchPackageSummaryUpdateTypeEnum {
	values := make([]OsPatchPackageSummaryUpdateTypeEnum, 0)
	for _, v := range mappingOsPatchPackageSummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOsPatchPackageSummaryUpdateTypeEnumStringValues Enumerates the set of values in String for OsPatchPackageSummaryUpdateTypeEnum
func GetOsPatchPackageSummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"INSTALL",
		"REMOVE",
		"UPDATE",
	}
}

// GetMappingOsPatchPackageSummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsPatchPackageSummaryUpdateTypeEnum(val string) (OsPatchPackageSummaryUpdateTypeEnum, bool) {
	enum, ok := mappingOsPatchPackageSummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OsPatchPackageSummaryPackageTypeEnum Enum with underlying type: string
type OsPatchPackageSummaryPackageTypeEnum string

// Set of constants representing the allowable values for OsPatchPackageSummaryPackageTypeEnum
const (
	OsPatchPackageSummaryPackageTypeRpm   OsPatchPackageSummaryPackageTypeEnum = "RPM"
	OsPatchPackageSummaryPackageTypePip   OsPatchPackageSummaryPackageTypeEnum = "PIP"
	OsPatchPackageSummaryPackageTypePip3  OsPatchPackageSummaryPackageTypeEnum = "PIP3"
	OsPatchPackageSummaryPackageTypePip38 OsPatchPackageSummaryPackageTypeEnum = "PIP3_8"
)

var mappingOsPatchPackageSummaryPackageTypeEnum = map[string]OsPatchPackageSummaryPackageTypeEnum{
	"RPM":    OsPatchPackageSummaryPackageTypeRpm,
	"PIP":    OsPatchPackageSummaryPackageTypePip,
	"PIP3":   OsPatchPackageSummaryPackageTypePip3,
	"PIP3_8": OsPatchPackageSummaryPackageTypePip38,
}

var mappingOsPatchPackageSummaryPackageTypeEnumLowerCase = map[string]OsPatchPackageSummaryPackageTypeEnum{
	"rpm":    OsPatchPackageSummaryPackageTypeRpm,
	"pip":    OsPatchPackageSummaryPackageTypePip,
	"pip3":   OsPatchPackageSummaryPackageTypePip3,
	"pip3_8": OsPatchPackageSummaryPackageTypePip38,
}

// GetOsPatchPackageSummaryPackageTypeEnumValues Enumerates the set of values for OsPatchPackageSummaryPackageTypeEnum
func GetOsPatchPackageSummaryPackageTypeEnumValues() []OsPatchPackageSummaryPackageTypeEnum {
	values := make([]OsPatchPackageSummaryPackageTypeEnum, 0)
	for _, v := range mappingOsPatchPackageSummaryPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOsPatchPackageSummaryPackageTypeEnumStringValues Enumerates the set of values in String for OsPatchPackageSummaryPackageTypeEnum
func GetOsPatchPackageSummaryPackageTypeEnumStringValues() []string {
	return []string{
		"RPM",
		"PIP",
		"PIP3",
		"PIP3_8",
	}
}

// GetMappingOsPatchPackageSummaryPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsPatchPackageSummaryPackageTypeEnum(val string) (OsPatchPackageSummaryPackageTypeEnum, bool) {
	enum, ok := mappingOsPatchPackageSummaryPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
