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

// FileSystemDefinition File system configuration details for Exadata VM cluster.
type FileSystemDefinition struct {

	// Details of the file system definitions of the Exadata VM cluster.
	DefinedFileSystemConfigurations []FileSystemConfigurationDefinition `mandatory:"true" json:"definedFileSystemConfigurations"`

	// Specifies shape parameter.
	Shape *string `mandatory:"false" json:"shape"`

	// Specifies availability domain.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The type of Exascale storage used for Exadata VM cluster. The default is SMART_STORAGE which supports Oracle Database 23ai and later
	ShapeAttribute FileSystemDefinitionShapeAttributeEnum `mandatory:"false" json:"shapeAttribute,omitempty"`

	// Maximum size of Vm file systems.
	MaximumVmSizeInGBs *int `mandatory:"false" json:"maximumVmSizeInGBs"`
}

func (m FileSystemDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFileSystemDefinitionShapeAttributeEnum(string(m.ShapeAttribute)); !ok && m.ShapeAttribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeAttribute: %s. Supported values are: %s.", m.ShapeAttribute, strings.Join(GetFileSystemDefinitionShapeAttributeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSystemDefinitionShapeAttributeEnum Enum with underlying type: string
type FileSystemDefinitionShapeAttributeEnum string

// Set of constants representing the allowable values for FileSystemDefinitionShapeAttributeEnum
const (
	FileSystemDefinitionShapeAttributeSmartStorage FileSystemDefinitionShapeAttributeEnum = "SMART_STORAGE"
	FileSystemDefinitionShapeAttributeBlockStorage FileSystemDefinitionShapeAttributeEnum = "BLOCK_STORAGE"
)

var mappingFileSystemDefinitionShapeAttributeEnum = map[string]FileSystemDefinitionShapeAttributeEnum{
	"SMART_STORAGE": FileSystemDefinitionShapeAttributeSmartStorage,
	"BLOCK_STORAGE": FileSystemDefinitionShapeAttributeBlockStorage,
}

var mappingFileSystemDefinitionShapeAttributeEnumLowerCase = map[string]FileSystemDefinitionShapeAttributeEnum{
	"smart_storage": FileSystemDefinitionShapeAttributeSmartStorage,
	"block_storage": FileSystemDefinitionShapeAttributeBlockStorage,
}

// GetFileSystemDefinitionShapeAttributeEnumValues Enumerates the set of values for FileSystemDefinitionShapeAttributeEnum
func GetFileSystemDefinitionShapeAttributeEnumValues() []FileSystemDefinitionShapeAttributeEnum {
	values := make([]FileSystemDefinitionShapeAttributeEnum, 0)
	for _, v := range mappingFileSystemDefinitionShapeAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemDefinitionShapeAttributeEnumStringValues Enumerates the set of values in String for FileSystemDefinitionShapeAttributeEnum
func GetFileSystemDefinitionShapeAttributeEnumStringValues() []string {
	return []string{
		"SMART_STORAGE",
		"BLOCK_STORAGE",
	}
}

// GetMappingFileSystemDefinitionShapeAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemDefinitionShapeAttributeEnum(val string) (FileSystemDefinitionShapeAttributeEnum, bool) {
	enum, ok := mappingFileSystemDefinitionShapeAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
