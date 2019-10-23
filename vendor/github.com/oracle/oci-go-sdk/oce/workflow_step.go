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

// WorkflowStep Workflow step of workflow monitor.
type WorkflowStep struct {

	// workflow step name
	StepName *string `mandatory:"false" json:"stepName"`

	// workflow step status
	Status *string `mandatory:"false" json:"status"`
}

func (m WorkflowStep) String() string {
	return common.PointerString(m)
}
