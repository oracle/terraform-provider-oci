// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGovernanceTargetDetails Parameters used to create a new governance target.
type CreateGovernanceTargetDetails struct {

	// Display name for the governance target.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	GovernanceScope GovernanceScope `mandatory:"true" json:"governanceScope"`

	// The governance target description.
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes to be created in Subject Tenancies
	GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipeId `mandatory:"false" json:"governanceTargetDetectorRecipes"`

	// Security Zone Recipe Id, which will be used to create the similar recipe in the subject tenancies.
	SecurityRecipeId *string `mandatory:"false" json:"securityRecipeId"`

	// List of locks associated with this resource
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateGovernanceTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGovernanceTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateGovernanceTargetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                     *string                            `json:"description"`
		GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipeId `json:"governanceTargetDetectorRecipes"`
		SecurityRecipeId                *string                            `json:"securityRecipeId"`
		Locks                           []ResourceLock                     `json:"locks"`
		FreeformTags                    map[string]string                  `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}  `json:"definedTags"`
		DisplayName                     *string                            `json:"displayName"`
		CompartmentId                   *string                            `json:"compartmentId"`
		GovernanceScope                 governancescope                    `json:"governanceScope"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.GovernanceTargetDetectorRecipes = make([]GovernanceTargetDetectorRecipeId, len(model.GovernanceTargetDetectorRecipes))
	copy(m.GovernanceTargetDetectorRecipes, model.GovernanceTargetDetectorRecipes)
	m.SecurityRecipeId = model.SecurityRecipeId

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	nn, e = model.GovernanceScope.UnmarshalPolymorphicJSON(model.GovernanceScope.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GovernanceScope = nn.(GovernanceScope)
	} else {
		m.GovernanceScope = nil
	}

	return
}
