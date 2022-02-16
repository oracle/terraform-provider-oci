// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryMirrorRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryMirrorRecordMirrorStatusEnum(string(m.MirrorStatus)); !ok && m.MirrorStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MirrorStatus: %s. Supported values are: %s.", m.MirrorStatus, strings.Join(GetRepositoryMirrorRecordMirrorStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingRepositoryMirrorRecordMirrorStatusEnum = map[string]RepositoryMirrorRecordMirrorStatusEnum{
	"NONE":    RepositoryMirrorRecordMirrorStatusNone,
	"QUEUED":  RepositoryMirrorRecordMirrorStatusQueued,
	"RUNNING": RepositoryMirrorRecordMirrorStatusRunning,
	"PASSED":  RepositoryMirrorRecordMirrorStatusPassed,
	"FAILED":  RepositoryMirrorRecordMirrorStatusFailed,
}

// GetRepositoryMirrorRecordMirrorStatusEnumValues Enumerates the set of values for RepositoryMirrorRecordMirrorStatusEnum
func GetRepositoryMirrorRecordMirrorStatusEnumValues() []RepositoryMirrorRecordMirrorStatusEnum {
	values := make([]RepositoryMirrorRecordMirrorStatusEnum, 0)
	for _, v := range mappingRepositoryMirrorRecordMirrorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryMirrorRecordMirrorStatusEnumStringValues Enumerates the set of values in String for RepositoryMirrorRecordMirrorStatusEnum
func GetRepositoryMirrorRecordMirrorStatusEnumStringValues() []string {
	return []string{
		"NONE",
		"QUEUED",
		"RUNNING",
		"PASSED",
		"FAILED",
	}
}

// GetMappingRepositoryMirrorRecordMirrorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryMirrorRecordMirrorStatusEnum(val string) (RepositoryMirrorRecordMirrorStatusEnum, bool) {
	mappingRepositoryMirrorRecordMirrorStatusEnumIgnoreCase := make(map[string]RepositoryMirrorRecordMirrorStatusEnum)
	for k, v := range mappingRepositoryMirrorRecordMirrorStatusEnum {
		mappingRepositoryMirrorRecordMirrorStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRepositoryMirrorRecordMirrorStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
