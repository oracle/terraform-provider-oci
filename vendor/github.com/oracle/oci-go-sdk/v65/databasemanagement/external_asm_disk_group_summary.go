// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalAsmDiskGroupSummary The summary of an external ASM disk group.
type ExternalAsmDiskGroupSummary struct {

	// The name of the ASM disk group.
	Name *string `mandatory:"true" json:"name"`

	// The number of ASM instances that have the disk group in mounted state.
	MountingInstanceCount *int `mandatory:"false" json:"mountingInstanceCount"`

	// The number of ASM instances that have the disk group in dismounted state.
	DismountingInstanceCount *int `mandatory:"false" json:"dismountingInstanceCount"`

	// The redundancy type of the disk group.
	RedundancyType ExternalAsmDiskGroupSummaryRedundancyTypeEnum `mandatory:"false" json:"redundancyType,omitempty"`

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

func (m ExternalAsmDiskGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalAsmDiskGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExternalAsmDiskGroupSummaryRedundancyTypeEnum(string(m.RedundancyType)); !ok && m.RedundancyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RedundancyType: %s. Supported values are: %s.", m.RedundancyType, strings.Join(GetExternalAsmDiskGroupSummaryRedundancyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalAsmDiskGroupSummaryRedundancyTypeEnum Enum with underlying type: string
type ExternalAsmDiskGroupSummaryRedundancyTypeEnum string

// Set of constants representing the allowable values for ExternalAsmDiskGroupSummaryRedundancyTypeEnum
const (
	ExternalAsmDiskGroupSummaryRedundancyTypeExtend ExternalAsmDiskGroupSummaryRedundancyTypeEnum = "EXTEND"
	ExternalAsmDiskGroupSummaryRedundancyTypeExtern ExternalAsmDiskGroupSummaryRedundancyTypeEnum = "EXTERN"
	ExternalAsmDiskGroupSummaryRedundancyTypeFlex   ExternalAsmDiskGroupSummaryRedundancyTypeEnum = "FLEX"
	ExternalAsmDiskGroupSummaryRedundancyTypeHigh   ExternalAsmDiskGroupSummaryRedundancyTypeEnum = "HIGH"
	ExternalAsmDiskGroupSummaryRedundancyTypeNormal ExternalAsmDiskGroupSummaryRedundancyTypeEnum = "NORMAL"
)

var mappingExternalAsmDiskGroupSummaryRedundancyTypeEnum = map[string]ExternalAsmDiskGroupSummaryRedundancyTypeEnum{
	"EXTEND": ExternalAsmDiskGroupSummaryRedundancyTypeExtend,
	"EXTERN": ExternalAsmDiskGroupSummaryRedundancyTypeExtern,
	"FLEX":   ExternalAsmDiskGroupSummaryRedundancyTypeFlex,
	"HIGH":   ExternalAsmDiskGroupSummaryRedundancyTypeHigh,
	"NORMAL": ExternalAsmDiskGroupSummaryRedundancyTypeNormal,
}

var mappingExternalAsmDiskGroupSummaryRedundancyTypeEnumLowerCase = map[string]ExternalAsmDiskGroupSummaryRedundancyTypeEnum{
	"extend": ExternalAsmDiskGroupSummaryRedundancyTypeExtend,
	"extern": ExternalAsmDiskGroupSummaryRedundancyTypeExtern,
	"flex":   ExternalAsmDiskGroupSummaryRedundancyTypeFlex,
	"high":   ExternalAsmDiskGroupSummaryRedundancyTypeHigh,
	"normal": ExternalAsmDiskGroupSummaryRedundancyTypeNormal,
}

// GetExternalAsmDiskGroupSummaryRedundancyTypeEnumValues Enumerates the set of values for ExternalAsmDiskGroupSummaryRedundancyTypeEnum
func GetExternalAsmDiskGroupSummaryRedundancyTypeEnumValues() []ExternalAsmDiskGroupSummaryRedundancyTypeEnum {
	values := make([]ExternalAsmDiskGroupSummaryRedundancyTypeEnum, 0)
	for _, v := range mappingExternalAsmDiskGroupSummaryRedundancyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalAsmDiskGroupSummaryRedundancyTypeEnumStringValues Enumerates the set of values in String for ExternalAsmDiskGroupSummaryRedundancyTypeEnum
func GetExternalAsmDiskGroupSummaryRedundancyTypeEnumStringValues() []string {
	return []string{
		"EXTEND",
		"EXTERN",
		"FLEX",
		"HIGH",
		"NORMAL",
	}
}

// GetMappingExternalAsmDiskGroupSummaryRedundancyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalAsmDiskGroupSummaryRedundancyTypeEnum(val string) (ExternalAsmDiskGroupSummaryRedundancyTypeEnum, bool) {
	enum, ok := mappingExternalAsmDiskGroupSummaryRedundancyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
