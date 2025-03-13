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

// RetryPatchOperationDetails Retry specification for patch operation.
type RetryPatchOperationDetails struct {

	// PatchOperation Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Type of patch operation task
	TaskType RetryPatchOperationDetailsTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	Resources *Resources `mandatory:"false" json:"resources"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RetryPatchOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RetryPatchOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRetryPatchOperationDetailsTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetRetryPatchOperationDetailsTaskTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RetryPatchOperationDetailsTaskTypeEnum Enum with underlying type: string
type RetryPatchOperationDetailsTaskTypeEnum string

// Set of constants representing the allowable values for RetryPatchOperationDetailsTaskTypeEnum
const (
	RetryPatchOperationDetailsTaskTypeDeploy           RetryPatchOperationDetailsTaskTypeEnum = "DEPLOY"
	RetryPatchOperationDetailsTaskTypeMigrateListener  RetryPatchOperationDetailsTaskTypeEnum = "MIGRATE_LISTENER"
	RetryPatchOperationDetailsTaskTypeUpdate           RetryPatchOperationDetailsTaskTypeEnum = "UPDATE"
	RetryPatchOperationDetailsTaskTypeRollback         RetryPatchOperationDetailsTaskTypeEnum = "ROLLBACK"
	RetryPatchOperationDetailsTaskTypeRollbackListener RetryPatchOperationDetailsTaskTypeEnum = "ROLLBACK_LISTENER"
	RetryPatchOperationDetailsTaskTypeCleanup          RetryPatchOperationDetailsTaskTypeEnum = "CLEANUP"
)

var mappingRetryPatchOperationDetailsTaskTypeEnum = map[string]RetryPatchOperationDetailsTaskTypeEnum{
	"DEPLOY":            RetryPatchOperationDetailsTaskTypeDeploy,
	"MIGRATE_LISTENER":  RetryPatchOperationDetailsTaskTypeMigrateListener,
	"UPDATE":            RetryPatchOperationDetailsTaskTypeUpdate,
	"ROLLBACK":          RetryPatchOperationDetailsTaskTypeRollback,
	"ROLLBACK_LISTENER": RetryPatchOperationDetailsTaskTypeRollbackListener,
	"CLEANUP":           RetryPatchOperationDetailsTaskTypeCleanup,
}

var mappingRetryPatchOperationDetailsTaskTypeEnumLowerCase = map[string]RetryPatchOperationDetailsTaskTypeEnum{
	"deploy":            RetryPatchOperationDetailsTaskTypeDeploy,
	"migrate_listener":  RetryPatchOperationDetailsTaskTypeMigrateListener,
	"update":            RetryPatchOperationDetailsTaskTypeUpdate,
	"rollback":          RetryPatchOperationDetailsTaskTypeRollback,
	"rollback_listener": RetryPatchOperationDetailsTaskTypeRollbackListener,
	"cleanup":           RetryPatchOperationDetailsTaskTypeCleanup,
}

// GetRetryPatchOperationDetailsTaskTypeEnumValues Enumerates the set of values for RetryPatchOperationDetailsTaskTypeEnum
func GetRetryPatchOperationDetailsTaskTypeEnumValues() []RetryPatchOperationDetailsTaskTypeEnum {
	values := make([]RetryPatchOperationDetailsTaskTypeEnum, 0)
	for _, v := range mappingRetryPatchOperationDetailsTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRetryPatchOperationDetailsTaskTypeEnumStringValues Enumerates the set of values in String for RetryPatchOperationDetailsTaskTypeEnum
func GetRetryPatchOperationDetailsTaskTypeEnumStringValues() []string {
	return []string{
		"DEPLOY",
		"MIGRATE_LISTENER",
		"UPDATE",
		"ROLLBACK",
		"ROLLBACK_LISTENER",
		"CLEANUP",
	}
}

// GetMappingRetryPatchOperationDetailsTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetryPatchOperationDetailsTaskTypeEnum(val string) (RetryPatchOperationDetailsTaskTypeEnum, bool) {
	enum, ok := mappingRetryPatchOperationDetailsTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
