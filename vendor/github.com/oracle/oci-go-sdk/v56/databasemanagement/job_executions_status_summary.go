// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// JobExecutionsStatusSummary A summary of the status of the job executions.
type JobExecutionsStatusSummary struct {

	// The status of the job execution.
	Status JobExecutionsStatusSummaryStatusEnum `mandatory:"true" json:"status"`

	// The number of job executions of a particular status.
	Count *int `mandatory:"true" json:"count"`
}

func (m JobExecutionsStatusSummary) String() string {
	return common.PointerString(m)
}

// JobExecutionsStatusSummaryStatusEnum Enum with underlying type: string
type JobExecutionsStatusSummaryStatusEnum string

// Set of constants representing the allowable values for JobExecutionsStatusSummaryStatusEnum
const (
	JobExecutionsStatusSummaryStatusSucceeded  JobExecutionsStatusSummaryStatusEnum = "SUCCEEDED"
	JobExecutionsStatusSummaryStatusFailed     JobExecutionsStatusSummaryStatusEnum = "FAILED"
	JobExecutionsStatusSummaryStatusInProgress JobExecutionsStatusSummaryStatusEnum = "IN_PROGRESS"
)

var mappingJobExecutionsStatusSummaryStatus = map[string]JobExecutionsStatusSummaryStatusEnum{
	"SUCCEEDED":   JobExecutionsStatusSummaryStatusSucceeded,
	"FAILED":      JobExecutionsStatusSummaryStatusFailed,
	"IN_PROGRESS": JobExecutionsStatusSummaryStatusInProgress,
}

// GetJobExecutionsStatusSummaryStatusEnumValues Enumerates the set of values for JobExecutionsStatusSummaryStatusEnum
func GetJobExecutionsStatusSummaryStatusEnumValues() []JobExecutionsStatusSummaryStatusEnum {
	values := make([]JobExecutionsStatusSummaryStatusEnum, 0)
	for _, v := range mappingJobExecutionsStatusSummaryStatus {
		values = append(values, v)
	}
	return values
}
