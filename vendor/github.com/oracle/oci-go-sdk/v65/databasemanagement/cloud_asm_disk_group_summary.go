// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAsmDiskGroupSummary The summary of a cloud ASM disk group.
type CloudAsmDiskGroupSummary struct {

	// The name of the ASM disk group.
	Name *string `mandatory:"true" json:"name"`

	// The number of ASM instances that have the disk group in mounted state.
	MountingInstanceCount *int `mandatory:"false" json:"mountingInstanceCount"`

	// The number of ASM instances that have the disk group in dismounted state.
	DismountingInstanceCount *int `mandatory:"false" json:"dismountingInstanceCount"`

	// The redundancy type of the disk group.
	RedundancyType CloudAsmDiskGroupSummaryRedundancyTypeEnum `mandatory:"false" json:"redundancyType,omitempty"`

	// Indicates whether the disk group is a sparse disk group or not.
	IsSparse *bool `mandatory:"false" json:"isSparse"`

	// The unique names of the databases using the disk group.
	Databases []string `mandatory:"false" json:"databases"`

	// The total capacity of the disk group (in megabytes).
	TotalSizeInMBs *int64 `mandatory:"false" json:"totalSizeInMBs"`

	// The used capacity of the disk group (in megabytes).
	UsedSizeInMBs *int64 `mandatory:"false" json:"usedSizeInMBs"`

	// The percentage of used space in the disk group.
	UsedPercent *float32 `mandatory:"false" json:"usedPercent"`
}

func (m CloudAsmDiskGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudAsmDiskGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCloudAsmDiskGroupSummaryRedundancyTypeEnum(string(m.RedundancyType)); !ok && m.RedundancyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RedundancyType: %s. Supported values are: %s.", m.RedundancyType, strings.Join(GetCloudAsmDiskGroupSummaryRedundancyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudAsmDiskGroupSummaryRedundancyTypeEnum Enum with underlying type: string
type CloudAsmDiskGroupSummaryRedundancyTypeEnum string

// Set of constants representing the allowable values for CloudAsmDiskGroupSummaryRedundancyTypeEnum
const (
	CloudAsmDiskGroupSummaryRedundancyTypeExtend CloudAsmDiskGroupSummaryRedundancyTypeEnum = "EXTEND"
	CloudAsmDiskGroupSummaryRedundancyTypeExtern CloudAsmDiskGroupSummaryRedundancyTypeEnum = "EXTERN"
	CloudAsmDiskGroupSummaryRedundancyTypeFlex   CloudAsmDiskGroupSummaryRedundancyTypeEnum = "FLEX"
	CloudAsmDiskGroupSummaryRedundancyTypeHigh   CloudAsmDiskGroupSummaryRedundancyTypeEnum = "HIGH"
	CloudAsmDiskGroupSummaryRedundancyTypeNormal CloudAsmDiskGroupSummaryRedundancyTypeEnum = "NORMAL"
)

var mappingCloudAsmDiskGroupSummaryRedundancyTypeEnum = map[string]CloudAsmDiskGroupSummaryRedundancyTypeEnum{
	"EXTEND": CloudAsmDiskGroupSummaryRedundancyTypeExtend,
	"EXTERN": CloudAsmDiskGroupSummaryRedundancyTypeExtern,
	"FLEX":   CloudAsmDiskGroupSummaryRedundancyTypeFlex,
	"HIGH":   CloudAsmDiskGroupSummaryRedundancyTypeHigh,
	"NORMAL": CloudAsmDiskGroupSummaryRedundancyTypeNormal,
}

var mappingCloudAsmDiskGroupSummaryRedundancyTypeEnumLowerCase = map[string]CloudAsmDiskGroupSummaryRedundancyTypeEnum{
	"extend": CloudAsmDiskGroupSummaryRedundancyTypeExtend,
	"extern": CloudAsmDiskGroupSummaryRedundancyTypeExtern,
	"flex":   CloudAsmDiskGroupSummaryRedundancyTypeFlex,
	"high":   CloudAsmDiskGroupSummaryRedundancyTypeHigh,
	"normal": CloudAsmDiskGroupSummaryRedundancyTypeNormal,
}

// GetCloudAsmDiskGroupSummaryRedundancyTypeEnumValues Enumerates the set of values for CloudAsmDiskGroupSummaryRedundancyTypeEnum
func GetCloudAsmDiskGroupSummaryRedundancyTypeEnumValues() []CloudAsmDiskGroupSummaryRedundancyTypeEnum {
	values := make([]CloudAsmDiskGroupSummaryRedundancyTypeEnum, 0)
	for _, v := range mappingCloudAsmDiskGroupSummaryRedundancyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAsmDiskGroupSummaryRedundancyTypeEnumStringValues Enumerates the set of values in String for CloudAsmDiskGroupSummaryRedundancyTypeEnum
func GetCloudAsmDiskGroupSummaryRedundancyTypeEnumStringValues() []string {
	return []string{
		"EXTEND",
		"EXTERN",
		"FLEX",
		"HIGH",
		"NORMAL",
	}
}

// GetMappingCloudAsmDiskGroupSummaryRedundancyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAsmDiskGroupSummaryRedundancyTypeEnum(val string) (CloudAsmDiskGroupSummaryRedundancyTypeEnum, bool) {
	enum, ok := mappingCloudAsmDiskGroupSummaryRedundancyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
