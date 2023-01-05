// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetSummary Summary of the Target.
type TargetSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// possible type of targets(compartment/HCMCloud/ERPCloud)
	TargetResourceType TargetResourceTypeEnum `mandatory:"true" json:"targetResourceType"`

	// Resource ID which the target uses to monitor
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// Total number of recipes attached to target
	RecipeCount *int `mandatory:"true" json:"recipeCount"`

	// DetectorTemplate Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	TargetSummaryAdditionalDetails TargetSummaryAdditionalDetails `mandatory:"false" json:"targetSummaryAdditionalDetails"`

	// The date and time the target was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TargetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetResourceTypeEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetTargetResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TargetSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                    *string                           `json:"displayName"`
		TargetSummaryAdditionalDetails targetsummaryadditionaldetails    `json:"targetSummaryAdditionalDetails"`
		TimeCreated                    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                    *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState                 LifecycleStateEnum                `json:"lifecycleState"`
		LifecyleDetails                *string                           `json:"lifecyleDetails"`
		FreeformTags                   map[string]string                 `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                     map[string]map[string]interface{} `json:"systemTags"`
		Id                             *string                           `json:"id"`
		CompartmentId                  *string                           `json:"compartmentId"`
		TargetResourceType             TargetResourceTypeEnum            `json:"targetResourceType"`
		TargetResourceId               *string                           `json:"targetResourceId"`
		RecipeCount                    *int                              `json:"recipeCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.TargetSummaryAdditionalDetails.UnmarshalPolymorphicJSON(model.TargetSummaryAdditionalDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetSummaryAdditionalDetails = nn.(TargetSummaryAdditionalDetails)
	} else {
		m.TargetSummaryAdditionalDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecyleDetails = model.LifecyleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TargetResourceType = model.TargetResourceType

	m.TargetResourceId = model.TargetResourceId

	m.RecipeCount = model.RecipeCount

	return
}
