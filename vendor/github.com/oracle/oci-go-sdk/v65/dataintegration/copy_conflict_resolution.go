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

// CopyConflictResolution Copy Object Conflict resolution.
type CopyConflictResolution struct {

	// Copy Object Conflict Resolution Type (RETAIN/DUPLICATE/REPLACE).
	RequestType CopyConflictResolutionRequestTypeEnum `mandatory:"true" json:"requestType"`

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
	if _, ok := GetMappingCopyConflictResolutionRequestTypeEnum(string(m.RequestType)); !ok && m.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", m.RequestType, strings.Join(GetCopyConflictResolutionRequestTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CopyConflictResolutionRequestTypeEnum Enum with underlying type: string
type CopyConflictResolutionRequestTypeEnum string

// Set of constants representing the allowable values for CopyConflictResolutionRequestTypeEnum
const (
	CopyConflictResolutionRequestTypeRetain    CopyConflictResolutionRequestTypeEnum = "RETAIN"
	CopyConflictResolutionRequestTypeDuplicate CopyConflictResolutionRequestTypeEnum = "DUPLICATE"
	CopyConflictResolutionRequestTypeReplace   CopyConflictResolutionRequestTypeEnum = "REPLACE"
)

var mappingCopyConflictResolutionRequestTypeEnum = map[string]CopyConflictResolutionRequestTypeEnum{
	"RETAIN":    CopyConflictResolutionRequestTypeRetain,
	"DUPLICATE": CopyConflictResolutionRequestTypeDuplicate,
	"REPLACE":   CopyConflictResolutionRequestTypeReplace,
}

var mappingCopyConflictResolutionRequestTypeEnumLowerCase = map[string]CopyConflictResolutionRequestTypeEnum{
	"retain":    CopyConflictResolutionRequestTypeRetain,
	"duplicate": CopyConflictResolutionRequestTypeDuplicate,
	"replace":   CopyConflictResolutionRequestTypeReplace,
}

// GetCopyConflictResolutionRequestTypeEnumValues Enumerates the set of values for CopyConflictResolutionRequestTypeEnum
func GetCopyConflictResolutionRequestTypeEnumValues() []CopyConflictResolutionRequestTypeEnum {
	values := make([]CopyConflictResolutionRequestTypeEnum, 0)
	for _, v := range mappingCopyConflictResolutionRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCopyConflictResolutionRequestTypeEnumStringValues Enumerates the set of values in String for CopyConflictResolutionRequestTypeEnum
func GetCopyConflictResolutionRequestTypeEnumStringValues() []string {
	return []string{
		"RETAIN",
		"DUPLICATE",
		"REPLACE",
	}
}

// GetMappingCopyConflictResolutionRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCopyConflictResolutionRequestTypeEnum(val string) (CopyConflictResolutionRequestTypeEnum, bool) {
	enum, ok := mappingCopyConflictResolutionRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
