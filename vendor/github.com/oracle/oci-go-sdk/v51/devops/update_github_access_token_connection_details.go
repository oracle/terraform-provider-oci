// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v51/common"
)

// UpdateGithubAccessTokenConnectionDetails The details for updating a connection of the type `GITHUB_ACCESS_TOKEN`.
// This type corresponds to a connection in GitHub that is authenticated with a personal access token.
type UpdateGithubAccessTokenConnectionDetails struct {

	// Optional description about the Connection
	Description *string `mandatory:"false" json:"description"`

	// Optional Connection display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// OCID of personal access token saved in secret store
	AccessToken *string `mandatory:"false" json:"accessToken"`
}

//GetDescription returns Description
func (m UpdateGithubAccessTokenConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m UpdateGithubAccessTokenConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m UpdateGithubAccessTokenConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateGithubAccessTokenConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateGithubAccessTokenConnectionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateGithubAccessTokenConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateGithubAccessTokenConnectionDetails UpdateGithubAccessTokenConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateGithubAccessTokenConnectionDetails
	}{
		"GITHUB_ACCESS_TOKEN",
		(MarshalTypeUpdateGithubAccessTokenConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
