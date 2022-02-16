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

// DeployStageExecutionStep Details about each steps in stage execution for a target environment.
type DeployStageExecutionStep struct {

	// Name of the step.
	Name *string `mandatory:"false" json:"name"`

	// State of the step.
	State DeployStageExecutionStepStateEnum `mandatory:"false" json:"state,omitempty"`

	// Time when the step started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time when the step finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m DeployStageExecutionStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployStageExecutionStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployStageExecutionStepStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetDeployStageExecutionStepStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeployStageExecutionStepStateEnum Enum with underlying type: string
type DeployStageExecutionStepStateEnum string

// Set of constants representing the allowable values for DeployStageExecutionStepStateEnum
const (
	DeployStageExecutionStepStateWaiting    DeployStageExecutionStepStateEnum = "WAITING"
	DeployStageExecutionStepStateInProgress DeployStageExecutionStepStateEnum = "IN_PROGRESS"
	DeployStageExecutionStepStateFailed     DeployStageExecutionStepStateEnum = "FAILED"
	DeployStageExecutionStepStateSucceeded  DeployStageExecutionStepStateEnum = "SUCCEEDED"
	DeployStageExecutionStepStateCanceled   DeployStageExecutionStepStateEnum = "CANCELED"
)

var mappingDeployStageExecutionStepStateEnum = map[string]DeployStageExecutionStepStateEnum{
	"WAITING":     DeployStageExecutionStepStateWaiting,
	"IN_PROGRESS": DeployStageExecutionStepStateInProgress,
	"FAILED":      DeployStageExecutionStepStateFailed,
	"SUCCEEDED":   DeployStageExecutionStepStateSucceeded,
	"CANCELED":    DeployStageExecutionStepStateCanceled,
}

// GetDeployStageExecutionStepStateEnumValues Enumerates the set of values for DeployStageExecutionStepStateEnum
func GetDeployStageExecutionStepStateEnumValues() []DeployStageExecutionStepStateEnum {
	values := make([]DeployStageExecutionStepStateEnum, 0)
	for _, v := range mappingDeployStageExecutionStepStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployStageExecutionStepStateEnumStringValues Enumerates the set of values in String for DeployStageExecutionStepStateEnum
func GetDeployStageExecutionStepStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
	}
}

// GetMappingDeployStageExecutionStepStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployStageExecutionStepStateEnum(val string) (DeployStageExecutionStepStateEnum, bool) {
	mappingDeployStageExecutionStepStateEnumIgnoreCase := make(map[string]DeployStageExecutionStepStateEnum)
	for k, v := range mappingDeployStageExecutionStepStateEnum {
		mappingDeployStageExecutionStepStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDeployStageExecutionStepStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
