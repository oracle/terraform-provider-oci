// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API.
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateJobDetails Defines the requirements and properties of a job to create and run against the specified stack.
type CreateJobDetails struct {

	// OCID of the stack that is associated with the current job.
	StackId *string `mandatory:"true" json:"stackId"`

	// Terraform-specific operation to execute.
	Operation JobOperationEnum `mandatory:"true" json:"operation"`

	// Description of the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}
