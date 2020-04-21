// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
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
