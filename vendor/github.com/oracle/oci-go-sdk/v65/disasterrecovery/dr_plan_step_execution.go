// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanStepExecution The details of a step execution in a DR plan execution.
type DrPlanStepExecution struct {

	// The unique id of the step. Must not be modified by user.
	// Example: `sgid1.step..uniqueID`
	StepId *string `mandatory:"true" json:"stepId"`

	// The step type.
	Type DrPlanStepTypeEnum `mandatory:"true" json:"type"`

	// The unique id of the group to which this step belongs. Must not be modified by user.
	// Example: `sgid1.group..uniqueID`
	GroupId *string `mandatory:"true" json:"groupId"`

	// The display name of the step execution.
	// Example: `DATABASE_SWITCHOVER`
	DisplayName *string `mandatory:"true" json:"displayName"`

	LogLocation *ObjectStorageLogLocation `mandatory:"true" json:"logLocation"`

	// The status of the step execution.
	Status DrPlanStepExecutionStatusEnum `mandatory:"true" json:"status"`

	// Additional details on the step execution status.
	// Example: `This step failed to complete due to a timeout`
	StatusDetails *string `mandatory:"false" json:"statusDetails"`

	// The time when step execution began. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time when execution ended. An RFC3339 formatted datetime string.
	// Example: `2019-03-29T09:36:42Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The total duration in seconds taken to complete the step execution.
	// Example: `35`
	ExecutionDurationInSec *int `mandatory:"false" json:"executionDurationInSec"`
}

func (m DrPlanStepExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanStepExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanStepTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanStepTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrPlanStepExecutionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDrPlanStepExecutionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
