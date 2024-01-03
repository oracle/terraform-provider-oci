// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content Management API
//
// Oracle Content Management is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkflowMonitor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
