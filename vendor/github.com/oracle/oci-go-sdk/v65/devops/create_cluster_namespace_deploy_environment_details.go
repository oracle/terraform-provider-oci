// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateClusterNamespaceDeployEnvironmentDetails Specifies the OMK cluster namespace environment.
type CreateClusterNamespaceDeployEnvironmentDetails struct {

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the Cluster Namespace to deploy to.
	ClusterNamespaceId *string `mandatory:"true" json:"clusterNamespaceId"`

	// Optional description about the deployment environment.
	Description *string `mandatory:"false" json:"description"`

	// Deployment environment display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	NetworkChannel NetworkChannel `mandatory:"false" json:"networkChannel"`
}

//GetDescription returns Description
func (m CreateClusterNamespaceDeployEnvironmentDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateClusterNamespaceDeployEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m CreateClusterNamespaceDeployEnvironmentDetails) GetProjectId() *string {
	return m.ProjectId
}

//GetFreeformTags returns FreeformTags
func (m CreateClusterNamespaceDeployEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateClusterNamespaceDeployEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateClusterNamespaceDeployEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateClusterNamespaceDeployEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateClusterNamespaceDeployEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateClusterNamespaceDeployEnvironmentDetails CreateClusterNamespaceDeployEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"deployEnvironmentType"`
		MarshalTypeCreateClusterNamespaceDeployEnvironmentDetails
	}{
		"CLUSTER_NAMESPACE",
		(MarshalTypeCreateClusterNamespaceDeployEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateClusterNamespaceDeployEnvironmentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                           `json:"description"`
		DisplayName        *string                           `json:"displayName"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		NetworkChannel     networkchannel                    `json:"networkChannel"`
		ProjectId          *string                           `json:"projectId"`
		ClusterNamespaceId *string                           `json:"clusterNamespaceId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.NetworkChannel.UnmarshalPolymorphicJSON(model.NetworkChannel.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkChannel = nn.(NetworkChannel)
	} else {
		m.NetworkChannel = nil
	}

	m.ProjectId = model.ProjectId

	m.ClusterNamespaceId = model.ClusterNamespaceId

	return
}
