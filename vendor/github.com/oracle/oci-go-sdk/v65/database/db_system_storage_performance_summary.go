// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystemStoragePerformanceSummary Representation of storage performance summary per shapeType .
type DbSystemStoragePerformanceSummary struct {

	// ShapeType of the DbSystems INTEL , AMD, INTEL_FLEX_X9 or AMPERE_FLEX_A1
	ShapeType DbSystemStoragePerformanceSummaryShapeTypeEnum `mandatory:"true" json:"shapeType"`

	// List of storage performance for the DATA disks
	DataStoragePerformanceList []StoragePerformanceDetails `mandatory:"true" json:"dataStoragePerformanceList"`

	// List of storage performance for the RECO disks
	RecoStoragePerformanceList []StoragePerformanceDetails `mandatory:"true" json:"recoStoragePerformanceList"`
}

func (m DbSystemStoragePerformanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemStoragePerformanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemStoragePerformanceSummaryShapeTypeEnum(string(m.ShapeType)); !ok && m.ShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeType: %s. Supported values are: %s.", m.ShapeType, strings.Join(GetDbSystemStoragePerformanceSummaryShapeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemStoragePerformanceSummaryShapeTypeEnum Enum with underlying type: string
type DbSystemStoragePerformanceSummaryShapeTypeEnum string

// Set of constants representing the allowable values for DbSystemStoragePerformanceSummaryShapeTypeEnum
const (
	DbSystemStoragePerformanceSummaryShapeTypeAmd          DbSystemStoragePerformanceSummaryShapeTypeEnum = "AMD"
	DbSystemStoragePerformanceSummaryShapeTypeIntel        DbSystemStoragePerformanceSummaryShapeTypeEnum = "INTEL"
	DbSystemStoragePerformanceSummaryShapeTypeIntelFlexX9  DbSystemStoragePerformanceSummaryShapeTypeEnum = "INTEL_FLEX_X9"
	DbSystemStoragePerformanceSummaryShapeTypeAmpereFlexA1 DbSystemStoragePerformanceSummaryShapeTypeEnum = "AMPERE_FLEX_A1"
)

var mappingDbSystemStoragePerformanceSummaryShapeTypeEnum = map[string]DbSystemStoragePerformanceSummaryShapeTypeEnum{
	"AMD":            DbSystemStoragePerformanceSummaryShapeTypeAmd,
	"INTEL":          DbSystemStoragePerformanceSummaryShapeTypeIntel,
	"INTEL_FLEX_X9":  DbSystemStoragePerformanceSummaryShapeTypeIntelFlexX9,
	"AMPERE_FLEX_A1": DbSystemStoragePerformanceSummaryShapeTypeAmpereFlexA1,
}

var mappingDbSystemStoragePerformanceSummaryShapeTypeEnumLowerCase = map[string]DbSystemStoragePerformanceSummaryShapeTypeEnum{
	"amd":            DbSystemStoragePerformanceSummaryShapeTypeAmd,
	"intel":          DbSystemStoragePerformanceSummaryShapeTypeIntel,
	"intel_flex_x9":  DbSystemStoragePerformanceSummaryShapeTypeIntelFlexX9,
	"ampere_flex_a1": DbSystemStoragePerformanceSummaryShapeTypeAmpereFlexA1,
}

// GetDbSystemStoragePerformanceSummaryShapeTypeEnumValues Enumerates the set of values for DbSystemStoragePerformanceSummaryShapeTypeEnum
func GetDbSystemStoragePerformanceSummaryShapeTypeEnumValues() []DbSystemStoragePerformanceSummaryShapeTypeEnum {
	values := make([]DbSystemStoragePerformanceSummaryShapeTypeEnum, 0)
	for _, v := range mappingDbSystemStoragePerformanceSummaryShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemStoragePerformanceSummaryShapeTypeEnumStringValues Enumerates the set of values in String for DbSystemStoragePerformanceSummaryShapeTypeEnum
func GetDbSystemStoragePerformanceSummaryShapeTypeEnumStringValues() []string {
	return []string{
		"AMD",
		"INTEL",
		"INTEL_FLEX_X9",
		"AMPERE_FLEX_A1",
	}
}

// GetMappingDbSystemStoragePerformanceSummaryShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemStoragePerformanceSummaryShapeTypeEnum(val string) (DbSystemStoragePerformanceSummaryShapeTypeEnum, bool) {
	enum, ok := mappingDbSystemStoragePerformanceSummaryShapeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
