// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PlatformConfigurationSummary Summary of the PlatformConfiguration.
type PlatformConfigurationSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Associated region
	ResourceRegion *string `mandatory:"true" json:"resourceRegion"`

	// The current state of the PlatformConfiguration.
	LifecycleState PlatformConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The type of the configuration.
	Type PlatformConfigurationTypeEnum `mandatory:"false" json:"type,omitempty"`

	ConfigCategoryDetails ConfigCategoryDetails `mandatory:"false" json:"configCategoryDetails"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PlatformConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlatformConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPlatformConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPlatformConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPlatformConfigurationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPlatformConfigurationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PlatformConfigurationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                                 `json:"description"`
		TimeUpdated           *common.SDKTime                         `json:"timeUpdated"`
		Type                  PlatformConfigurationTypeEnum           `json:"type"`
		ConfigCategoryDetails configcategorydetails                   `json:"configCategoryDetails"`
		LifecycleDetails      *string                                 `json:"lifecycleDetails"`
		FreeformTags          map[string]string                       `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}       `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}       `json:"systemTags"`
		Id                    *string                                 `json:"id"`
		CompartmentId         *string                                 `json:"compartmentId"`
		DisplayName           *string                                 `json:"displayName"`
		TimeCreated           *common.SDKTime                         `json:"timeCreated"`
		ResourceRegion        *string                                 `json:"resourceRegion"`
		LifecycleState        PlatformConfigurationLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.Type = model.Type

	nn, e = model.ConfigCategoryDetails.UnmarshalPolymorphicJSON(model.ConfigCategoryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigCategoryDetails = nn.(ConfigCategoryDetails)
	} else {
		m.ConfigCategoryDetails = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.ResourceRegion = model.ResourceRegion

	m.LifecycleState = model.LifecycleState

	return
}
