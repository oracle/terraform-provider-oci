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

// BuildRunSnapshot Snapshot of a build run. Contains information including pipelineId, commitId.
type BuildRunSnapshot struct {

	// The OCID of the build pipeline where the build was triggered.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The commit id which the build was triggered from.
	CommitId *string `mandatory:"true" json:"commitId"`

	// The OCID of the build run.
	BuildRunId *string `mandatory:"true" json:"buildRunId"`

	// The current status of the build run.
	LifecycleState BuildRunSnapshotLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The display name of the build run.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the build run was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the build run was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BuildRunSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildRunSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBuildRunSnapshotLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildRunSnapshotLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildRunSnapshotLifecycleStateEnum Enum with underlying type: string
type BuildRunSnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for BuildRunSnapshotLifecycleStateEnum
const (
	BuildRunSnapshotLifecycleStateAccepted   BuildRunSnapshotLifecycleStateEnum = "ACCEPTED"
	BuildRunSnapshotLifecycleStateInProgress BuildRunSnapshotLifecycleStateEnum = "IN_PROGRESS"
	BuildRunSnapshotLifecycleStateFailed     BuildRunSnapshotLifecycleStateEnum = "FAILED"
	BuildRunSnapshotLifecycleStateSucceeded  BuildRunSnapshotLifecycleStateEnum = "SUCCEEDED"
	BuildRunSnapshotLifecycleStateCanceling  BuildRunSnapshotLifecycleStateEnum = "CANCELING"
	BuildRunSnapshotLifecycleStateCanceled   BuildRunSnapshotLifecycleStateEnum = "CANCELED"
)

var mappingBuildRunSnapshotLifecycleStateEnum = map[string]BuildRunSnapshotLifecycleStateEnum{
	"ACCEPTED":    BuildRunSnapshotLifecycleStateAccepted,
	"IN_PROGRESS": BuildRunSnapshotLifecycleStateInProgress,
	"FAILED":      BuildRunSnapshotLifecycleStateFailed,
	"SUCCEEDED":   BuildRunSnapshotLifecycleStateSucceeded,
	"CANCELING":   BuildRunSnapshotLifecycleStateCanceling,
	"CANCELED":    BuildRunSnapshotLifecycleStateCanceled,
}

var mappingBuildRunSnapshotLifecycleStateEnumLowerCase = map[string]BuildRunSnapshotLifecycleStateEnum{
	"accepted":    BuildRunSnapshotLifecycleStateAccepted,
	"in_progress": BuildRunSnapshotLifecycleStateInProgress,
	"failed":      BuildRunSnapshotLifecycleStateFailed,
	"succeeded":   BuildRunSnapshotLifecycleStateSucceeded,
	"canceling":   BuildRunSnapshotLifecycleStateCanceling,
	"canceled":    BuildRunSnapshotLifecycleStateCanceled,
}

// GetBuildRunSnapshotLifecycleStateEnumValues Enumerates the set of values for BuildRunSnapshotLifecycleStateEnum
func GetBuildRunSnapshotLifecycleStateEnumValues() []BuildRunSnapshotLifecycleStateEnum {
	values := make([]BuildRunSnapshotLifecycleStateEnum, 0)
	for _, v := range mappingBuildRunSnapshotLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildRunSnapshotLifecycleStateEnumStringValues Enumerates the set of values in String for BuildRunSnapshotLifecycleStateEnum
func GetBuildRunSnapshotLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBuildRunSnapshotLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildRunSnapshotLifecycleStateEnum(val string) (BuildRunSnapshotLifecycleStateEnum, bool) {
	enum, ok := mappingBuildRunSnapshotLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
