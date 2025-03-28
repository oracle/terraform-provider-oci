// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobRunSummary A summary of a specific job run.
type JobRunSummary struct {

	// The identifier of the job run.
	Id *string `mandatory:"true" json:"id"`

	// The name of the job run.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the parent job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The status of the job run.
	RunStatus JobRunRunStatusEnum `mandatory:"true" json:"runStatus"`

	// The date and time when the job run was submitted.
	TimeSubmitted *common.SDKTime `mandatory:"true" json:"timeSubmitted"`

	// The date and time when the job run was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the parent job has to be executed.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database where the parent job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`
}

func (m JobRunSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobRunSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobRunRunStatusEnum(string(m.RunStatus)); !ok && m.RunStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunStatus: %s. Supported values are: %s.", m.RunStatus, strings.Join(GetJobRunRunStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
