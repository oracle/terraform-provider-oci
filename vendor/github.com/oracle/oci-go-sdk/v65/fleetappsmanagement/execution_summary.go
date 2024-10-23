// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecutionSummary A task associated with the Job.
type ExecutionSummary struct {

	// Unique Id associated with the task execution.
	Id *string `mandatory:"true" json:"id"`

	// Status of the Task.
	Status JobStatusEnum `mandatory:"true" json:"status"`

	// The OCID of taskRecord.
	TaskRecordId *string `mandatory:"false" json:"taskRecordId"`

	// Name of the Step.
	StepName *string `mandatory:"false" json:"stepName"`

	// Unique process-reference identifier returned by the execution client.
	// In some cases, this can be a runcommand OCID.
	ProcessReferenceId *string `mandatory:"false" json:"processReferenceId"`

	// The sequence of the task.
	Sequence *string `mandatory:"false" json:"sequence"`

	// Target associated with the execution.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The time the task started. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the task ended. An RFC3339 formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Is this a rollback task?
	IsRollbackTask *bool `mandatory:"false" json:"isRollbackTask"`

	// Description of the Execution status.
	// If there are any errors, this can also include a short error message.
	Description *string `mandatory:"false" json:"description"`

	// Resource Identifier associated with the Work Request.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
