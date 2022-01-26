// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateFunctionDeployEnvironmentDetails Specifies the Function environment.
type CreateFunctionDeployEnvironmentDetails struct {

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the Function.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// Optional description about the deployment environment.
	Description *string `mandatory:"false" json:"description"`

	// Deployment environment display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDescription returns Description
func (m CreateFunctionDeployEnvironmentDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateFunctionDeployEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m CreateFunctionDeployEnvironmentDetails) GetProjectId() *string {
	return m.ProjectId
}

//GetFreeformTags returns FreeformTags
func (m CreateFunctionDeployEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateFunctionDeployEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateFunctionDeployEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateFunctionDeployEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateFunctionDeployEnvironmentDetails CreateFunctionDeployEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"deployEnvironmentType"`
		MarshalTypeCreateFunctionDeployEnvironmentDetails
	}{
		"FUNCTION",
		(MarshalTypeCreateFunctionDeployEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}
