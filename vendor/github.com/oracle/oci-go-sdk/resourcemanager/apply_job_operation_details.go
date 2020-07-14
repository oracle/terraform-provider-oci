// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ApplyJobOperationDetails Job details that are specific to apply operations.
type ApplyJobOperationDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the plan job that contains the execution plan used for this job,
	// or `null` if no execution plan was used.
	ExecutionPlanJobId *string `mandatory:"false" json:"executionPlanJobId"`

	// Specifies the source of the execution plan to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan.
	ExecutionPlanStrategy ApplyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m ApplyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ApplyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApplyJobOperationDetails ApplyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeApplyJobOperationDetails
	}{
		"APPLY",
		(MarshalTypeApplyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}

// ApplyJobOperationDetailsExecutionPlanStrategyEnum Enum with underlying type: string
type ApplyJobOperationDetailsExecutionPlanStrategyEnum string

// Set of constants representing the allowable values for ApplyJobOperationDetailsExecutionPlanStrategyEnum
const (
	ApplyJobOperationDetailsExecutionPlanStrategyFromPlanJobId     ApplyJobOperationDetailsExecutionPlanStrategyEnum = "FROM_PLAN_JOB_ID"
	ApplyJobOperationDetailsExecutionPlanStrategyFromLatestPlanJob ApplyJobOperationDetailsExecutionPlanStrategyEnum = "FROM_LATEST_PLAN_JOB"
	ApplyJobOperationDetailsExecutionPlanStrategyAutoApproved      ApplyJobOperationDetailsExecutionPlanStrategyEnum = "AUTO_APPROVED"
)

var mappingApplyJobOperationDetailsExecutionPlanStrategy = map[string]ApplyJobOperationDetailsExecutionPlanStrategyEnum{
	"FROM_PLAN_JOB_ID":     ApplyJobOperationDetailsExecutionPlanStrategyFromPlanJobId,
	"FROM_LATEST_PLAN_JOB": ApplyJobOperationDetailsExecutionPlanStrategyFromLatestPlanJob,
	"AUTO_APPROVED":        ApplyJobOperationDetailsExecutionPlanStrategyAutoApproved,
}

// GetApplyJobOperationDetailsExecutionPlanStrategyEnumValues Enumerates the set of values for ApplyJobOperationDetailsExecutionPlanStrategyEnum
func GetApplyJobOperationDetailsExecutionPlanStrategyEnumValues() []ApplyJobOperationDetailsExecutionPlanStrategyEnum {
	values := make([]ApplyJobOperationDetailsExecutionPlanStrategyEnum, 0)
	for _, v := range mappingApplyJobOperationDetailsExecutionPlanStrategy {
		values = append(values, v)
	}
	return values
}
