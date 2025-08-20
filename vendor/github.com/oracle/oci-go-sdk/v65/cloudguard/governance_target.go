// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceTarget A governance target defines the scope of resources that Cloud Guard monitors and the rules to be enforced in that monitoring for assigned Subject Tenancies.
type GovernanceTarget struct {

	// Unique identifier that can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	GovernanceScope GovernanceScope `mandatory:"true" json:"governanceScope"`

	// Total number of recipes attached to target
	RecipeCount *int `mandatory:"true" json:"recipeCount"`

	// The date and time the governance target was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the governance target
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Governance Target display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The governance target description
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes to be created in Subject Tenancies
	GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipe `mandatory:"false" json:"governanceTargetDetectorRecipes"`

	// The date and time the governance target was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current lifecycle state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GovernanceTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GovernanceTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *GovernanceTarget) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                     *string                           `json:"displayName"`
		Description                     *string                           `json:"description"`
		GovernanceTargetDetectorRecipes []GovernanceTargetDetectorRecipe  `json:"governanceTargetDetectorRecipes"`
		TimeUpdated                     *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails                *string                           `json:"lifecycleDetails"`
		Locks                           []ResourceLock                    `json:"locks"`
		FreeformTags                    map[string]string                 `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                      map[string]map[string]interface{} `json:"systemTags"`
		Id                              *string                           `json:"id"`
		CompartmentId                   *string                           `json:"compartmentId"`
		GovernanceScope                 governancescope                   `json:"governanceScope"`
		RecipeCount                     *int                              `json:"recipeCount"`
		TimeCreated                     *common.SDKTime                   `json:"timeCreated"`
		LifecycleState                  LifecycleStateEnum                `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.GovernanceTargetDetectorRecipes = make([]GovernanceTargetDetectorRecipe, len(model.GovernanceTargetDetectorRecipes))
	copy(m.GovernanceTargetDetectorRecipes, model.GovernanceTargetDetectorRecipes)
	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

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

	m.RecipeCount = model.RecipeCount

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
