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

// DbSystemShapeSummary The shape of the DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type DbSystemShapeSummary struct {

	// The name of the shape used for the DB system.
	Name *string `mandatory:"true" json:"name"`

	// The maximum number of CPU cores that can be enabled on the DB system for this shape.
	AvailableCoreCount *int `mandatory:"true" json:"availableCoreCount"`

	// The family of the shape used for the DB system.
	ShapeFamily *string `mandatory:"false" json:"shapeFamily"`

	// The shape type for the virtual machine DB system. Shape type is determined by CPU hardware. Valid values are `AMD` , `INTEL`, `INTEL_FLEX_X9` or `AMPERE_FLEX_A1`.
	ShapeType DbSystemShapeSummaryShapeTypeEnum `mandatory:"false" json:"shapeType,omitempty"`

	// Deprecated. Use `name` instead of `shape`.
	Shape *string `mandatory:"false" json:"shape"`

	// The minimum number of CPU cores that can be enabled on the DB system for this shape.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The discrete number by which the CPU core count for this shape can be increased or decreased.
	CoreCountIncrement *int `mandatory:"false" json:"coreCountIncrement"`

	// The minimum number of Exadata storage servers available for the Exadata infrastructure.
	MinStorageCount *int `mandatory:"false" json:"minStorageCount"`

	// The maximum number of Exadata storage servers available for the Exadata infrastructure.
	MaxStorageCount *int `mandatory:"false" json:"maxStorageCount"`

	// The maximum data storage available per storage server for this shape. Only applicable to ExaCC Elastic shapes.
	AvailableDataStoragePerServerInTBs *float64 `mandatory:"false" json:"availableDataStoragePerServerInTBs"`

	// The maximum memory available per database node for this shape. Only applicable to ExaCC Elastic shapes.
	AvailableMemoryPerNodeInGBs *int `mandatory:"false" json:"availableMemoryPerNodeInGBs"`

	// The maximum Db Node storage available per database node for this shape. Only applicable to ExaCC Elastic shapes.
	AvailableDbNodePerNodeInGBs *int `mandatory:"false" json:"availableDbNodePerNodeInGBs"`

	// The minimum number of CPU cores that can be enabled per node for this shape.
	MinCoreCountPerNode *int `mandatory:"false" json:"minCoreCountPerNode"`

	// The maximum memory that can be enabled for this shape.
	AvailableMemoryInGBs *int `mandatory:"false" json:"availableMemoryInGBs"`

	// The minimum memory that need be allocated per node for this shape.
	MinMemoryPerNodeInGBs *int `mandatory:"false" json:"minMemoryPerNodeInGBs"`

	// The maximum Db Node storage that can be enabled for this shape.
	AvailableDbNodeStorageInGBs *int `mandatory:"false" json:"availableDbNodeStorageInGBs"`

	// The minimum Db Node storage that need be allocated per node for this shape.
	MinDbNodeStoragePerNodeInGBs *int `mandatory:"false" json:"minDbNodeStoragePerNodeInGBs"`

	// The maximum DATA storage that can be enabled for this shape.
	AvailableDataStorageInTBs *int `mandatory:"false" json:"availableDataStorageInTBs"`

	// The minimum data storage that need be allocated for this shape.
	MinDataStorageInTBs *int `mandatory:"false" json:"minDataStorageInTBs"`

	// The minimum number of compute servers available for this shape.
	MinimumNodeCount *int `mandatory:"false" json:"minimumNodeCount"`

	// The maximum number of compute servers available for this shape.
	MaximumNodeCount *int `mandatory:"false" json:"maximumNodeCount"`

	// The maximum number of CPU cores per database node that can be enabled for this shape. Only applicable to the flex Exadata shape, ExaCC Elastic shapes and VM Flex shapes.
	AvailableCoreCountPerNode *int `mandatory:"false" json:"availableCoreCountPerNode"`
}

func (m DbSystemShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbSystemShapeSummaryShapeTypeEnum(string(m.ShapeType)); !ok && m.ShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeType: %s. Supported values are: %s.", m.ShapeType, strings.Join(GetDbSystemShapeSummaryShapeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemShapeSummaryShapeTypeEnum Enum with underlying type: string
type DbSystemShapeSummaryShapeTypeEnum string

// Set of constants representing the allowable values for DbSystemShapeSummaryShapeTypeEnum
const (
	DbSystemShapeSummaryShapeTypeAmd          DbSystemShapeSummaryShapeTypeEnum = "AMD"
	DbSystemShapeSummaryShapeTypeIntel        DbSystemShapeSummaryShapeTypeEnum = "INTEL"
	DbSystemShapeSummaryShapeTypeIntelFlexX9  DbSystemShapeSummaryShapeTypeEnum = "INTEL_FLEX_X9"
	DbSystemShapeSummaryShapeTypeAmpereFlexA1 DbSystemShapeSummaryShapeTypeEnum = "AMPERE_FLEX_A1"
)

var mappingDbSystemShapeSummaryShapeTypeEnum = map[string]DbSystemShapeSummaryShapeTypeEnum{
	"AMD":            DbSystemShapeSummaryShapeTypeAmd,
	"INTEL":          DbSystemShapeSummaryShapeTypeIntel,
	"INTEL_FLEX_X9":  DbSystemShapeSummaryShapeTypeIntelFlexX9,
	"AMPERE_FLEX_A1": DbSystemShapeSummaryShapeTypeAmpereFlexA1,
}

var mappingDbSystemShapeSummaryShapeTypeEnumLowerCase = map[string]DbSystemShapeSummaryShapeTypeEnum{
	"amd":            DbSystemShapeSummaryShapeTypeAmd,
	"intel":          DbSystemShapeSummaryShapeTypeIntel,
	"intel_flex_x9":  DbSystemShapeSummaryShapeTypeIntelFlexX9,
	"ampere_flex_a1": DbSystemShapeSummaryShapeTypeAmpereFlexA1,
}

// GetDbSystemShapeSummaryShapeTypeEnumValues Enumerates the set of values for DbSystemShapeSummaryShapeTypeEnum
func GetDbSystemShapeSummaryShapeTypeEnumValues() []DbSystemShapeSummaryShapeTypeEnum {
	values := make([]DbSystemShapeSummaryShapeTypeEnum, 0)
	for _, v := range mappingDbSystemShapeSummaryShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemShapeSummaryShapeTypeEnumStringValues Enumerates the set of values in String for DbSystemShapeSummaryShapeTypeEnum
func GetDbSystemShapeSummaryShapeTypeEnumStringValues() []string {
	return []string{
		"AMD",
		"INTEL",
		"INTEL_FLEX_X9",
		"AMPERE_FLEX_A1",
	}
}

// GetMappingDbSystemShapeSummaryShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemShapeSummaryShapeTypeEnum(val string) (DbSystemShapeSummaryShapeTypeEnum, bool) {
	enum, ok := mappingDbSystemShapeSummaryShapeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
