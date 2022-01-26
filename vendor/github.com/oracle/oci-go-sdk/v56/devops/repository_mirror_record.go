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

// RepositoryMirrorRecord Object containing information about a mirror record.
type RepositoryMirrorRecord struct {

	// Mirror status of current mirror entry.
	// QUEUED - Mirroring Queued
	// RUNNING - Mirroring is Running
	// PASSED - Mirroring Passed
	// FAILED - Mirroring Failed
	MirrorStatus RepositoryMirrorRecordMirrorStatusEnum `mandatory:"true" json:"mirrorStatus"`

	// Workrequest ID to track current mirror operation.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The time to enqueue a mirror operation.
	TimeEnqueued *common.SDKTime `mandatory:"false" json:"timeEnqueued"`

	// The time to start a mirror operation.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time taken to complete a mirror operation. Value is null if not completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m RepositoryMirrorRecord) String() string {
	return common.PointerString(m)
}

// RepositoryMirrorRecordMirrorStatusEnum Enum with underlying type: string
type RepositoryMirrorRecordMirrorStatusEnum string

// Set of constants representing the allowable values for RepositoryMirrorRecordMirrorStatusEnum
const (
	RepositoryMirrorRecordMirrorStatusNone    RepositoryMirrorRecordMirrorStatusEnum = "NONE"
	RepositoryMirrorRecordMirrorStatusQueued  RepositoryMirrorRecordMirrorStatusEnum = "QUEUED"
	RepositoryMirrorRecordMirrorStatusRunning RepositoryMirrorRecordMirrorStatusEnum = "RUNNING"
	RepositoryMirrorRecordMirrorStatusPassed  RepositoryMirrorRecordMirrorStatusEnum = "PASSED"
	RepositoryMirrorRecordMirrorStatusFailed  RepositoryMirrorRecordMirrorStatusEnum = "FAILED"
)

var mappingRepositoryMirrorRecordMirrorStatus = map[string]RepositoryMirrorRecordMirrorStatusEnum{
	"NONE":    RepositoryMirrorRecordMirrorStatusNone,
	"QUEUED":  RepositoryMirrorRecordMirrorStatusQueued,
	"RUNNING": RepositoryMirrorRecordMirrorStatusRunning,
	"PASSED":  RepositoryMirrorRecordMirrorStatusPassed,
	"FAILED":  RepositoryMirrorRecordMirrorStatusFailed,
}

// GetRepositoryMirrorRecordMirrorStatusEnumValues Enumerates the set of values for RepositoryMirrorRecordMirrorStatusEnum
func GetRepositoryMirrorRecordMirrorStatusEnumValues() []RepositoryMirrorRecordMirrorStatusEnum {
	values := make([]RepositoryMirrorRecordMirrorStatusEnum, 0)
	for _, v := range mappingRepositoryMirrorRecordMirrorStatus {
		values = append(values, v)
	}
	return values
}
