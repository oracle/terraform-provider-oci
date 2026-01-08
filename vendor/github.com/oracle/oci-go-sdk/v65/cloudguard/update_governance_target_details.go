// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateGovernanceTargetDetails The governance target information to be updated.
type UpdateGovernanceTargetDetails struct {

	// Display name of a governance target.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The governance target description.
	Description *string `mandatory:"false" json:"description"`

	GovernanceScope GovernanceScope `mandatory:"false" json:"governanceScope"`

	// List of detector recipes to be created in Subject Tenancies
	GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipeId `mandatory:"false" json:"governanceTargetDetectorRecipes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateGovernanceTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGovernanceTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateGovernanceTargetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                     *string                            `json:"displayName"`
		Description                     *string                            `json:"description"`
		GovernanceScope                 governancescope                    `json:"governanceScope"`
		GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipeId `json:"governanceTargetDetectorRecipes"`
		FreeformTags                    map[string]string                  `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}  `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.GovernanceScope.UnmarshalPolymorphicJSON(model.GovernanceScope.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GovernanceScope = nn.(GovernanceScope)
	} else {
		m.GovernanceScope = nil
	}

	m.GovernanceTargetDetectorRecipes = make([]GovernanceTargetDetectorRecipeId, len(model.GovernanceTargetDetectorRecipes))
	copy(m.GovernanceTargetDetectorRecipes, model.GovernanceTargetDetectorRecipes)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
