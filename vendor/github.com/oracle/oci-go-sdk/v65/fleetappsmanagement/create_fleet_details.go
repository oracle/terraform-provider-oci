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

// CreateFleetDetails The information about new Fleet.
type CreateFleetDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ResourceSelection ResourceSelection `mandatory:"true" json:"resourceSelection"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	Details FleetDetails `mandatory:"false" json:"details"`

	// Products associated with the Fleet.
	Products []string `mandatory:"false" json:"products"`

	// Environment Type associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`

	// Notification Preferences associated with the Fleet.
	NotificationPreferences []NotificationPreference `mandatory:"false" json:"notificationPreferences"`

	// Resources associated with the Fleet if resourceSelectionType is MANUAL.
	Resources []AssociatedFleetResourceDetails `mandatory:"false" json:"resources"`

	// Credentials associated with the Fleet.
	Credentials []AssociatedFleetCredentialDetails `mandatory:"false" json:"credentials"`

	// Properties associated with the Fleet.
	Properties []AssociatedFleetPropertyDetails `mandatory:"false" json:"properties"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet that would be the parent for this fleet.
	ParentFleetId *string `mandatory:"false" json:"parentFleetId"`

	// A value that represents if auto-confirming of the targets can be enabled.
	// This will allow targets to be auto-confirmed in the fleet without manual intervention.
	IsTargetAutoConfirm *bool `mandatory:"false" json:"isTargetAutoConfirm"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFleetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFleetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFleetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                            `json:"description"`
		Details                 fleetdetails                       `json:"details"`
		Products                []string                           `json:"products"`
		EnvironmentType         *string                            `json:"environmentType"`
		NotificationPreferences []NotificationPreference           `json:"notificationPreferences"`
		Resources               []AssociatedFleetResourceDetails   `json:"resources"`
		Credentials             []AssociatedFleetCredentialDetails `json:"credentials"`
		Properties              []AssociatedFleetPropertyDetails   `json:"properties"`
		ParentFleetId           *string                            `json:"parentFleetId"`
		IsTargetAutoConfirm     *bool                              `json:"isTargetAutoConfirm"`
		FreeformTags            map[string]string                  `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}  `json:"definedTags"`
		DisplayName             *string                            `json:"displayName"`
		CompartmentId           *string                            `json:"compartmentId"`
		ResourceSelection       resourceselection                  `json:"resourceSelection"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(FleetDetails)
	} else {
		m.Details = nil
	}

	m.Products = make([]string, len(model.Products))
	copy(m.Products, model.Products)
	m.EnvironmentType = model.EnvironmentType

	m.NotificationPreferences = make([]NotificationPreference, len(model.NotificationPreferences))
	copy(m.NotificationPreferences, model.NotificationPreferences)
	m.Resources = make([]AssociatedFleetResourceDetails, len(model.Resources))
	copy(m.Resources, model.Resources)
	m.Credentials = make([]AssociatedFleetCredentialDetails, len(model.Credentials))
	copy(m.Credentials, model.Credentials)
	m.Properties = make([]AssociatedFleetPropertyDetails, len(model.Properties))
	copy(m.Properties, model.Properties)
	m.ParentFleetId = model.ParentFleetId

	m.IsTargetAutoConfirm = model.IsTargetAutoConfirm

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	nn, e = model.ResourceSelection.UnmarshalPolymorphicJSON(model.ResourceSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResourceSelection = nn.(ResourceSelection)
	} else {
		m.ResourceSelection = nil
	}

	return
}
