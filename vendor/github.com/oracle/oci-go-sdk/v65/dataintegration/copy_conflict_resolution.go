// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CopyConflictResolution Copy Object Conflict resolution.
type CopyConflictResolution struct {

	// Copy Object Conflict Resolution Type (RETAIN/DUPLICATE/REPLACE).
	CopyConflictResolutionType CopyConflictResolutionCopyConflictResolutionTypeEnum `mandatory:"true" json:"copyConflictResolutionType"`

	// In case of DUPLICATE mode, this prefix will be used to disambiguate the object.
	DuplicatePrefix *string `mandatory:"false" json:"duplicatePrefix"`

	// In case of DUPLICATE mode, this suffix will be used to disambiguate the object.
	DuplicateSuffix *string `mandatory:"false" json:"duplicateSuffix"`
}

func (m CopyConflictResolution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CopyConflictResolution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCopyConflictResolutionCopyConflictResolutionTypeEnum(string(m.CopyConflictResolutionType)); !ok && m.CopyConflictResolutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CopyConflictResolutionType: %s. Supported values are: %s.", m.CopyConflictResolutionType, strings.Join(GetCopyConflictResolutionCopyConflictResolutionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CopyConflictResolutionCopyConflictResolutionTypeEnum Enum with underlying type: string
type CopyConflictResolutionCopyConflictResolutionTypeEnum string

// Set of constants representing the allowable values for CopyConflictResolutionCopyConflictResolutionTypeEnum
const (
	CopyConflictResolutionCopyConflictResolutionTypeRetain    CopyConflictResolutionCopyConflictResolutionTypeEnum = "RETAIN"
	CopyConflictResolutionCopyConflictResolutionTypeDuplicate CopyConflictResolutionCopyConflictResolutionTypeEnum = "DUPLICATE"
	CopyConflictResolutionCopyConflictResolutionTypeReplace   CopyConflictResolutionCopyConflictResolutionTypeEnum = "REPLACE"
)

var mappingCopyConflictResolutionCopyConflictResolutionTypeEnum = map[string]CopyConflictResolutionCopyConflictResolutionTypeEnum{
	"RETAIN":    CopyConflictResolutionCopyConflictResolutionTypeRetain,
	"DUPLICATE": CopyConflictResolutionCopyConflictResolutionTypeDuplicate,
	"REPLACE":   CopyConflictResolutionCopyConflictResolutionTypeReplace,
}

var mappingCopyConflictResolutionCopyConflictResolutionTypeEnumLowerCase = map[string]CopyConflictResolutionCopyConflictResolutionTypeEnum{
	"retain":    CopyConflictResolutionCopyConflictResolutionTypeRetain,
	"duplicate": CopyConflictResolutionCopyConflictResolutionTypeDuplicate,
	"replace":   CopyConflictResolutionCopyConflictResolutionTypeReplace,
}

// GetCopyConflictResolutionCopyConflictResolutionTypeEnumValues Enumerates the set of values for CopyConflictResolutionCopyConflictResolutionTypeEnum
func GetCopyConflictResolutionCopyConflictResolutionTypeEnumValues() []CopyConflictResolutionCopyConflictResolutionTypeEnum {
	values := make([]CopyConflictResolutionCopyConflictResolutionTypeEnum, 0)
	for _, v := range mappingCopyConflictResolutionCopyConflictResolutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCopyConflictResolutionCopyConflictResolutionTypeEnumStringValues Enumerates the set of values in String for CopyConflictResolutionCopyConflictResolutionTypeEnum
func GetCopyConflictResolutionCopyConflictResolutionTypeEnumStringValues() []string {
	return []string{
		"RETAIN",
		"DUPLICATE",
		"REPLACE",
	}
}

// GetMappingCopyConflictResolutionCopyConflictResolutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCopyConflictResolutionCopyConflictResolutionTypeEnum(val string) (CopyConflictResolutionCopyConflictResolutionTypeEnum, bool) {
	enum, ok := mappingCopyConflictResolutionCopyConflictResolutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
