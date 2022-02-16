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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobExecutionsStatusSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobExecutionsStatusSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobExecutionsStatusSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobExecutionsStatusSummaryStatusEnum Enum with underlying type: string
type JobExecutionsStatusSummaryStatusEnum string

// Set of constants representing the allowable values for JobExecutionsStatusSummaryStatusEnum
const (
	JobExecutionsStatusSummaryStatusSucceeded  JobExecutionsStatusSummaryStatusEnum = "SUCCEEDED"
	JobExecutionsStatusSummaryStatusFailed     JobExecutionsStatusSummaryStatusEnum = "FAILED"
	JobExecutionsStatusSummaryStatusInProgress JobExecutionsStatusSummaryStatusEnum = "IN_PROGRESS"
)

var mappingJobExecutionsStatusSummaryStatusEnum = map[string]JobExecutionsStatusSummaryStatusEnum{
	"SUCCEEDED":   JobExecutionsStatusSummaryStatusSucceeded,
	"FAILED":      JobExecutionsStatusSummaryStatusFailed,
	"IN_PROGRESS": JobExecutionsStatusSummaryStatusInProgress,
}

// GetJobExecutionsStatusSummaryStatusEnumValues Enumerates the set of values for JobExecutionsStatusSummaryStatusEnum
func GetJobExecutionsStatusSummaryStatusEnumValues() []JobExecutionsStatusSummaryStatusEnum {
	values := make([]JobExecutionsStatusSummaryStatusEnum, 0)
	for _, v := range mappingJobExecutionsStatusSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJobExecutionsStatusSummaryStatusEnumStringValues Enumerates the set of values in String for JobExecutionsStatusSummaryStatusEnum
func GetJobExecutionsStatusSummaryStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingJobExecutionsStatusSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobExecutionsStatusSummaryStatusEnum(val string) (JobExecutionsStatusSummaryStatusEnum, bool) {
	mappingJobExecutionsStatusSummaryStatusEnumIgnoreCase := make(map[string]JobExecutionsStatusSummaryStatusEnum)
	for k, v := range mappingJobExecutionsStatusSummaryStatusEnum {
		mappingJobExecutionsStatusSummaryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobExecutionsStatusSummaryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
