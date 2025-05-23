// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PackageGroup Yum or DNF package group, category, or environment.
type PackageGroup struct {

	// Package group identifier.
	Id *string `mandatory:"true" json:"id"`

	// Package group name.
	Name *string `mandatory:"true" json:"name"`

	// The list of packages in the package group.
	Packages []string `mandatory:"true" json:"packages"`

	// The repository IDs of the package group's repositories.
	Repositories []string `mandatory:"false" json:"repositories"`

	// Description of the package group.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if this package group is visible to users.
	IsUserVisible *bool `mandatory:"false" json:"isUserVisible"`

	// Indicates if this package group is the default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Indicates if this is a group, category, or environment.
	GroupType PackageGroupGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`

	// Indicates the order to display category or environment.
	DisplayOrder *int `mandatory:"false" json:"displayOrder"`
}

func (m PackageGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PackageGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPackageGroupGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetPackageGroupGroupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PackageGroupGroupTypeEnum Enum with underlying type: string
type PackageGroupGroupTypeEnum string

// Set of constants representing the allowable values for PackageGroupGroupTypeEnum
const (
	PackageGroupGroupTypeGroup       PackageGroupGroupTypeEnum = "GROUP"
	PackageGroupGroupTypeEnvironment PackageGroupGroupTypeEnum = "ENVIRONMENT"
	PackageGroupGroupTypeCategory    PackageGroupGroupTypeEnum = "CATEGORY"
)

var mappingPackageGroupGroupTypeEnum = map[string]PackageGroupGroupTypeEnum{
	"GROUP":       PackageGroupGroupTypeGroup,
	"ENVIRONMENT": PackageGroupGroupTypeEnvironment,
	"CATEGORY":    PackageGroupGroupTypeCategory,
}

var mappingPackageGroupGroupTypeEnumLowerCase = map[string]PackageGroupGroupTypeEnum{
	"group":       PackageGroupGroupTypeGroup,
	"environment": PackageGroupGroupTypeEnvironment,
	"category":    PackageGroupGroupTypeCategory,
}

// GetPackageGroupGroupTypeEnumValues Enumerates the set of values for PackageGroupGroupTypeEnum
func GetPackageGroupGroupTypeEnumValues() []PackageGroupGroupTypeEnum {
	values := make([]PackageGroupGroupTypeEnum, 0)
	for _, v := range mappingPackageGroupGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageGroupGroupTypeEnumStringValues Enumerates the set of values in String for PackageGroupGroupTypeEnum
func GetPackageGroupGroupTypeEnumStringValues() []string {
	return []string{
		"GROUP",
		"ENVIRONMENT",
		"CATEGORY",
	}
}

// GetMappingPackageGroupGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageGroupGroupTypeEnum(val string) (PackageGroupGroupTypeEnum, bool) {
	enum, ok := mappingPackageGroupGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
