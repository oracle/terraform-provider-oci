// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// GenericRepository The metadata for the artifact repository.
type GenericRepository struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	// Example: `ocid1.artifactrepository.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The repository name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the repository's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The repository description.
	Description *string `mandatory:"true" json:"description"`

	// Whether the repository is immutable. The artifacts of an immutable repository cannot be overwritten.
	IsImmutable *bool `mandatory:"true" json:"isImmutable"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the repository.
	LifecycleState RepositoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m GenericRepository) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m GenericRepository) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m GenericRepository) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m GenericRepository) GetDescription() *string {
	return m.Description
}

//GetIsImmutable returns IsImmutable
func (m GenericRepository) GetIsImmutable() *bool {
	return m.IsImmutable
}

//GetLifecycleState returns LifecycleState
func (m GenericRepository) GetLifecycleState() RepositoryLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m GenericRepository) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GenericRepository) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetTimeCreated returns TimeCreated
func (m GenericRepository) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m GenericRepository) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GenericRepository) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericRepository GenericRepository
	s := struct {
		DiscriminatorParam string `json:"repositoryType"`
		MarshalTypeGenericRepository
	}{
		"GENERIC",
		(MarshalTypeGenericRepository)(m),
	}

	return json.Marshal(&s)
}
