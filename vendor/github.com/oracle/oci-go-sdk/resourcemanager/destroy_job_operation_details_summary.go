// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DestroyJobOperationDetailsSummary Job details that are specific to destroy operations.
type DestroyJobOperationDetailsSummary struct {

	// Specifies the source of the execution plan to apply.
	// Currently, only `AUTO_APPROVED` is allowed, which indicates that the job
	// will be run without an execution plan.
	ExecutionPlanStrategy DestroyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m DestroyJobOperationDetailsSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DestroyJobOperationDetailsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDestroyJobOperationDetailsSummary DestroyJobOperationDetailsSummary
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeDestroyJobOperationDetailsSummary
	}{
		"DESTROY",
		(MarshalTypeDestroyJobOperationDetailsSummary)(m),
	}

	return json.Marshal(&s)
}
