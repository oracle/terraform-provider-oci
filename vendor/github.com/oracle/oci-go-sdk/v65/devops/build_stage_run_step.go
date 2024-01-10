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

// BuildStageRunStep The details about each step in a build stage.
type BuildStageRunStep struct {

	// Name of the step.
	Name *string `mandatory:"false" json:"name"`

	// State of the step.
	State BuildStageRunStepStateEnum `mandatory:"false" json:"state,omitempty"`

	// Time when the step started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time when the step finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m BuildStageRunStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildStageRunStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildStageRunStepStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetBuildStageRunStepStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildStageRunStepStateEnum Enum with underlying type: string
type BuildStageRunStepStateEnum string

// Set of constants representing the allowable values for BuildStageRunStepStateEnum
const (
	BuildStageRunStepStateWaiting    BuildStageRunStepStateEnum = "WAITING"
	BuildStageRunStepStateInProgress BuildStageRunStepStateEnum = "IN_PROGRESS"
	BuildStageRunStepStateFailed     BuildStageRunStepStateEnum = "FAILED"
	BuildStageRunStepStateSucceeded  BuildStageRunStepStateEnum = "SUCCEEDED"
)

var mappingBuildStageRunStepStateEnum = map[string]BuildStageRunStepStateEnum{
	"WAITING":     BuildStageRunStepStateWaiting,
	"IN_PROGRESS": BuildStageRunStepStateInProgress,
	"FAILED":      BuildStageRunStepStateFailed,
	"SUCCEEDED":   BuildStageRunStepStateSucceeded,
}

var mappingBuildStageRunStepStateEnumLowerCase = map[string]BuildStageRunStepStateEnum{
	"waiting":     BuildStageRunStepStateWaiting,
	"in_progress": BuildStageRunStepStateInProgress,
	"failed":      BuildStageRunStepStateFailed,
	"succeeded":   BuildStageRunStepStateSucceeded,
}

// GetBuildStageRunStepStateEnumValues Enumerates the set of values for BuildStageRunStepStateEnum
func GetBuildStageRunStepStateEnumValues() []BuildStageRunStepStateEnum {
	values := make([]BuildStageRunStepStateEnum, 0)
	for _, v := range mappingBuildStageRunStepStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildStageRunStepStateEnumStringValues Enumerates the set of values in String for BuildStageRunStepStateEnum
func GetBuildStageRunStepStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingBuildStageRunStepStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildStageRunStepStateEnum(val string) (BuildStageRunStepStateEnum, bool) {
	enum, ok := mappingBuildStageRunStepStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
