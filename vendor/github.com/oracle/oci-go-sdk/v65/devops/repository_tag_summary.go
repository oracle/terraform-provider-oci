// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RepositoryTagSummary The information needed to create a lightweight tag.
type RepositoryTagSummary struct {

	// Reference name inside a repository.
	RefName *string `mandatory:"true" json:"refName"`

	// Unique full reference name inside a repository.
	FullRefName *string `mandatory:"true" json:"fullRefName"`

	// The OCID of the repository containing the reference.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// SHA-1 hash value of the object pointed to by the tag.
	ObjectId *string `mandatory:"true" json:"objectId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetRefName returns RefName
func (m RepositoryTagSummary) GetRefName() *string {
	return m.RefName
}

// GetFullRefName returns FullRefName
func (m RepositoryTagSummary) GetFullRefName() *string {
	return m.FullRefName
}

// GetRepositoryId returns RepositoryId
func (m RepositoryTagSummary) GetRepositoryId() *string {
	return m.RepositoryId
}

// GetFreeformTags returns FreeformTags
func (m RepositoryTagSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m RepositoryTagSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m RepositoryTagSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryTagSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RepositoryTagSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRepositoryTagSummary RepositoryTagSummary
	s := struct {
		DiscriminatorParam string `json:"refType"`
		MarshalTypeRepositoryTagSummary
	}{
		"TAG",
		(MarshalTypeRepositoryTagSummary)(m),
	}

	return json.Marshal(&s)
}
