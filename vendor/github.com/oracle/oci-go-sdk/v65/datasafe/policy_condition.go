// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PolicyCondition The audit policy provisioning conditions.
type PolicyCondition interface {

	// Specifies whether to include or exclude the specified users or roles.
	GetEntitySelection() PolicyConditionEntitySelectionEnum

	// The operation status that the policy must be enabled for.
	GetOperationStatus() PolicyConditionOperationStatusEnum
}

type policycondition struct {
	JsonData        []byte
	EntitySelection PolicyConditionEntitySelectionEnum `mandatory:"true" json:"entitySelection"`
	OperationStatus PolicyConditionOperationStatusEnum `mandatory:"true" json:"operationStatus"`
	EntityType      string                             `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *policycondition) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpolicycondition policycondition
	s := struct {
		Model Unmarshalerpolicycondition
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntitySelection = s.Model.EntitySelection
	m.OperationStatus = s.Model.OperationStatus
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *policycondition) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "ATTRIBUTE_SET":
		mm := AttributeSetCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER":
		mm := UserCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROLE":
		mm := RoleCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALL_USERS":
		mm := AllUserCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PolicyCondition: %s.", m.EntityType)
		return *m, nil
	}
}

// GetEntitySelection returns EntitySelection
func (m policycondition) GetEntitySelection() PolicyConditionEntitySelectionEnum {
	return m.EntitySelection
}

// GetOperationStatus returns OperationStatus
func (m policycondition) GetOperationStatus() PolicyConditionOperationStatusEnum {
	return m.OperationStatus
}

func (m policycondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m policycondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPolicyConditionEntitySelectionEnum(string(m.EntitySelection)); !ok && m.EntitySelection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntitySelection: %s. Supported values are: %s.", m.EntitySelection, strings.Join(GetPolicyConditionEntitySelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPolicyConditionOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetPolicyConditionOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PolicyConditionEntitySelectionEnum Enum with underlying type: string
type PolicyConditionEntitySelectionEnum string

// Set of constants representing the allowable values for PolicyConditionEntitySelectionEnum
const (
	PolicyConditionEntitySelectionInclude PolicyConditionEntitySelectionEnum = "INCLUDE"
	PolicyConditionEntitySelectionExclude PolicyConditionEntitySelectionEnum = "EXCLUDE"
)

var mappingPolicyConditionEntitySelectionEnum = map[string]PolicyConditionEntitySelectionEnum{
	"INCLUDE": PolicyConditionEntitySelectionInclude,
	"EXCLUDE": PolicyConditionEntitySelectionExclude,
}

var mappingPolicyConditionEntitySelectionEnumLowerCase = map[string]PolicyConditionEntitySelectionEnum{
	"include": PolicyConditionEntitySelectionInclude,
	"exclude": PolicyConditionEntitySelectionExclude,
}

// GetPolicyConditionEntitySelectionEnumValues Enumerates the set of values for PolicyConditionEntitySelectionEnum
func GetPolicyConditionEntitySelectionEnumValues() []PolicyConditionEntitySelectionEnum {
	values := make([]PolicyConditionEntitySelectionEnum, 0)
	for _, v := range mappingPolicyConditionEntitySelectionEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyConditionEntitySelectionEnumStringValues Enumerates the set of values in String for PolicyConditionEntitySelectionEnum
func GetPolicyConditionEntitySelectionEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingPolicyConditionEntitySelectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyConditionEntitySelectionEnum(val string) (PolicyConditionEntitySelectionEnum, bool) {
	enum, ok := mappingPolicyConditionEntitySelectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PolicyConditionOperationStatusEnum Enum with underlying type: string
type PolicyConditionOperationStatusEnum string

// Set of constants representing the allowable values for PolicyConditionOperationStatusEnum
const (
	PolicyConditionOperationStatusSuccess PolicyConditionOperationStatusEnum = "SUCCESS"
	PolicyConditionOperationStatusFailure PolicyConditionOperationStatusEnum = "FAILURE"
	PolicyConditionOperationStatusBoth    PolicyConditionOperationStatusEnum = "BOTH"
)

var mappingPolicyConditionOperationStatusEnum = map[string]PolicyConditionOperationStatusEnum{
	"SUCCESS": PolicyConditionOperationStatusSuccess,
	"FAILURE": PolicyConditionOperationStatusFailure,
	"BOTH":    PolicyConditionOperationStatusBoth,
}

var mappingPolicyConditionOperationStatusEnumLowerCase = map[string]PolicyConditionOperationStatusEnum{
	"success": PolicyConditionOperationStatusSuccess,
	"failure": PolicyConditionOperationStatusFailure,
	"both":    PolicyConditionOperationStatusBoth,
}

// GetPolicyConditionOperationStatusEnumValues Enumerates the set of values for PolicyConditionOperationStatusEnum
func GetPolicyConditionOperationStatusEnumValues() []PolicyConditionOperationStatusEnum {
	values := make([]PolicyConditionOperationStatusEnum, 0)
	for _, v := range mappingPolicyConditionOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyConditionOperationStatusEnumStringValues Enumerates the set of values in String for PolicyConditionOperationStatusEnum
func GetPolicyConditionOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
		"BOTH",
	}
}

// GetMappingPolicyConditionOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyConditionOperationStatusEnum(val string) (PolicyConditionOperationStatusEnum, bool) {
	enum, ok := mappingPolicyConditionOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PolicyConditionEntityTypeEnum Enum with underlying type: string
type PolicyConditionEntityTypeEnum string

// Set of constants representing the allowable values for PolicyConditionEntityTypeEnum
const (
	PolicyConditionEntityTypeUser         PolicyConditionEntityTypeEnum = "USER"
	PolicyConditionEntityTypeRole         PolicyConditionEntityTypeEnum = "ROLE"
	PolicyConditionEntityTypeAllUsers     PolicyConditionEntityTypeEnum = "ALL_USERS"
	PolicyConditionEntityTypeAttributeSet PolicyConditionEntityTypeEnum = "ATTRIBUTE_SET"
)

var mappingPolicyConditionEntityTypeEnum = map[string]PolicyConditionEntityTypeEnum{
	"USER":          PolicyConditionEntityTypeUser,
	"ROLE":          PolicyConditionEntityTypeRole,
	"ALL_USERS":     PolicyConditionEntityTypeAllUsers,
	"ATTRIBUTE_SET": PolicyConditionEntityTypeAttributeSet,
}

var mappingPolicyConditionEntityTypeEnumLowerCase = map[string]PolicyConditionEntityTypeEnum{
	"user":          PolicyConditionEntityTypeUser,
	"role":          PolicyConditionEntityTypeRole,
	"all_users":     PolicyConditionEntityTypeAllUsers,
	"attribute_set": PolicyConditionEntityTypeAttributeSet,
}

// GetPolicyConditionEntityTypeEnumValues Enumerates the set of values for PolicyConditionEntityTypeEnum
func GetPolicyConditionEntityTypeEnumValues() []PolicyConditionEntityTypeEnum {
	values := make([]PolicyConditionEntityTypeEnum, 0)
	for _, v := range mappingPolicyConditionEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyConditionEntityTypeEnumStringValues Enumerates the set of values in String for PolicyConditionEntityTypeEnum
func GetPolicyConditionEntityTypeEnumStringValues() []string {
	return []string{
		"USER",
		"ROLE",
		"ALL_USERS",
		"ATTRIBUTE_SET",
	}
}

// GetMappingPolicyConditionEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyConditionEntityTypeEnum(val string) (PolicyConditionEntityTypeEnum, bool) {
	enum, ok := mappingPolicyConditionEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
