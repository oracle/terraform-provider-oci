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

// UpdateFleetDetails The information to be updated.
type UpdateFleetDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Notification Preferences associated with the Fleet.
	// An UPDATE operation replaces the existing notification preferences list entirely
	NotificationPreferences []NotificationPreference `mandatory:"false" json:"notificationPreferences"`

	// A value that represents if auto-confirming of the targets can be enabled.
	// This will allow targets to be auto-confirmed in the fleet without manual intervention.
	IsTargetAutoConfirm *bool `mandatory:"false" json:"isTargetAutoConfirm"`

	ResourceSelection ResourceSelection `mandatory:"false" json:"resourceSelection"`

	// Products associated with the Fleet.
	// Provide PlatformConfiguration Ids corresponding to all the Products that need to be managed.
	Products []string `mandatory:"false" json:"products"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Environment Type associated with the Fleet.
	// Applicable for ENVIRONMENT fleet types.
	EnvironmentType *string `mandatory:"false" json:"environmentType"`
}

func (m UpdateFleetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFleetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateFleetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		Description             *string                           `json:"description"`
		NotificationPreferences []NotificationPreference          `json:"notificationPreferences"`
		IsTargetAutoConfirm     *bool                             `json:"isTargetAutoConfirm"`
		ResourceSelection       resourceselection                 `json:"resourceSelection"`
		Products                []string                          `json:"products"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		EnvironmentType         *string                           `json:"environmentType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.NotificationPreferences = make([]NotificationPreference, len(model.NotificationPreferences))
	copy(m.NotificationPreferences, model.NotificationPreferences)
	m.IsTargetAutoConfirm = model.IsTargetAutoConfirm

	nn, e = model.ResourceSelection.UnmarshalPolymorphicJSON(model.ResourceSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResourceSelection = nn.(ResourceSelection)
	} else {
		m.ResourceSelection = nil
	}

	m.Products = make([]string, len(model.Products))
	copy(m.Products, model.Products)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.EnvironmentType = model.EnvironmentType

	return
}
