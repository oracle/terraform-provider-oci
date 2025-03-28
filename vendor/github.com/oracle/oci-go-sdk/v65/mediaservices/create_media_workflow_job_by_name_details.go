// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMediaWorkflowJobByNameDetails Information to run a system MediaWorkflow identified by its name.
type CreateMediaWorkflowJobByNameDetails struct {

	// ID of the compartment in which the job should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Configurations to be applied to this run of the workflow.
	MediaWorkflowConfigurationIds []string `mandatory:"false" json:"mediaWorkflowConfigurationIds"`

	// Name of the Media Workflow Job. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Parameters that override parameters specified in MediaWorkflowTaskDeclarations, the MediaWorkflow,
	// the MediaWorkflow's MediaWorkflowConfigurations and the MediaWorkflowConfigurations of this
	// MediaWorkflowJob. The parameters are given as JSON. The top level and 2nd level elements must be
	// JSON objects (vs arrays, scalars, etc). The top level keys refer to a task's key and the 2nd level
	// keys refer to a parameter's name.
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Name of the system MediaWorkflow that should be run.
	MediaWorkflowName *string `mandatory:"false" json:"mediaWorkflowName"`
}

// GetMediaWorkflowConfigurationIds returns MediaWorkflowConfigurationIds
func (m CreateMediaWorkflowJobByNameDetails) GetMediaWorkflowConfigurationIds() []string {
	return m.MediaWorkflowConfigurationIds
}

// GetCompartmentId returns CompartmentId
func (m CreateMediaWorkflowJobByNameDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateMediaWorkflowJobByNameDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetParameters returns Parameters
func (m CreateMediaWorkflowJobByNameDetails) GetParameters() map[string]interface{} {
	return m.Parameters
}

// GetFreeformTags returns FreeformTags
func (m CreateMediaWorkflowJobByNameDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMediaWorkflowJobByNameDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateMediaWorkflowJobByNameDetails) GetLocks() []ResourceLock {
	return m.Locks
}

func (m CreateMediaWorkflowJobByNameDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMediaWorkflowJobByNameDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMediaWorkflowJobByNameDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMediaWorkflowJobByNameDetails CreateMediaWorkflowJobByNameDetails
	s := struct {
		DiscriminatorParam string `json:"workflowIdentifierType"`
		MarshalTypeCreateMediaWorkflowJobByNameDetails
	}{
		"NAME",
		(MarshalTypeCreateMediaWorkflowJobByNameDetails)(m),
	}

	return json.Marshal(&s)
}
