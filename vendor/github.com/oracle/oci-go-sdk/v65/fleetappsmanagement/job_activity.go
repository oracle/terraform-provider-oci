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

// JobActivity Activity details including status corresponding to an Action Group.
type JobActivity struct {

	// Unique activity id at the action group level.
	// In most cases, this would be a generated ActionGroupId.
	Id *string `mandatory:"true" json:"id"`

	// Status of the Job at Action Group Level.
	Status JobStatusEnum `mandatory:"true" json:"status"`

	// The time the execution for the Action Group started. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the execution for the Action Group ended. An RFC3339 formatted datetime string
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// OCID of the runbook associated with the Action Group.
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the runbook associated with the Action Group.
	RunbookName *string `mandatory:"false" json:"runbookName"`

	// A description of the Job Activity status.
	// If there are any errors, this can also include a short error message.
	Description *string `mandatory:"false" json:"description"`

	// List of Resource executions associated with the Action Group.
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
