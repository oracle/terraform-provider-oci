// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportConflictResolution Import Objects Conflict resolution.
type ImportConflictResolution struct {

	// Import Objects Conflict resolution Type (RETAIN/DUPLICATE/REPLACE).
	ImportConflictResolutionType ImportConflictResolutionImportConflictResolutionTypeEnum `mandatory:"true" json:"importConflictResolutionType"`

	// In case of DUPLICATE mode, prefix will be used to disambiguate the object.
	DuplicatePrefix *string `mandatory:"false" json:"duplicatePrefix"`

	// In case of DUPLICATE mode, suffix will be used to disambiguate the object.
	DuplicateSuffix *string `mandatory:"false" json:"duplicateSuffix"`
}

func (m ImportConflictResolution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportConflictResolution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportConflictResolutionImportConflictResolutionTypeEnum(string(m.ImportConflictResolutionType)); !ok && m.ImportConflictResolutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImportConflictResolutionType: %s. Supported values are: %s.", m.ImportConflictResolutionType, strings.Join(GetImportConflictResolutionImportConflictResolutionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportConflictResolutionImportConflictResolutionTypeEnum Enum with underlying type: string
type ImportConflictResolutionImportConflictResolutionTypeEnum string

// Set of constants representing the allowable values for ImportConflictResolutionImportConflictResolutionTypeEnum
const (
	ImportConflictResolutionImportConflictResolutionTypeDuplicate ImportConflictResolutionImportConflictResolutionTypeEnum = "DUPLICATE"
	ImportConflictResolutionImportConflictResolutionTypeReplace   ImportConflictResolutionImportConflictResolutionTypeEnum = "REPLACE"
	ImportConflictResolutionImportConflictResolutionTypeRetain    ImportConflictResolutionImportConflictResolutionTypeEnum = "RETAIN"
)

var mappingImportConflictResolutionImportConflictResolutionTypeEnum = map[string]ImportConflictResolutionImportConflictResolutionTypeEnum{
	"DUPLICATE": ImportConflictResolutionImportConflictResolutionTypeDuplicate,
	"REPLACE":   ImportConflictResolutionImportConflictResolutionTypeReplace,
	"RETAIN":    ImportConflictResolutionImportConflictResolutionTypeRetain,
}

var mappingImportConflictResolutionImportConflictResolutionTypeEnumLowerCase = map[string]ImportConflictResolutionImportConflictResolutionTypeEnum{
	"duplicate": ImportConflictResolutionImportConflictResolutionTypeDuplicate,
	"replace":   ImportConflictResolutionImportConflictResolutionTypeReplace,
	"retain":    ImportConflictResolutionImportConflictResolutionTypeRetain,
}

// GetImportConflictResolutionImportConflictResolutionTypeEnumValues Enumerates the set of values for ImportConflictResolutionImportConflictResolutionTypeEnum
func GetImportConflictResolutionImportConflictResolutionTypeEnumValues() []ImportConflictResolutionImportConflictResolutionTypeEnum {
	values := make([]ImportConflictResolutionImportConflictResolutionTypeEnum, 0)
	for _, v := range mappingImportConflictResolutionImportConflictResolutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetImportConflictResolutionImportConflictResolutionTypeEnumStringValues Enumerates the set of values in String for ImportConflictResolutionImportConflictResolutionTypeEnum
func GetImportConflictResolutionImportConflictResolutionTypeEnumStringValues() []string {
	return []string{
		"DUPLICATE",
		"REPLACE",
		"RETAIN",
	}
}

// GetMappingImportConflictResolutionImportConflictResolutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportConflictResolutionImportConflictResolutionTypeEnum(val string) (ImportConflictResolutionImportConflictResolutionTypeEnum, bool) {
	enum, ok := mappingImportConflictResolutionImportConflictResolutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
