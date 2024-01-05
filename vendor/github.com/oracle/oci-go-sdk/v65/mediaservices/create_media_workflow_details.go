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

// CreateMediaWorkflowDetails The information about new MediaWorkflow.
type CreateMediaWorkflowDetails struct {

	// Name for the MediaWorkflow. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array must be unique
	// within the array. The order of tasks given here will be preserved.
	Tasks []MediaWorkflowTask `mandatory:"false" json:"tasks"`

	// Configurations to be applied to all the jobs for this workflow. Parameters in these configurations are
	// overridden by parameters in the MediaWorkflowConfigurations of the MediaWorkflowJob and the
	// parameters of the MediaWorkflowJob.
	MediaWorkflowConfigurationIds []string `mandatory:"false" json:"mediaWorkflowConfigurationIds"`

	// JSON object representing named parameters and their default values that can be referenced throughout this workflow.
	// The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating
	// MediaWorkflowJobs from this MediaWorkflow.
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMediaWorkflowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMediaWorkflowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
