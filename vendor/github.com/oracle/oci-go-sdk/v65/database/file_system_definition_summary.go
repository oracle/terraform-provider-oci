// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FileSystemDefinitionSummary File system configuration details for Exadata VM cluster.
type FileSystemDefinitionSummary struct {

	// Details of the file system definitions of the Exadata VM cluster.
	DefinedFileSystemConfigurations []FileSystemConfigurationDefinition `mandatory:"true" json:"definedFileSystemConfigurations"`

	// Specifies shape parameter.
	Shape *string `mandatory:"false" json:"shape"`

	// Specifies availability domain.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The type of Exascale storage used for Exadata VM cluster. The default is SMART_STORAGE which supports Oracle Database 23ai and later
	ShapeAttribute FileSystemDefinitionSummaryShapeAttributeEnum `mandatory:"false" json:"shapeAttribute,omitempty"`

	// Maximum size of Vm file systems.
	MaximumVmSizeInGBs *int `mandatory:"false" json:"maximumVmSizeInGBs"`
}

func (m FileSystemDefinitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemDefinitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFileSystemDefinitionSummaryShapeAttributeEnum(string(m.ShapeAttribute)); !ok && m.ShapeAttribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeAttribute: %s. Supported values are: %s.", m.ShapeAttribute, strings.Join(GetFileSystemDefinitionSummaryShapeAttributeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSystemDefinitionSummaryShapeAttributeEnum Enum with underlying type: string
type FileSystemDefinitionSummaryShapeAttributeEnum string

// Set of constants representing the allowable values for FileSystemDefinitionSummaryShapeAttributeEnum
const (
	FileSystemDefinitionSummaryShapeAttributeSmartStorage FileSystemDefinitionSummaryShapeAttributeEnum = "SMART_STORAGE"
	FileSystemDefinitionSummaryShapeAttributeBlockStorage FileSystemDefinitionSummaryShapeAttributeEnum = "BLOCK_STORAGE"
)

var mappingFileSystemDefinitionSummaryShapeAttributeEnum = map[string]FileSystemDefinitionSummaryShapeAttributeEnum{
	"SMART_STORAGE": FileSystemDefinitionSummaryShapeAttributeSmartStorage,
	"BLOCK_STORAGE": FileSystemDefinitionSummaryShapeAttributeBlockStorage,
}

var mappingFileSystemDefinitionSummaryShapeAttributeEnumLowerCase = map[string]FileSystemDefinitionSummaryShapeAttributeEnum{
	"smart_storage": FileSystemDefinitionSummaryShapeAttributeSmartStorage,
	"block_storage": FileSystemDefinitionSummaryShapeAttributeBlockStorage,
}

// GetFileSystemDefinitionSummaryShapeAttributeEnumValues Enumerates the set of values for FileSystemDefinitionSummaryShapeAttributeEnum
func GetFileSystemDefinitionSummaryShapeAttributeEnumValues() []FileSystemDefinitionSummaryShapeAttributeEnum {
	values := make([]FileSystemDefinitionSummaryShapeAttributeEnum, 0)
	for _, v := range mappingFileSystemDefinitionSummaryShapeAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemDefinitionSummaryShapeAttributeEnumStringValues Enumerates the set of values in String for FileSystemDefinitionSummaryShapeAttributeEnum
func GetFileSystemDefinitionSummaryShapeAttributeEnumStringValues() []string {
	return []string{
		"SMART_STORAGE",
		"BLOCK_STORAGE",
	}
}

// GetMappingFileSystemDefinitionSummaryShapeAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemDefinitionSummaryShapeAttributeEnum(val string) (FileSystemDefinitionSummaryShapeAttributeEnum, bool) {
	enum, ok := mappingFileSystemDefinitionSummaryShapeAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
