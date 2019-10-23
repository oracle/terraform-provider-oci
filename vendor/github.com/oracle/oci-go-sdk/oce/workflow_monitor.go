// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkflowMonitor The workflow monitor for this work request.
type WorkflowMonitor struct {

	// workflow name for this work request
	WorkflowName *string `mandatory:"false" json:"workflowName"`

	// resource name for this work request
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Workflow step of workflow monitor.
	WorkflowSteps []WorkflowStep `mandatory:"false" json:"workflowSteps"`
}

func (m WorkflowMonitor) String() string {
	return common.PointerString(m)
}
