// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousPatch The representation of AutonomousPatch
type AutonomousPatch struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	Id *string `mandatory:"true" json:"id"`

	// The text describing this patch package.
	Description *string `mandatory:"true" json:"description"`

	// The type of patch. BUNDLE is one example.
	Type *string `mandatory:"true" json:"type"`

	// The date and time that the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of this patch package.
	Version *string `mandatory:"true" json:"version"`

	// A descriptive text associated with the lifecycleState.
	// Typically can contain additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the patch as a result of lastAction.
	LifecycleState AutonomousPatchLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Database patching model preference. See My Oracle Support note 2285040.1 (https://support.oracle.com/rs?type=doc&id=2285040.1) for information on the Release Update (RU) and Release Update Revision (RUR) patching models.
	PatchModel AutonomousPatchPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	// First month of the quarter in which the patch was released.
	Quarter *string `mandatory:"false" json:"quarter"`

	// Year in which the patch was released.
	Year *string `mandatory:"false" json:"year"`
}

func (m AutonomousPatch) String() string {
	return common.PointerString(m)
}

// AutonomousPatchLifecycleStateEnum Enum with underlying type: string
type AutonomousPatchLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousPatchLifecycleStateEnum
const (
	AutonomousPatchLifecycleStateAvailable  AutonomousPatchLifecycleStateEnum = "AVAILABLE"
	AutonomousPatchLifecycleStateSuccess    AutonomousPatchLifecycleStateEnum = "SUCCESS"
	AutonomousPatchLifecycleStateInProgress AutonomousPatchLifecycleStateEnum = "IN_PROGRESS"
	AutonomousPatchLifecycleStateFailed     AutonomousPatchLifecycleStateEnum = "FAILED"
)

var mappingAutonomousPatchLifecycleState = map[string]AutonomousPatchLifecycleStateEnum{
	"AVAILABLE":   AutonomousPatchLifecycleStateAvailable,
	"SUCCESS":     AutonomousPatchLifecycleStateSuccess,
	"IN_PROGRESS": AutonomousPatchLifecycleStateInProgress,
	"FAILED":      AutonomousPatchLifecycleStateFailed,
}

// GetAutonomousPatchLifecycleStateEnumValues Enumerates the set of values for AutonomousPatchLifecycleStateEnum
func GetAutonomousPatchLifecycleStateEnumValues() []AutonomousPatchLifecycleStateEnum {
	values := make([]AutonomousPatchLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousPatchLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousPatchPatchModelEnum Enum with underlying type: string
type AutonomousPatchPatchModelEnum string

// Set of constants representing the allowable values for AutonomousPatchPatchModelEnum
const (
	AutonomousPatchPatchModelUpdates         AutonomousPatchPatchModelEnum = "RELEASE_UPDATES"
	AutonomousPatchPatchModelUpdateRevisions AutonomousPatchPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousPatchPatchModel = map[string]AutonomousPatchPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousPatchPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousPatchPatchModelUpdateRevisions,
}

// GetAutonomousPatchPatchModelEnumValues Enumerates the set of values for AutonomousPatchPatchModelEnum
func GetAutonomousPatchPatchModelEnumValues() []AutonomousPatchPatchModelEnum {
	values := make([]AutonomousPatchPatchModelEnum, 0)
	for _, v := range mappingAutonomousPatchPatchModel {
		values = append(values, v)
	}
	return values
}
