// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"strings"
)

// Target Description of Target.
type Target interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// Compartment Identifier where the resource is created
	GetCompartmentId() *string

	// Resource ID which the target uses to monitor
	GetTargetResourceId() *string

	// Total number of recipes attached to target
	GetRecipeCount() *int

	// Target display name, can be renamed.
	GetDisplayName() *string

	// The target description.
	GetDescription() *string

	// List of detector recipes associated with target
	GetTargetDetectorRecipes() []TargetDetectorRecipe

	// List of responder recipes associated with target
	GetTargetResponderRecipes() []TargetResponderRecipe

	// List of inherited compartments
	GetInheritedByCompartments() []string

	// The date and time the target was created. Format defined by RFC3339.
	GetTimeCreated() *common.SDKTime

	// The date and time the target was updated. Format defined by RFC3339.
	GetTimeUpdated() *common.SDKTime

	// The current state of the Target.
	GetLifecycleState() LifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecyleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type target struct {
	JsonData                []byte
	Id                      *string                           `mandatory:"true" json:"id"`
	CompartmentId           *string                           `mandatory:"true" json:"compartmentId"`
	TargetResourceId        *string                           `mandatory:"true" json:"targetResourceId"`
	RecipeCount             *int                              `mandatory:"true" json:"recipeCount"`
	DisplayName             *string                           `mandatory:"false" json:"displayName"`
	Description             *string                           `mandatory:"false" json:"description"`
	TargetDetectorRecipes   []TargetDetectorRecipe            `mandatory:"false" json:"targetDetectorRecipes"`
	TargetResponderRecipes  []TargetResponderRecipe           `mandatory:"false" json:"targetResponderRecipes"`
	InheritedByCompartments []string                          `mandatory:"false" json:"inheritedByCompartments"`
	TimeCreated             *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated             *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState          LifecycleStateEnum                `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecyleDetails         *string                           `mandatory:"false" json:"lifecyleDetails"`
	FreeformTags            map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags             map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags              map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TargetResourceType      string                            `json:"targetResourceType"`
}

// UnmarshalJSON unmarshals json
func (m *target) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertarget target
	s := struct {
		Model Unmarshalertarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.TargetResourceId = s.Model.TargetResourceId
	m.RecipeCount = s.Model.RecipeCount
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.TargetDetectorRecipes = s.Model.TargetDetectorRecipes
	m.TargetResponderRecipes = s.Model.TargetResponderRecipes
	m.InheritedByCompartments = s.Model.InheritedByCompartments
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecyleDetails = s.Model.LifecyleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TargetResourceType = s.Model.TargetResourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *target) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetResourceType {
	case "COMPARTMENT":
		mm := OciTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECURITY_ZONE":
		mm := SzTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m target) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m target) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTargetResourceId returns TargetResourceId
func (m target) GetTargetResourceId() *string {
	return m.TargetResourceId
}

//GetRecipeCount returns RecipeCount
func (m target) GetRecipeCount() *int {
	return m.RecipeCount
}

//GetDisplayName returns DisplayName
func (m target) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m target) GetDescription() *string {
	return m.Description
}

//GetTargetDetectorRecipes returns TargetDetectorRecipes
func (m target) GetTargetDetectorRecipes() []TargetDetectorRecipe {
	return m.TargetDetectorRecipes
}

//GetTargetResponderRecipes returns TargetResponderRecipes
func (m target) GetTargetResponderRecipes() []TargetResponderRecipe {
	return m.TargetResponderRecipes
}

//GetInheritedByCompartments returns InheritedByCompartments
func (m target) GetInheritedByCompartments() []string {
	return m.InheritedByCompartments
}

//GetTimeCreated returns TimeCreated
func (m target) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m target) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m target) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m target) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetFreeformTags returns FreeformTags
func (m target) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m target) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m target) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m target) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m target) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
