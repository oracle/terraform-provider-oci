// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DataMaskRule Description of DataMaskRule.
type DataMaskRule struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// IAM Group id associated with the data mask rule
	IamGroupId *string `mandatory:"true" json:"iamGroupId"`

	TargetSelected TargetSelected `mandatory:"true" json:"targetSelected"`

	// Data Mask Rule Identifier, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The data mask rule description.
	Description *string `mandatory:"false" json:"description"`

	// Data Mask Categories
	DataMaskCategories []DataMaskCategoryEnum `mandatory:"false" json:"dataMaskCategories,omitempty"`

	// The date and time the target was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The status of the dataMaskRule.
	DataMaskRuleStatus DataMaskRuleStatusEnum `mandatory:"false" json:"dataMaskRuleStatus,omitempty"`

	// The current state of the DataMaskRule.
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

func (m DataMaskRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataMaskRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.DataMaskCategories {
		if _, ok := GetMappingDataMaskCategoryEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataMaskCategories: %s. Supported values are: %s.", val, strings.Join(GetDataMaskCategoryEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingDataMaskRuleStatusEnum(string(m.DataMaskRuleStatus)); !ok && m.DataMaskRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataMaskRuleStatus: %s. Supported values are: %s.", m.DataMaskRuleStatus, strings.Join(GetDataMaskRuleStatusEnumStringValues(), ",")))
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
func (m *DataMaskRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		Description        *string                           `json:"description"`
		DataMaskCategories []DataMaskCategoryEnum            `json:"dataMaskCategories"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
		DataMaskRuleStatus DataMaskRuleStatusEnum            `json:"dataMaskRuleStatus"`
		LifecycleState     LifecycleStateEnum                `json:"lifecycleState"`
		LifecyleDetails    *string                           `json:"lifecyleDetails"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		Id                 *string                           `json:"id"`
		CompartmentId      *string                           `json:"compartmentId"`
		IamGroupId         *string                           `json:"iamGroupId"`
		TargetSelected     targetselected                    `json:"targetSelected"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.DataMaskCategories = make([]DataMaskCategoryEnum, len(model.DataMaskCategories))
	copy(m.DataMaskCategories, model.DataMaskCategories)
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.DataMaskRuleStatus = model.DataMaskRuleStatus

	m.LifecycleState = model.LifecycleState

	m.LifecyleDetails = model.LifecyleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.IamGroupId = model.IamGroupId

	nn, e = model.TargetSelected.UnmarshalPolymorphicJSON(model.TargetSelected.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetSelected = nn.(TargetSelected)
	} else {
		m.TargetSelected = nil
	}

	return
}
