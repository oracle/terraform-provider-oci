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

// SystemMediaWorkflow A named list of tasks to be used to run a job or as a template to create a MediaWorkflow.
type SystemMediaWorkflow struct {

	// System provided unique identifier for this static media workflow.
	Name *string `mandatory:"true" json:"name"`

	// The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array is unique
	// within the array. The order of the items is preserved from the order of the tasks array in
	// CreateMediaWorkflowDetails or UpdateMediaWorkflowDetails.
	Tasks []MediaWorkflowTask `mandatory:"true" json:"tasks"`

	// Description of this workflow's processing and how that processing can be customized by
	// specifying parameter values.
	Description *string `mandatory:"false" json:"description"`

	// JSON object representing named parameters and their default values that can be referenced throughout this workflow.
	// The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating
	// MediaWorkflowJobs from this MediaWorkflow.
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m SystemMediaWorkflow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SystemMediaWorkflow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
