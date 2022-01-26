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

// CreateGenericRepositoryDetails Parameters needed to create an artifact repository.
type CreateGenericRepositoryDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the repository's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether to make the repository immutable. The artifacts of an immutable repository cannot be overwritten.
	IsImmutable *bool `mandatory:"true" json:"isImmutable"`

	// A user-friendly display name for the repository. If not present, will be auto-generated. It can be modified later. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the repository. It can be updated later.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateGenericRepositoryDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateGenericRepositoryDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m CreateGenericRepositoryDetails) GetDescription() *string {
	return m.Description
}

//GetIsImmutable returns IsImmutable
func (m CreateGenericRepositoryDetails) GetIsImmutable() *bool {
	return m.IsImmutable
}

//GetFreeformTags returns FreeformTags
func (m CreateGenericRepositoryDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateGenericRepositoryDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateGenericRepositoryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateGenericRepositoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGenericRepositoryDetails CreateGenericRepositoryDetails
	s := struct {
		DiscriminatorParam string `json:"repositoryType"`
		MarshalTypeCreateGenericRepositoryDetails
	}{
		"GENERIC",
		(MarshalTypeCreateGenericRepositoryDetails)(m),
	}

	return json.Marshal(&s)
}
