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

// CreateJobDetails Defines the requirements and properties of a job to create and run against the specified stack.
type CreateJobDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack that is associated with the current job.
	StackId *string `mandatory:"true" json:"stackId"`

	// Description of the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Terraform-specific operation to execute.
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	JobOperationDetails CreateJobOperationDetails `mandatory:"false" json:"jobOperationDetails"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string                           `json:"displayName"`
		Operation              JobOperationEnum                  `json:"operation"`
		JobOperationDetails    createjoboperationdetails         `json:"jobOperationDetails"`
		ApplyJobPlanResolution *ApplyJobPlanResolution           `json:"applyJobPlanResolution"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		StackId                *string                           `json:"stackId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Operation = model.Operation

	nn, e = model.JobOperationDetails.UnmarshalPolymorphicJSON(model.JobOperationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobOperationDetails = nn.(CreateJobOperationDetails)
	} else {
		m.JobOperationDetails = nil
	}

	m.ApplyJobPlanResolution = model.ApplyJobPlanResolution

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.StackId = model.StackId

	return
}
