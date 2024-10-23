// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkflowTaskComponent Workflow Task Component Details.
type WorkflowTaskComponent struct {

	// Provide StepName for the Task.
	StepName *string `mandatory:"true" json:"stepName"`
}

func (m WorkflowTaskComponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkflowTaskComponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WorkflowTaskComponent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWorkflowTaskComponent WorkflowTaskComponent
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeWorkflowTaskComponent
	}{
		"TASK",
		(MarshalTypeWorkflowTaskComponent)(m),
	}

	return json.Marshal(&s)
}
