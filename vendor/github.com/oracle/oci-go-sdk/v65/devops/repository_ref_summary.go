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

// RepositoryRefSummary Summary of a reference.
type RepositoryRefSummary interface {

	// Reference name inside a repository.
	GetRefName() *string

	// Unique full reference name inside a repository.
	GetFullRefName() *string

	// The OCID of the repository containing the reference.
	GetRepositoryId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type repositoryrefsummary struct {
	JsonData     []byte
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags   map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	RefName      *string                           `mandatory:"true" json:"refName"`
	FullRefName  *string                           `mandatory:"true" json:"fullRefName"`
	RepositoryId *string                           `mandatory:"true" json:"repositoryId"`
	RefType      string                            `json:"refType"`
}

// UnmarshalJSON unmarshals json
func (m *repositoryrefsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrepositoryrefsummary repositoryrefsummary
	s := struct {
		Model Unmarshalerrepositoryrefsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RefName = s.Model.RefName
	m.FullRefName = s.Model.FullRefName
	m.RepositoryId = s.Model.RepositoryId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.RefType = s.Model.RefType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *repositoryrefsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RefType {
	case "BRANCH":
		mm := RepositoryBranchSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAG":
		mm := RepositoryTagSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RepositoryRefSummary: %s.", m.RefType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m repositoryrefsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m repositoryrefsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m repositoryrefsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetRefName returns RefName
func (m repositoryrefsummary) GetRefName() *string {
	return m.RefName
}

// GetFullRefName returns FullRefName
func (m repositoryrefsummary) GetFullRefName() *string {
	return m.FullRefName
}

// GetRepositoryId returns RepositoryId
func (m repositoryrefsummary) GetRepositoryId() *string {
	return m.RepositoryId
}

func (m repositoryrefsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m repositoryrefsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
