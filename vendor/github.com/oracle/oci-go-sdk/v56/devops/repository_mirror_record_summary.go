// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RepositoryMirrorRecordSummary Object containing information about a mirror record.
type RepositoryMirrorRecordSummary struct {

	// Mirror status of current mirror entry.
	// QUEUED - Mirroring Queued
	// RUNNING - Mirroring is Running
	// PASSED - Mirroring Passed
	// FAILED - Mirroring Failed
	MirrorStatus RepositoryMirrorRecordSummaryMirrorStatusEnum `mandatory:"true" json:"mirrorStatus"`

	// Workrequest ID to track current mirror operation.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The time to enqueue a mirror operation.
	TimeEnqueued *common.SDKTime `mandatory:"false" json:"timeEnqueued"`

	// The time to start a mirror operation.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time to complete a mirror operation.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RepositoryMirrorRecordSummary) String() string {
	return common.PointerString(m)
}

// RepositoryMirrorRecordSummaryMirrorStatusEnum Enum with underlying type: string
type RepositoryMirrorRecordSummaryMirrorStatusEnum string

// Set of constants representing the allowable values for RepositoryMirrorRecordSummaryMirrorStatusEnum
const (
	RepositoryMirrorRecordSummaryMirrorStatusNone    RepositoryMirrorRecordSummaryMirrorStatusEnum = "NONE"
	RepositoryMirrorRecordSummaryMirrorStatusQueued  RepositoryMirrorRecordSummaryMirrorStatusEnum = "QUEUED"
	RepositoryMirrorRecordSummaryMirrorStatusRunning RepositoryMirrorRecordSummaryMirrorStatusEnum = "RUNNING"
	RepositoryMirrorRecordSummaryMirrorStatusPassed  RepositoryMirrorRecordSummaryMirrorStatusEnum = "PASSED"
	RepositoryMirrorRecordSummaryMirrorStatusFailed  RepositoryMirrorRecordSummaryMirrorStatusEnum = "FAILED"
)

var mappingRepositoryMirrorRecordSummaryMirrorStatus = map[string]RepositoryMirrorRecordSummaryMirrorStatusEnum{
	"NONE":    RepositoryMirrorRecordSummaryMirrorStatusNone,
	"QUEUED":  RepositoryMirrorRecordSummaryMirrorStatusQueued,
	"RUNNING": RepositoryMirrorRecordSummaryMirrorStatusRunning,
	"PASSED":  RepositoryMirrorRecordSummaryMirrorStatusPassed,
	"FAILED":  RepositoryMirrorRecordSummaryMirrorStatusFailed,
}

// GetRepositoryMirrorRecordSummaryMirrorStatusEnumValues Enumerates the set of values for RepositoryMirrorRecordSummaryMirrorStatusEnum
func GetRepositoryMirrorRecordSummaryMirrorStatusEnumValues() []RepositoryMirrorRecordSummaryMirrorStatusEnum {
	values := make([]RepositoryMirrorRecordSummaryMirrorStatusEnum, 0)
	for _, v := range mappingRepositoryMirrorRecordSummaryMirrorStatus {
		values = append(values, v)
	}
	return values
}
