// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableConditions The details of the audit policy provisioning conditions.
type EnableConditions struct {

	// The entity include or exclude selection.
	EntitySelection EnableConditionsEntitySelectionEnum `mandatory:"true" json:"entitySelection"`

	// The entity type that the policy must be enabled for.
	EntityType EnableConditionsEntityTypeEnum `mandatory:"true" json:"entityType"`

	// The operation status that the policy must be enabled for.
	OperationStatus EnableConditionsOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// List of users or roles that the policy must be enabled for.
	EntityNames []string `mandatory:"false" json:"entityNames"`
}

func (m EnableConditions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableConditions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnableConditionsEntitySelectionEnum(string(m.EntitySelection)); !ok && m.EntitySelection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntitySelection: %s. Supported values are: %s.", m.EntitySelection, strings.Join(GetEnableConditionsEntitySelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnableConditionsEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetEnableConditionsEntityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnableConditionsOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetEnableConditionsOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableConditionsEntitySelectionEnum Enum with underlying type: string
type EnableConditionsEntitySelectionEnum string

// Set of constants representing the allowable values for EnableConditionsEntitySelectionEnum
const (
	EnableConditionsEntitySelectionInclude EnableConditionsEntitySelectionEnum = "INCLUDE"
	EnableConditionsEntitySelectionExclude EnableConditionsEntitySelectionEnum = "EXCLUDE"
)

var mappingEnableConditionsEntitySelectionEnum = map[string]EnableConditionsEntitySelectionEnum{
	"INCLUDE": EnableConditionsEntitySelectionInclude,
	"EXCLUDE": EnableConditionsEntitySelectionExclude,
}

var mappingEnableConditionsEntitySelectionEnumLowerCase = map[string]EnableConditionsEntitySelectionEnum{
	"include": EnableConditionsEntitySelectionInclude,
	"exclude": EnableConditionsEntitySelectionExclude,
}

// GetEnableConditionsEntitySelectionEnumValues Enumerates the set of values for EnableConditionsEntitySelectionEnum
func GetEnableConditionsEntitySelectionEnumValues() []EnableConditionsEntitySelectionEnum {
	values := make([]EnableConditionsEntitySelectionEnum, 0)
	for _, v := range mappingEnableConditionsEntitySelectionEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableConditionsEntitySelectionEnumStringValues Enumerates the set of values in String for EnableConditionsEntitySelectionEnum
func GetEnableConditionsEntitySelectionEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingEnableConditionsEntitySelectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableConditionsEntitySelectionEnum(val string) (EnableConditionsEntitySelectionEnum, bool) {
	enum, ok := mappingEnableConditionsEntitySelectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EnableConditionsEntityTypeEnum Enum with underlying type: string
type EnableConditionsEntityTypeEnum string

// Set of constants representing the allowable values for EnableConditionsEntityTypeEnum
const (
	EnableConditionsEntityTypeUser     EnableConditionsEntityTypeEnum = "USER"
	EnableConditionsEntityTypeRole     EnableConditionsEntityTypeEnum = "ROLE"
	EnableConditionsEntityTypeAllUsers EnableConditionsEntityTypeEnum = "ALL_USERS"
)

var mappingEnableConditionsEntityTypeEnum = map[string]EnableConditionsEntityTypeEnum{
	"USER":      EnableConditionsEntityTypeUser,
	"ROLE":      EnableConditionsEntityTypeRole,
	"ALL_USERS": EnableConditionsEntityTypeAllUsers,
}

var mappingEnableConditionsEntityTypeEnumLowerCase = map[string]EnableConditionsEntityTypeEnum{
	"user":      EnableConditionsEntityTypeUser,
	"role":      EnableConditionsEntityTypeRole,
	"all_users": EnableConditionsEntityTypeAllUsers,
}

// GetEnableConditionsEntityTypeEnumValues Enumerates the set of values for EnableConditionsEntityTypeEnum
func GetEnableConditionsEntityTypeEnumValues() []EnableConditionsEntityTypeEnum {
	values := make([]EnableConditionsEntityTypeEnum, 0)
	for _, v := range mappingEnableConditionsEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableConditionsEntityTypeEnumStringValues Enumerates the set of values in String for EnableConditionsEntityTypeEnum
func GetEnableConditionsEntityTypeEnumStringValues() []string {
	return []string{
		"USER",
		"ROLE",
		"ALL_USERS",
	}
}

// GetMappingEnableConditionsEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableConditionsEntityTypeEnum(val string) (EnableConditionsEntityTypeEnum, bool) {
	enum, ok := mappingEnableConditionsEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EnableConditionsOperationStatusEnum Enum with underlying type: string
type EnableConditionsOperationStatusEnum string

// Set of constants representing the allowable values for EnableConditionsOperationStatusEnum
const (
	EnableConditionsOperationStatusSuccess EnableConditionsOperationStatusEnum = "SUCCESS"
	EnableConditionsOperationStatusFailure EnableConditionsOperationStatusEnum = "FAILURE"
	EnableConditionsOperationStatusBoth    EnableConditionsOperationStatusEnum = "BOTH"
)

var mappingEnableConditionsOperationStatusEnum = map[string]EnableConditionsOperationStatusEnum{
	"SUCCESS": EnableConditionsOperationStatusSuccess,
	"FAILURE": EnableConditionsOperationStatusFailure,
	"BOTH":    EnableConditionsOperationStatusBoth,
}

var mappingEnableConditionsOperationStatusEnumLowerCase = map[string]EnableConditionsOperationStatusEnum{
	"success": EnableConditionsOperationStatusSuccess,
	"failure": EnableConditionsOperationStatusFailure,
	"both":    EnableConditionsOperationStatusBoth,
}

// GetEnableConditionsOperationStatusEnumValues Enumerates the set of values for EnableConditionsOperationStatusEnum
func GetEnableConditionsOperationStatusEnumValues() []EnableConditionsOperationStatusEnum {
	values := make([]EnableConditionsOperationStatusEnum, 0)
	for _, v := range mappingEnableConditionsOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableConditionsOperationStatusEnumStringValues Enumerates the set of values in String for EnableConditionsOperationStatusEnum
func GetEnableConditionsOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
		"BOTH",
	}
}

// GetMappingEnableConditionsOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableConditionsOperationStatusEnum(val string) (EnableConditionsOperationStatusEnum, bool) {
	enum, ok := mappingEnableConditionsOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
