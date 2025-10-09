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

// FileSystemDefinitionDetails File system definition request details.
type FileSystemDefinitionDetails struct {

	// Specifies shape parameter.
	Shape *string `mandatory:"true" json:"shape"`

	// Specifies shape parameter.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The type of Exascale storage used for Exadata VM cluster. The default is SMART_STORAGE which supports Oracle Database 23ai and later
	ShapeAttribute FileSystemDefinitionDetailsShapeAttributeEnum `mandatory:"false" json:"shapeAttribute,omitempty"`
}

func (m FileSystemDefinitionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemDefinitionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFileSystemDefinitionDetailsShapeAttributeEnum(string(m.ShapeAttribute)); !ok && m.ShapeAttribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeAttribute: %s. Supported values are: %s.", m.ShapeAttribute, strings.Join(GetFileSystemDefinitionDetailsShapeAttributeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSystemDefinitionDetailsShapeAttributeEnum Enum with underlying type: string
type FileSystemDefinitionDetailsShapeAttributeEnum string

// Set of constants representing the allowable values for FileSystemDefinitionDetailsShapeAttributeEnum
const (
	FileSystemDefinitionDetailsShapeAttributeSmartStorage FileSystemDefinitionDetailsShapeAttributeEnum = "SMART_STORAGE"
	FileSystemDefinitionDetailsShapeAttributeBlockStorage FileSystemDefinitionDetailsShapeAttributeEnum = "BLOCK_STORAGE"
)

var mappingFileSystemDefinitionDetailsShapeAttributeEnum = map[string]FileSystemDefinitionDetailsShapeAttributeEnum{
	"SMART_STORAGE": FileSystemDefinitionDetailsShapeAttributeSmartStorage,
	"BLOCK_STORAGE": FileSystemDefinitionDetailsShapeAttributeBlockStorage,
}

var mappingFileSystemDefinitionDetailsShapeAttributeEnumLowerCase = map[string]FileSystemDefinitionDetailsShapeAttributeEnum{
	"smart_storage": FileSystemDefinitionDetailsShapeAttributeSmartStorage,
	"block_storage": FileSystemDefinitionDetailsShapeAttributeBlockStorage,
}

// GetFileSystemDefinitionDetailsShapeAttributeEnumValues Enumerates the set of values for FileSystemDefinitionDetailsShapeAttributeEnum
func GetFileSystemDefinitionDetailsShapeAttributeEnumValues() []FileSystemDefinitionDetailsShapeAttributeEnum {
	values := make([]FileSystemDefinitionDetailsShapeAttributeEnum, 0)
	for _, v := range mappingFileSystemDefinitionDetailsShapeAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemDefinitionDetailsShapeAttributeEnumStringValues Enumerates the set of values in String for FileSystemDefinitionDetailsShapeAttributeEnum
func GetFileSystemDefinitionDetailsShapeAttributeEnumStringValues() []string {
	return []string{
		"SMART_STORAGE",
		"BLOCK_STORAGE",
	}
}

// GetMappingFileSystemDefinitionDetailsShapeAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemDefinitionDetailsShapeAttributeEnum(val string) (FileSystemDefinitionDetailsShapeAttributeEnum, bool) {
	enum, ok := mappingFileSystemDefinitionDetailsShapeAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
