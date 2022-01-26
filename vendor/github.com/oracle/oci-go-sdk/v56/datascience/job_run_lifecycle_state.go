// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

// JobRunLifecycleStateEnum Enum with underlying type: string
type JobRunLifecycleStateEnum string

// Set of constants representing the allowable values for JobRunLifecycleStateEnum
const (
	JobRunLifecycleStateAccepted       JobRunLifecycleStateEnum = "ACCEPTED"
	JobRunLifecycleStateInProgress     JobRunLifecycleStateEnum = "IN_PROGRESS"
	JobRunLifecycleStateFailed         JobRunLifecycleStateEnum = "FAILED"
	JobRunLifecycleStateSucceeded      JobRunLifecycleStateEnum = "SUCCEEDED"
	JobRunLifecycleStateCanceling      JobRunLifecycleStateEnum = "CANCELING"
	JobRunLifecycleStateCanceled       JobRunLifecycleStateEnum = "CANCELED"
	JobRunLifecycleStateDeleted        JobRunLifecycleStateEnum = "DELETED"
	JobRunLifecycleStateNeedsAttention JobRunLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingJobRunLifecycleState = map[string]JobRunLifecycleStateEnum{
	"ACCEPTED":        JobRunLifecycleStateAccepted,
	"IN_PROGRESS":     JobRunLifecycleStateInProgress,
	"FAILED":          JobRunLifecycleStateFailed,
	"SUCCEEDED":       JobRunLifecycleStateSucceeded,
	"CANCELING":       JobRunLifecycleStateCanceling,
	"CANCELED":        JobRunLifecycleStateCanceled,
	"DELETED":         JobRunLifecycleStateDeleted,
	"NEEDS_ATTENTION": JobRunLifecycleStateNeedsAttention,
}

// GetJobRunLifecycleStateEnumValues Enumerates the set of values for JobRunLifecycleStateEnum
func GetJobRunLifecycleStateEnumValues() []JobRunLifecycleStateEnum {
	values := make([]JobRunLifecycleStateEnum, 0)
	for _, v := range mappingJobRunLifecycleState {
		values = append(values, v)
	}
	return values
}
