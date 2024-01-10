// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// API covering the Artifacts and Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as generic artifacts and container images.
//

package artifacts

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RepositorySummary Summary information for a repository.
type RepositorySummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	// Example: `ocid1.artifactrepository.oc1..exampleuniqueID`
	GetId() *string

	// The repository name.
	GetDisplayName() *string

	// The OCID of the repository's compartment.
	GetCompartmentId() *string

	// Whether the repository is immutable. The artifacts of an immutable repository cannot be overwritten.
	GetIsImmutable() *bool

	// The current state of the artifact repository.
	GetLifecycleState() RepositoryLifecycleStateEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// An RFC 3339 timestamp indicating when the repository was created.
	GetTimeCreated() *common.SDKTime

	// The repository description.
	GetDescription() *string
}

type repositorysummary struct {
	JsonData       []byte
	Description    *string                           `mandatory:"false" json:"description"`
	Id             *string                           `mandatory:"true" json:"id"`
	DisplayName    *string                           `mandatory:"true" json:"displayName"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	IsImmutable    *bool                             `mandatory:"true" json:"isImmutable"`
	LifecycleState RepositoryLifecycleStateEnum      `mandatory:"true" json:"lifecycleState"`
	FreeformTags   map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	TimeCreated    *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	RepositoryType string                            `json:"repositoryType"`
}

// UnmarshalJSON unmarshals json
func (m *repositorysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrepositorysummary repositorysummary
	s := struct {
		Model Unmarshalerrepositorysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.IsImmutable = s.Model.IsImmutable
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TimeCreated = s.Model.TimeCreated
	m.Description = s.Model.Description
	m.RepositoryType = s.Model.RepositoryType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *repositorysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RepositoryType {
	case "GENERIC":
		mm := GenericRepositorySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RepositorySummary: %s.", m.RepositoryType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m repositorysummary) GetDescription() *string {
	return m.Description
}

// GetId returns Id
func (m repositorysummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m repositorysummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m repositorysummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetIsImmutable returns IsImmutable
func (m repositorysummary) GetIsImmutable() *bool {
	return m.IsImmutable
}

// GetLifecycleState returns LifecycleState
func (m repositorysummary) GetLifecycleState() RepositoryLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m repositorysummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m repositorysummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetTimeCreated returns TimeCreated
func (m repositorysummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m repositorysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m repositorysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRepositoryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
