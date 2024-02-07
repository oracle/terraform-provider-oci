// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MediaWorkflowTask Defines the type of processing to be run at a given point in the workflow, parameters to configure the
// processing, and any processing that must be completed before this processing begins.
type MediaWorkflowTask struct {

	// The type of process to run at this task. Refers to the name of a MediaWorkflowTaskDeclaration.
	Type *string `mandatory:"true" json:"type"`

	// The version of the MediaWorkflowTaskDeclaration.
	Version *int64 `mandatory:"true" json:"version"`

	// A unique identifier for this task within its workflow. Keys are used to reference a task within workflows
	// and MediaWorkflowJobs. Tasks are referenced as prerequisites and to track output and state.
	Key *string `mandatory:"true" json:"key"`

	// Data specifiying how this task is to be run. The data is a JSON object that must conform to the JSON Schema
	// specified by the parameters of the MediaWorkflowTaskDeclaration this task references. The parameters may
	// contain values or references to other parameters.
	Parameters map[string]interface{} `mandatory:"true" json:"parameters"`

	// Keys to the other tasks in this workflow that must be completed before execution of this task can begin.
	Prerequisites []string `mandatory:"false" json:"prerequisites"`

	// Allows this task to be conditionally enabled.  If no value or a blank value is given, the task is
	// unconditionally enbled.  Otherwise the given string specifies a parameter of the job created for this task's
	// workflow using the JSON pointer syntax. The JSON pointer is validated when a job is created from the workflow of this task.
	EnableParameterReference *string `mandatory:"false" json:"enableParameterReference"`

	// Used in conjunction with enableParameterReference to conditionally enable a task.  When a job is created
	// from the workflow of this task, the task will only be enabled if the value of the parameter specified by
	// enableParameterReference is equal to the value of this property. This property must be prenset if and only if
	// a enableParameterReference is given. The value is a JSON node.
	EnableWhenReferencedParameterEquals map[string]interface{} `mandatory:"false" json:"enableWhenReferencedParameterEquals"`
}

func (m MediaWorkflowTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflowTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
