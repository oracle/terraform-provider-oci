// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Fleet A fleet is a collection or grouping of resources based on criteria.
type Fleet struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the Fleet.
	LifecycleState FleetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Products associated with the Fleet.
	Products []string `mandatory:"false" json:"products"`

	Details FleetDetails `mandatory:"false" json:"details"`

	// Environment Type associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	ResourceSelection ResourceSelection `mandatory:"false" json:"resourceSelection"`

	// Notification Preferences associated with the Fleet.
	NotificationPreferences []NotificationPreference `mandatory:"false" json:"notificationPreferences"`

	// Resources associated with the Fleet if resourceSelectionType is MANUAL.
	Resources []AssociatedFleetResourceDetails `mandatory:"false" json:"resources"`

	// Properties associated with the Fleet.
	Properties []AssociatedFleetPropertyDetails `mandatory:"false" json:"properties"`

	// Credentials associated with the Fleet.
	Credentials []AssociatedFleetCredentialDetails `mandatory:"false" json:"credentials"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet that would be the parent for this fleet.
	ParentFleetId *string `mandatory:"false" json:"parentFleetId"`

	// A value that represents if auto-confirming of the targets can be enabled.
	// This will allow targets to be auto-confirmed in the fleet without manual intervention.
	IsTargetAutoConfirm *bool `mandatory:"false" json:"isTargetAutoConfirm"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Fleet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Fleet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Fleet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceRegion          *string                            `json:"resourceRegion"`
		Description             *string                            `json:"description"`
		TimeUpdated             *common.SDKTime                    `json:"timeUpdated"`
		Products                []string                           `json:"products"`
		Details                 fleetdetails                       `json:"details"`
		EnvironmentType         *string                            `json:"environmentType"`
		ResourceSelection       resourceselection                  `json:"resourceSelection"`
		NotificationPreferences []NotificationPreference           `json:"notificationPreferences"`
		Resources               []AssociatedFleetResourceDetails   `json:"resources"`
		Properties              []AssociatedFleetPropertyDetails   `json:"properties"`
		Credentials             []AssociatedFleetCredentialDetails `json:"credentials"`
		ParentFleetId           *string                            `json:"parentFleetId"`
		IsTargetAutoConfirm     *bool                              `json:"isTargetAutoConfirm"`
		LifecycleDetails        *string                            `json:"lifecycleDetails"`
		SystemTags              map[string]map[string]interface{}  `json:"systemTags"`
		Id                      *string                            `json:"id"`
		CompartmentId           *string                            `json:"compartmentId"`
		DisplayName             *string                            `json:"displayName"`
		TimeCreated             *common.SDKTime                    `json:"timeCreated"`
		LifecycleState          FleetLifecycleStateEnum            `json:"lifecycleState"`
		FreeformTags            map[string]string                  `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}  `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResourceRegion = model.ResourceRegion

	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.Products = make([]string, len(model.Products))
	copy(m.Products, model.Products)
	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(FleetDetails)
	} else {
		m.Details = nil
	}

	m.EnvironmentType = model.EnvironmentType

	nn, e = model.ResourceSelection.UnmarshalPolymorphicJSON(model.ResourceSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResourceSelection = nn.(ResourceSelection)
	} else {
		m.ResourceSelection = nil
	}

	m.NotificationPreferences = make([]NotificationPreference, len(model.NotificationPreferences))
	copy(m.NotificationPreferences, model.NotificationPreferences)
	m.Resources = make([]AssociatedFleetResourceDetails, len(model.Resources))
	copy(m.Resources, model.Resources)
	m.Properties = make([]AssociatedFleetPropertyDetails, len(model.Properties))
	copy(m.Properties, model.Properties)
	m.Credentials = make([]AssociatedFleetCredentialDetails, len(model.Credentials))
	copy(m.Credentials, model.Credentials)
	m.ParentFleetId = model.ParentFleetId

	m.IsTargetAutoConfirm = model.IsTargetAutoConfirm

	m.LifecycleDetails = model.LifecycleDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// FleetLifecycleStateEnum Enum with underlying type: string
type FleetLifecycleStateEnum string

// Set of constants representing the allowable values for FleetLifecycleStateEnum
const (
	FleetLifecycleStateActive         FleetLifecycleStateEnum = "ACTIVE"
	FleetLifecycleStateInactive       FleetLifecycleStateEnum = "INACTIVE"
	FleetLifecycleStateCreating       FleetLifecycleStateEnum = "CREATING"
	FleetLifecycleStateDeleted        FleetLifecycleStateEnum = "DELETED"
	FleetLifecycleStateDeleting       FleetLifecycleStateEnum = "DELETING"
	FleetLifecycleStateFailed         FleetLifecycleStateEnum = "FAILED"
	FleetLifecycleStateUpdating       FleetLifecycleStateEnum = "UPDATING"
	FleetLifecycleStateNeedsAttention FleetLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingFleetLifecycleStateEnum = map[string]FleetLifecycleStateEnum{
	"ACTIVE":          FleetLifecycleStateActive,
	"INACTIVE":        FleetLifecycleStateInactive,
	"CREATING":        FleetLifecycleStateCreating,
	"DELETED":         FleetLifecycleStateDeleted,
	"DELETING":        FleetLifecycleStateDeleting,
	"FAILED":          FleetLifecycleStateFailed,
	"UPDATING":        FleetLifecycleStateUpdating,
	"NEEDS_ATTENTION": FleetLifecycleStateNeedsAttention,
}

var mappingFleetLifecycleStateEnumLowerCase = map[string]FleetLifecycleStateEnum{
	"active":          FleetLifecycleStateActive,
	"inactive":        FleetLifecycleStateInactive,
	"creating":        FleetLifecycleStateCreating,
	"deleted":         FleetLifecycleStateDeleted,
	"deleting":        FleetLifecycleStateDeleting,
	"failed":          FleetLifecycleStateFailed,
	"updating":        FleetLifecycleStateUpdating,
	"needs_attention": FleetLifecycleStateNeedsAttention,
}

// GetFleetLifecycleStateEnumValues Enumerates the set of values for FleetLifecycleStateEnum
func GetFleetLifecycleStateEnumValues() []FleetLifecycleStateEnum {
	values := make([]FleetLifecycleStateEnum, 0)
	for _, v := range mappingFleetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetLifecycleStateEnumStringValues Enumerates the set of values in String for FleetLifecycleStateEnum
func GetFleetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingFleetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetLifecycleStateEnum(val string) (FleetLifecycleStateEnum, bool) {
	enum, ok := mappingFleetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
