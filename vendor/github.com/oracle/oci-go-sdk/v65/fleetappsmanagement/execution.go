// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Execution Task Execution associated with the Job.
type Execution struct {

	// Unique Id assocaited with the Task Execution
	Id *string `mandatory:"true" json:"id"`

	// Status of the Task
	Status JobStatusEnum `mandatory:"true" json:"status"`

	// The OCID of taskRecord
	TaskRecordId *string `mandatory:"false" json:"taskRecordId"`

	// Name of the step
	StepName *string `mandatory:"false" json:"stepName"`

	// Unique process reference identifier returned by the execution client
	ProcessReferenceId *string `mandatory:"false" json:"processReferenceId"`

	// The sequence of the task
	Sequence *string `mandatory:"false" json:"sequence"`

	// Subjects which are tied to the task
	Subjects []string `mandatory:"false" json:"subjects"`

	Outcome *Outcome `mandatory:"false" json:"outcome"`

	// Target associated with the execution
	TargetId *string `mandatory:"false" json:"targetId"`

	// The time the task started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the task ended. An RFC3339 formatted datetime string
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Execution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Execution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
