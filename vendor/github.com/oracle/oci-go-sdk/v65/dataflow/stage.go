// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Stage Spark job stage events.
type Stage struct {

	// Spark job stage name.
	Name *string `mandatory:"true" json:"name"`

	// Spark job stage id.
	Id *int `mandatory:"true" json:"id"`

	// Spark job stage status.
	StageStatus *string `mandatory:"false" json:"stageStatus"`

	// The number of tasks that have been successfully completed within the Spark stage.
	SuccessfulTaskCount *int64 `mandatory:"false" json:"successfulTaskCount"`

	// The number of currently running tasks within the Spark stage.
	RunningTaskCount *int64 `mandatory:"false" json:"runningTaskCount"`

	// The number of failed tasks within the Spark stage.
	FailedTaskCount *int64 `mandatory:"false" json:"failedTaskCount"`

	// The number of killed tasks within the Spark stage.
	KilledTaskCount *int64 `mandatory:"false" json:"killedTaskCount"`

	// Total number of tasks for the Spark stage.
	NumTasks *int `mandatory:"false" json:"numTasks"`

	// The attempt number of a Spark stage.
	AttemptNumber *int `mandatory:"false" json:"attemptNumber"`

	// The time when a Spark stage was submitted.
	SubmissionTime *int64 `mandatory:"false" json:"submissionTime"`

	// The time when a Spark stage was completed.
	CompletionTime *int64 `mandatory:"false" json:"completionTime"`
}

func (m Stage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Stage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
