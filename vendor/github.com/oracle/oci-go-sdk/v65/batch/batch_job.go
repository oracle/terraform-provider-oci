// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchJob A batch job represents the execution unit that is invoked on a batch's pools of compute resources.
// A job is composed of one or more tasks, which are the executable commands.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type BatchJob struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent batch job pool.
	BatchJobPoolId *string `mandatory:"true" json:"batchJobPoolId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent batch context of the parent job pool.
	BatchContextId *string `mandatory:"true" json:"batchContextId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the batch job was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the batch job.
	LifecycleState BatchJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A value calculated by the system based on the priority tags.
	Priority *int `mandatory:"false" json:"priority"`

	// A percentage value (0..100) of the job completion.
	ProgressPercentage *int `mandatory:"false" json:"progressPercentage"`

	// A more detailed textual representation of the job completion.
	ProgressDetails *string `mandatory:"false" json:"progressDetails"`

	// An environment variables to use for the job's tasks (can be overridden by task's environment variables).
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// The date and time the batch job was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the batch job in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m BatchJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchJobLifecycleStateEnum Enum with underlying type: string
type BatchJobLifecycleStateEnum string

// Set of constants representing the allowable values for BatchJobLifecycleStateEnum
const (
	BatchJobLifecycleStateAccepted       BatchJobLifecycleStateEnum = "ACCEPTED"
	BatchJobLifecycleStateWaiting        BatchJobLifecycleStateEnum = "WAITING"
	BatchJobLifecycleStateInProgress     BatchJobLifecycleStateEnum = "IN_PROGRESS"
	BatchJobLifecycleStateSucceeded      BatchJobLifecycleStateEnum = "SUCCEEDED"
	BatchJobLifecycleStateNeedsAttention BatchJobLifecycleStateEnum = "NEEDS_ATTENTION"
	BatchJobLifecycleStateFailed         BatchJobLifecycleStateEnum = "FAILED"
	BatchJobLifecycleStateCanceling      BatchJobLifecycleStateEnum = "CANCELING"
	BatchJobLifecycleStateCanceled       BatchJobLifecycleStateEnum = "CANCELED"
)

var mappingBatchJobLifecycleStateEnum = map[string]BatchJobLifecycleStateEnum{
	"ACCEPTED":        BatchJobLifecycleStateAccepted,
	"WAITING":         BatchJobLifecycleStateWaiting,
	"IN_PROGRESS":     BatchJobLifecycleStateInProgress,
	"SUCCEEDED":       BatchJobLifecycleStateSucceeded,
	"NEEDS_ATTENTION": BatchJobLifecycleStateNeedsAttention,
	"FAILED":          BatchJobLifecycleStateFailed,
	"CANCELING":       BatchJobLifecycleStateCanceling,
	"CANCELED":        BatchJobLifecycleStateCanceled,
}

var mappingBatchJobLifecycleStateEnumLowerCase = map[string]BatchJobLifecycleStateEnum{
	"accepted":        BatchJobLifecycleStateAccepted,
	"waiting":         BatchJobLifecycleStateWaiting,
	"in_progress":     BatchJobLifecycleStateInProgress,
	"succeeded":       BatchJobLifecycleStateSucceeded,
	"needs_attention": BatchJobLifecycleStateNeedsAttention,
	"failed":          BatchJobLifecycleStateFailed,
	"canceling":       BatchJobLifecycleStateCanceling,
	"canceled":        BatchJobLifecycleStateCanceled,
}

// GetBatchJobLifecycleStateEnumValues Enumerates the set of values for BatchJobLifecycleStateEnum
func GetBatchJobLifecycleStateEnumValues() []BatchJobLifecycleStateEnum {
	values := make([]BatchJobLifecycleStateEnum, 0)
	for _, v := range mappingBatchJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchJobLifecycleStateEnumStringValues Enumerates the set of values in String for BatchJobLifecycleStateEnum
func GetBatchJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"WAITING",
		"IN_PROGRESS",
		"SUCCEEDED",
		"NEEDS_ATTENTION",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBatchJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchJobLifecycleStateEnum(val string) (BatchJobLifecycleStateEnum, bool) {
	enum, ok := mappingBatchJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
