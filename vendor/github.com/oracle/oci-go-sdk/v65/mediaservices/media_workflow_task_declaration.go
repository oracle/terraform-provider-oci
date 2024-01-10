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

// MediaWorkflowTaskDeclaration The declaration of a type of task that can be used in a MediaWorkflow.
type MediaWorkflowTaskDeclaration struct {

	// MediaWorkflowTaskDeclaration identifier. The name and version should be unique among
	// MediaWorkflowTaskDeclarations.
	Name *string `mandatory:"true" json:"name"`

	// The version of MediaWorkflowTaskDeclaration, incremented whenever the team implementing the task processor
	// modifies the JSON schema of this declaration's definitions, parameters or list of required parameters.
	Version *int `mandatory:"true" json:"version"`

	// JSON schema specifying the parameters supported by this type of task. This is used to validate tasks'
	// parameters when jobs are created.
	ParametersSchema map[string]interface{} `mandatory:"true" json:"parametersSchema"`

	// JSON schema similar to the parameterSchema, but permits parameter values to refer to other parameters using the
	// ${/path/to/another/parmeter} syntax.  This is used to validate task parameters when workflows are created.
	ParametersSchemaAllowingReferences map[string]interface{} `mandatory:"true" json:"parametersSchemaAllowingReferences"`
}

func (m MediaWorkflowTaskDeclaration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflowTaskDeclaration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
