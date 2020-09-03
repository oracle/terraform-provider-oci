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

// AutonomousPatchSummary A patch for an Autonomous Exadata Infrastructure or Autonomous Container Database.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type AutonomousPatchSummary struct {

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
	LifecycleState AutonomousPatchSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Database patching model preference. See My Oracle Support note 2285040.1 (https://support.oracle.com/rs?type=doc&id=2285040.1) for information on the Release Update (RU) and Release Update Revision (RUR) patching models.
	PatchModel AutonomousPatchSummaryPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	// First month of the quarter in which the patch was released.
	Quarter *string `mandatory:"false" json:"quarter"`

	// Year in which the patch was released.
	Year *string `mandatory:"false" json:"year"`
}

func (m AutonomousPatchSummary) String() string {
	return common.PointerString(m)
}

// AutonomousPatchSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousPatchSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousPatchSummaryLifecycleStateEnum
const (
	AutonomousPatchSummaryLifecycleStateAvailable  AutonomousPatchSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousPatchSummaryLifecycleStateSuccess    AutonomousPatchSummaryLifecycleStateEnum = "SUCCESS"
	AutonomousPatchSummaryLifecycleStateInProgress AutonomousPatchSummaryLifecycleStateEnum = "IN_PROGRESS"
	AutonomousPatchSummaryLifecycleStateFailed     AutonomousPatchSummaryLifecycleStateEnum = "FAILED"
)

var mappingAutonomousPatchSummaryLifecycleState = map[string]AutonomousPatchSummaryLifecycleStateEnum{
	"AVAILABLE":   AutonomousPatchSummaryLifecycleStateAvailable,
	"SUCCESS":     AutonomousPatchSummaryLifecycleStateSuccess,
	"IN_PROGRESS": AutonomousPatchSummaryLifecycleStateInProgress,
	"FAILED":      AutonomousPatchSummaryLifecycleStateFailed,
}

// GetAutonomousPatchSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousPatchSummaryLifecycleStateEnum
func GetAutonomousPatchSummaryLifecycleStateEnumValues() []AutonomousPatchSummaryLifecycleStateEnum {
	values := make([]AutonomousPatchSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousPatchSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousPatchSummaryPatchModelEnum Enum with underlying type: string
type AutonomousPatchSummaryPatchModelEnum string

// Set of constants representing the allowable values for AutonomousPatchSummaryPatchModelEnum
const (
	AutonomousPatchSummaryPatchModelUpdates         AutonomousPatchSummaryPatchModelEnum = "RELEASE_UPDATES"
	AutonomousPatchSummaryPatchModelUpdateRevisions AutonomousPatchSummaryPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousPatchSummaryPatchModel = map[string]AutonomousPatchSummaryPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousPatchSummaryPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousPatchSummaryPatchModelUpdateRevisions,
}

// GetAutonomousPatchSummaryPatchModelEnumValues Enumerates the set of values for AutonomousPatchSummaryPatchModelEnum
func GetAutonomousPatchSummaryPatchModelEnumValues() []AutonomousPatchSummaryPatchModelEnum {
	values := make([]AutonomousPatchSummaryPatchModelEnum, 0)
	for _, v := range mappingAutonomousPatchSummaryPatchModel {
		values = append(values, v)
	}
	return values
}
