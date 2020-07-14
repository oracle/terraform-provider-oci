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

// DestroyJobOperationDetails Job details that are specific to destroy operations.
type DestroyJobOperationDetails struct {

	// Specifies the source of the execution plan to apply.
	// Currently, only `AUTO_APPROVED` is allowed, which indicates that the job
	// will be run without an execution plan.
	ExecutionPlanStrategy DestroyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m DestroyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DestroyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDestroyJobOperationDetails DestroyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeDestroyJobOperationDetails
	}{
		"DESTROY",
		(MarshalTypeDestroyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}

// DestroyJobOperationDetailsExecutionPlanStrategyEnum Enum with underlying type: string
type DestroyJobOperationDetailsExecutionPlanStrategyEnum string

// Set of constants representing the allowable values for DestroyJobOperationDetailsExecutionPlanStrategyEnum
const (
	DestroyJobOperationDetailsExecutionPlanStrategyAutoApproved DestroyJobOperationDetailsExecutionPlanStrategyEnum = "AUTO_APPROVED"
)

var mappingDestroyJobOperationDetailsExecutionPlanStrategy = map[string]DestroyJobOperationDetailsExecutionPlanStrategyEnum{
	"AUTO_APPROVED": DestroyJobOperationDetailsExecutionPlanStrategyAutoApproved,
}

// GetDestroyJobOperationDetailsExecutionPlanStrategyEnumValues Enumerates the set of values for DestroyJobOperationDetailsExecutionPlanStrategyEnum
func GetDestroyJobOperationDetailsExecutionPlanStrategyEnumValues() []DestroyJobOperationDetailsExecutionPlanStrategyEnum {
	values := make([]DestroyJobOperationDetailsExecutionPlanStrategyEnum, 0)
	for _, v := range mappingDestroyJobOperationDetailsExecutionPlanStrategy {
		values = append(values, v)
	}
	return values
}
