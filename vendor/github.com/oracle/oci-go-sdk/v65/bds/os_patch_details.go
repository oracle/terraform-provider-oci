// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// OsPatchDetails Details of an os patch.
type OsPatchDetails struct {

	// Version of the os patch.
	OsPatchVersion *string `mandatory:"true" json:"osPatchVersion"`

	// Minimum BDS version required to install current OS patch.
	MinBdsVersion *string `mandatory:"true" json:"minBdsVersion"`

	// Map of major ODH version to minimum ODH version required to install current OS patch. e.g. {ODH0.9: 0.9.1}
	MinCompatibleOdhVersionMap map[string]string `mandatory:"true" json:"minCompatibleOdhVersionMap"`

	// List of summaries of individual target packages.
	TargetPackages []OsPatchPackageSummary `mandatory:"true" json:"targetPackages"`

	// Released date of the OS patch.
	ReleaseDate *common.SDKTime `mandatory:"true" json:"releaseDate"`

	// Type of a specific os patch.
	// REGULAR means standard released os patches.
	// CUSTOM means os patches with some customizations.
	// EMERGENT means os patches with some emergency fixes that should be prioritized.
	PatchType OsPatchDetailsPatchTypeEnum `mandatory:"true" json:"patchType"`
}

func (m OsPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OsPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOsPatchDetailsPatchTypeEnum(string(m.PatchType)); !ok && m.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", m.PatchType, strings.Join(GetOsPatchDetailsPatchTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OsPatchDetailsPatchTypeEnum Enum with underlying type: string
type OsPatchDetailsPatchTypeEnum string

// Set of constants representing the allowable values for OsPatchDetailsPatchTypeEnum
const (
	OsPatchDetailsPatchTypeRegular  OsPatchDetailsPatchTypeEnum = "REGULAR"
	OsPatchDetailsPatchTypeCustom   OsPatchDetailsPatchTypeEnum = "CUSTOM"
	OsPatchDetailsPatchTypeEmergent OsPatchDetailsPatchTypeEnum = "EMERGENT"
)

var mappingOsPatchDetailsPatchTypeEnum = map[string]OsPatchDetailsPatchTypeEnum{
	"REGULAR":  OsPatchDetailsPatchTypeRegular,
	"CUSTOM":   OsPatchDetailsPatchTypeCustom,
	"EMERGENT": OsPatchDetailsPatchTypeEmergent,
}

var mappingOsPatchDetailsPatchTypeEnumLowerCase = map[string]OsPatchDetailsPatchTypeEnum{
	"regular":  OsPatchDetailsPatchTypeRegular,
	"custom":   OsPatchDetailsPatchTypeCustom,
	"emergent": OsPatchDetailsPatchTypeEmergent,
}

// GetOsPatchDetailsPatchTypeEnumValues Enumerates the set of values for OsPatchDetailsPatchTypeEnum
func GetOsPatchDetailsPatchTypeEnumValues() []OsPatchDetailsPatchTypeEnum {
	values := make([]OsPatchDetailsPatchTypeEnum, 0)
	for _, v := range mappingOsPatchDetailsPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOsPatchDetailsPatchTypeEnumStringValues Enumerates the set of values in String for OsPatchDetailsPatchTypeEnum
func GetOsPatchDetailsPatchTypeEnumStringValues() []string {
	return []string{
		"REGULAR",
		"CUSTOM",
		"EMERGENT",
	}
}

// GetMappingOsPatchDetailsPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsPatchDetailsPatchTypeEnum(val string) (OsPatchDetailsPatchTypeEnum, bool) {
	enum, ok := mappingOsPatchDetailsPatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
