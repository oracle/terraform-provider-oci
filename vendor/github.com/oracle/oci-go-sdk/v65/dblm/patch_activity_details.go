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

// PatchActivityDetails Details of deploy, update and migrate-listener(only for single Instance database) operations for this resource.
type PatchActivityDetails struct {

	// Operation Identifier for deploy operation.
	DeployOperationId *string `mandatory:"false" json:"deployOperationId"`

	// Task identifier for deploy operation.
	DeployTaskId *string `mandatory:"false" json:"deployTaskId"`

	// Status of deploy operation.
	DeployStatus PatchActivityDetailsDeployStatusEnum `mandatory:"false" json:"deployStatus,omitempty"`

	// Operation Identifier for update operation.
	UpdateOperationId *string `mandatory:"false" json:"updateOperationId"`

	// Task identifier for update operation.
	UpdateTaskId *string `mandatory:"false" json:"updateTaskId"`

	// Status of update operation.
	UpdateStatus PatchActivityDetailsUpdateStatusEnum `mandatory:"false" json:"updateStatus,omitempty"`

	// Operation Identifier for migrate listener operation.
	MigrateListenerOperationId *string `mandatory:"false" json:"migrateListenerOperationId"`

	// Task identifier for migrate listener operation.
	MigrateListenerTaskId *string `mandatory:"false" json:"migrateListenerTaskId"`

	// Status of migrate listener operation.
	MigrateListenerStatus PatchActivityDetailsMigrateListenerStatusEnum `mandatory:"false" json:"migrateListenerStatus,omitempty"`
}

func (m PatchActivityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchActivityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchActivityDetailsDeployStatusEnum(string(m.DeployStatus)); !ok && m.DeployStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeployStatus: %s. Supported values are: %s.", m.DeployStatus, strings.Join(GetPatchActivityDetailsDeployStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchActivityDetailsUpdateStatusEnum(string(m.UpdateStatus)); !ok && m.UpdateStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateStatus: %s. Supported values are: %s.", m.UpdateStatus, strings.Join(GetPatchActivityDetailsUpdateStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchActivityDetailsMigrateListenerStatusEnum(string(m.MigrateListenerStatus)); !ok && m.MigrateListenerStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MigrateListenerStatus: %s. Supported values are: %s.", m.MigrateListenerStatus, strings.Join(GetPatchActivityDetailsMigrateListenerStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchActivityDetailsDeployStatusEnum Enum with underlying type: string
type PatchActivityDetailsDeployStatusEnum string

// Set of constants representing the allowable values for PatchActivityDetailsDeployStatusEnum
const (
	PatchActivityDetailsDeployStatusScheduled PatchActivityDetailsDeployStatusEnum = "SCHEDULED"
	PatchActivityDetailsDeployStatusRunning   PatchActivityDetailsDeployStatusEnum = "RUNNING"
	PatchActivityDetailsDeployStatusCompleted PatchActivityDetailsDeployStatusEnum = "COMPLETED"
	PatchActivityDetailsDeployStatusFailed    PatchActivityDetailsDeployStatusEnum = "FAILED"
)

var mappingPatchActivityDetailsDeployStatusEnum = map[string]PatchActivityDetailsDeployStatusEnum{
	"SCHEDULED": PatchActivityDetailsDeployStatusScheduled,
	"RUNNING":   PatchActivityDetailsDeployStatusRunning,
	"COMPLETED": PatchActivityDetailsDeployStatusCompleted,
	"FAILED":    PatchActivityDetailsDeployStatusFailed,
}

var mappingPatchActivityDetailsDeployStatusEnumLowerCase = map[string]PatchActivityDetailsDeployStatusEnum{
	"scheduled": PatchActivityDetailsDeployStatusScheduled,
	"running":   PatchActivityDetailsDeployStatusRunning,
	"completed": PatchActivityDetailsDeployStatusCompleted,
	"failed":    PatchActivityDetailsDeployStatusFailed,
}

// GetPatchActivityDetailsDeployStatusEnumValues Enumerates the set of values for PatchActivityDetailsDeployStatusEnum
func GetPatchActivityDetailsDeployStatusEnumValues() []PatchActivityDetailsDeployStatusEnum {
	values := make([]PatchActivityDetailsDeployStatusEnum, 0)
	for _, v := range mappingPatchActivityDetailsDeployStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchActivityDetailsDeployStatusEnumStringValues Enumerates the set of values in String for PatchActivityDetailsDeployStatusEnum
func GetPatchActivityDetailsDeployStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingPatchActivityDetailsDeployStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchActivityDetailsDeployStatusEnum(val string) (PatchActivityDetailsDeployStatusEnum, bool) {
	enum, ok := mappingPatchActivityDetailsDeployStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchActivityDetailsUpdateStatusEnum Enum with underlying type: string
type PatchActivityDetailsUpdateStatusEnum string

// Set of constants representing the allowable values for PatchActivityDetailsUpdateStatusEnum
const (
	PatchActivityDetailsUpdateStatusScheduled PatchActivityDetailsUpdateStatusEnum = "SCHEDULED"
	PatchActivityDetailsUpdateStatusRunning   PatchActivityDetailsUpdateStatusEnum = "RUNNING"
	PatchActivityDetailsUpdateStatusCompleted PatchActivityDetailsUpdateStatusEnum = "COMPLETED"
	PatchActivityDetailsUpdateStatusFailed    PatchActivityDetailsUpdateStatusEnum = "FAILED"
)

var mappingPatchActivityDetailsUpdateStatusEnum = map[string]PatchActivityDetailsUpdateStatusEnum{
	"SCHEDULED": PatchActivityDetailsUpdateStatusScheduled,
	"RUNNING":   PatchActivityDetailsUpdateStatusRunning,
	"COMPLETED": PatchActivityDetailsUpdateStatusCompleted,
	"FAILED":    PatchActivityDetailsUpdateStatusFailed,
}

var mappingPatchActivityDetailsUpdateStatusEnumLowerCase = map[string]PatchActivityDetailsUpdateStatusEnum{
	"scheduled": PatchActivityDetailsUpdateStatusScheduled,
	"running":   PatchActivityDetailsUpdateStatusRunning,
	"completed": PatchActivityDetailsUpdateStatusCompleted,
	"failed":    PatchActivityDetailsUpdateStatusFailed,
}

// GetPatchActivityDetailsUpdateStatusEnumValues Enumerates the set of values for PatchActivityDetailsUpdateStatusEnum
func GetPatchActivityDetailsUpdateStatusEnumValues() []PatchActivityDetailsUpdateStatusEnum {
	values := make([]PatchActivityDetailsUpdateStatusEnum, 0)
	for _, v := range mappingPatchActivityDetailsUpdateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchActivityDetailsUpdateStatusEnumStringValues Enumerates the set of values in String for PatchActivityDetailsUpdateStatusEnum
func GetPatchActivityDetailsUpdateStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingPatchActivityDetailsUpdateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchActivityDetailsUpdateStatusEnum(val string) (PatchActivityDetailsUpdateStatusEnum, bool) {
	enum, ok := mappingPatchActivityDetailsUpdateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchActivityDetailsMigrateListenerStatusEnum Enum with underlying type: string
type PatchActivityDetailsMigrateListenerStatusEnum string

// Set of constants representing the allowable values for PatchActivityDetailsMigrateListenerStatusEnum
const (
	PatchActivityDetailsMigrateListenerStatusScheduled PatchActivityDetailsMigrateListenerStatusEnum = "SCHEDULED"
	PatchActivityDetailsMigrateListenerStatusRunning   PatchActivityDetailsMigrateListenerStatusEnum = "RUNNING"
	PatchActivityDetailsMigrateListenerStatusCompleted PatchActivityDetailsMigrateListenerStatusEnum = "COMPLETED"
	PatchActivityDetailsMigrateListenerStatusFailed    PatchActivityDetailsMigrateListenerStatusEnum = "FAILED"
	PatchActivityDetailsMigrateListenerStatusNa        PatchActivityDetailsMigrateListenerStatusEnum = "NA"
)

var mappingPatchActivityDetailsMigrateListenerStatusEnum = map[string]PatchActivityDetailsMigrateListenerStatusEnum{
	"SCHEDULED": PatchActivityDetailsMigrateListenerStatusScheduled,
	"RUNNING":   PatchActivityDetailsMigrateListenerStatusRunning,
	"COMPLETED": PatchActivityDetailsMigrateListenerStatusCompleted,
	"FAILED":    PatchActivityDetailsMigrateListenerStatusFailed,
	"NA":        PatchActivityDetailsMigrateListenerStatusNa,
}

var mappingPatchActivityDetailsMigrateListenerStatusEnumLowerCase = map[string]PatchActivityDetailsMigrateListenerStatusEnum{
	"scheduled": PatchActivityDetailsMigrateListenerStatusScheduled,
	"running":   PatchActivityDetailsMigrateListenerStatusRunning,
	"completed": PatchActivityDetailsMigrateListenerStatusCompleted,
	"failed":    PatchActivityDetailsMigrateListenerStatusFailed,
	"na":        PatchActivityDetailsMigrateListenerStatusNa,
}

// GetPatchActivityDetailsMigrateListenerStatusEnumValues Enumerates the set of values for PatchActivityDetailsMigrateListenerStatusEnum
func GetPatchActivityDetailsMigrateListenerStatusEnumValues() []PatchActivityDetailsMigrateListenerStatusEnum {
	values := make([]PatchActivityDetailsMigrateListenerStatusEnum, 0)
	for _, v := range mappingPatchActivityDetailsMigrateListenerStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchActivityDetailsMigrateListenerStatusEnumStringValues Enumerates the set of values in String for PatchActivityDetailsMigrateListenerStatusEnum
func GetPatchActivityDetailsMigrateListenerStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"NA",
	}
}

// GetMappingPatchActivityDetailsMigrateListenerStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchActivityDetailsMigrateListenerStatusEnum(val string) (PatchActivityDetailsMigrateListenerStatusEnum, bool) {
	enum, ok := mappingPatchActivityDetailsMigrateListenerStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
