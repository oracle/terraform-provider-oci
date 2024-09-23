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

// JobActivity Description of JobActivity.
type JobActivity struct {

	// Unique activity id at action group level
	Id *string `mandatory:"true" json:"id"`

	// Status of the Job at Action Group Level
	Status JobStatusEnum `mandatory:"true" json:"status"`

	// The time the the Scheduler Job started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Scheduler Job ended. An RFC3339 formatted datetime string
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// ID of the runbook
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the runbook
	RunbookName *string `mandatory:"false" json:"runbookName"`

	// Resources execution details and outcomes associated with the Task.
	ResourceLevelExecutions []EntityExecutionDetails `mandatory:"false" json:"resourceLevelExecutions"`
}

func (m JobActivity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobActivity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
