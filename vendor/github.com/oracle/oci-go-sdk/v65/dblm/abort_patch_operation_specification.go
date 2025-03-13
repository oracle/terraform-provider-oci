// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AbortPatchOperationSpecification Specification of tasks to be aborted.
type AbortPatchOperationSpecification struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The type of PatchTask
	TaskType AbortPatchOperationSpecificationTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	Resources *Resources `mandatory:"false" json:"resources"`
}

func (m AbortPatchOperationSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AbortPatchOperationSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAbortPatchOperationSpecificationTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetAbortPatchOperationSpecificationTaskTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbortPatchOperationSpecificationTaskTypeEnum Enum with underlying type: string
type AbortPatchOperationSpecificationTaskTypeEnum string

// Set of constants representing the allowable values for AbortPatchOperationSpecificationTaskTypeEnum
const (
	AbortPatchOperationSpecificationTaskTypeDeploy           AbortPatchOperationSpecificationTaskTypeEnum = "DEPLOY"
	AbortPatchOperationSpecificationTaskTypeMigrateListener  AbortPatchOperationSpecificationTaskTypeEnum = "MIGRATE_LISTENER"
	AbortPatchOperationSpecificationTaskTypeUpdate           AbortPatchOperationSpecificationTaskTypeEnum = "UPDATE"
	AbortPatchOperationSpecificationTaskTypeRollback         AbortPatchOperationSpecificationTaskTypeEnum = "ROLLBACK"
	AbortPatchOperationSpecificationTaskTypeRollbackListener AbortPatchOperationSpecificationTaskTypeEnum = "ROLLBACK_LISTENER"
	AbortPatchOperationSpecificationTaskTypeCleanup          AbortPatchOperationSpecificationTaskTypeEnum = "CLEANUP"
)

var mappingAbortPatchOperationSpecificationTaskTypeEnum = map[string]AbortPatchOperationSpecificationTaskTypeEnum{
	"DEPLOY":            AbortPatchOperationSpecificationTaskTypeDeploy,
	"MIGRATE_LISTENER":  AbortPatchOperationSpecificationTaskTypeMigrateListener,
	"UPDATE":            AbortPatchOperationSpecificationTaskTypeUpdate,
	"ROLLBACK":          AbortPatchOperationSpecificationTaskTypeRollback,
	"ROLLBACK_LISTENER": AbortPatchOperationSpecificationTaskTypeRollbackListener,
	"CLEANUP":           AbortPatchOperationSpecificationTaskTypeCleanup,
}

var mappingAbortPatchOperationSpecificationTaskTypeEnumLowerCase = map[string]AbortPatchOperationSpecificationTaskTypeEnum{
	"deploy":            AbortPatchOperationSpecificationTaskTypeDeploy,
	"migrate_listener":  AbortPatchOperationSpecificationTaskTypeMigrateListener,
	"update":            AbortPatchOperationSpecificationTaskTypeUpdate,
	"rollback":          AbortPatchOperationSpecificationTaskTypeRollback,
	"rollback_listener": AbortPatchOperationSpecificationTaskTypeRollbackListener,
	"cleanup":           AbortPatchOperationSpecificationTaskTypeCleanup,
}

// GetAbortPatchOperationSpecificationTaskTypeEnumValues Enumerates the set of values for AbortPatchOperationSpecificationTaskTypeEnum
func GetAbortPatchOperationSpecificationTaskTypeEnumValues() []AbortPatchOperationSpecificationTaskTypeEnum {
	values := make([]AbortPatchOperationSpecificationTaskTypeEnum, 0)
	for _, v := range mappingAbortPatchOperationSpecificationTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbortPatchOperationSpecificationTaskTypeEnumStringValues Enumerates the set of values in String for AbortPatchOperationSpecificationTaskTypeEnum
func GetAbortPatchOperationSpecificationTaskTypeEnumStringValues() []string {
	return []string{
		"DEPLOY",
		"MIGRATE_LISTENER",
		"UPDATE",
		"ROLLBACK",
		"ROLLBACK_LISTENER",
		"CLEANUP",
	}
}

// GetMappingAbortPatchOperationSpecificationTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbortPatchOperationSpecificationTaskTypeEnum(val string) (AbortPatchOperationSpecificationTaskTypeEnum, bool) {
	enum, ok := mappingAbortPatchOperationSpecificationTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
