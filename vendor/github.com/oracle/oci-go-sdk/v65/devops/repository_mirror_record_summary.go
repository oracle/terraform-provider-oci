// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RepositoryMirrorRecordSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryMirrorRecordSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryMirrorRecordSummaryMirrorStatusEnum(string(m.MirrorStatus)); !ok && m.MirrorStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MirrorStatus: %s. Supported values are: %s.", m.MirrorStatus, strings.Join(GetRepositoryMirrorRecordSummaryMirrorStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingRepositoryMirrorRecordSummaryMirrorStatusEnum = map[string]RepositoryMirrorRecordSummaryMirrorStatusEnum{
	"NONE":    RepositoryMirrorRecordSummaryMirrorStatusNone,
	"QUEUED":  RepositoryMirrorRecordSummaryMirrorStatusQueued,
	"RUNNING": RepositoryMirrorRecordSummaryMirrorStatusRunning,
	"PASSED":  RepositoryMirrorRecordSummaryMirrorStatusPassed,
	"FAILED":  RepositoryMirrorRecordSummaryMirrorStatusFailed,
}

var mappingRepositoryMirrorRecordSummaryMirrorStatusEnumLowerCase = map[string]RepositoryMirrorRecordSummaryMirrorStatusEnum{
	"none":    RepositoryMirrorRecordSummaryMirrorStatusNone,
	"queued":  RepositoryMirrorRecordSummaryMirrorStatusQueued,
	"running": RepositoryMirrorRecordSummaryMirrorStatusRunning,
	"passed":  RepositoryMirrorRecordSummaryMirrorStatusPassed,
	"failed":  RepositoryMirrorRecordSummaryMirrorStatusFailed,
}

// GetRepositoryMirrorRecordSummaryMirrorStatusEnumValues Enumerates the set of values for RepositoryMirrorRecordSummaryMirrorStatusEnum
func GetRepositoryMirrorRecordSummaryMirrorStatusEnumValues() []RepositoryMirrorRecordSummaryMirrorStatusEnum {
	values := make([]RepositoryMirrorRecordSummaryMirrorStatusEnum, 0)
	for _, v := range mappingRepositoryMirrorRecordSummaryMirrorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryMirrorRecordSummaryMirrorStatusEnumStringValues Enumerates the set of values in String for RepositoryMirrorRecordSummaryMirrorStatusEnum
func GetRepositoryMirrorRecordSummaryMirrorStatusEnumStringValues() []string {
	return []string{
		"NONE",
		"QUEUED",
		"RUNNING",
		"PASSED",
		"FAILED",
	}
}

// GetMappingRepositoryMirrorRecordSummaryMirrorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryMirrorRecordSummaryMirrorStatusEnum(val string) (RepositoryMirrorRecordSummaryMirrorStatusEnum, bool) {
	enum, ok := mappingRepositoryMirrorRecordSummaryMirrorStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
