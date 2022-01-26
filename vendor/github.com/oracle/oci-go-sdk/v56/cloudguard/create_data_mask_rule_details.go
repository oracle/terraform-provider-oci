// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateDataMaskRuleDetails The information about new Data Mask Rule.
type CreateDataMaskRuleDetails struct {

	// Data Mask Rule name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// IAM Group id associated with the data mask rule
	IamGroupId *string `mandatory:"true" json:"iamGroupId"`

	TargetSelected TargetSelected `mandatory:"true" json:"targetSelected"`

	// Data Mask Categories
	DataMaskCategories []DataMaskCategoryEnum `mandatory:"true" json:"dataMaskCategories"`

	// The Data Mask Rule description.
	Description *string `mandatory:"false" json:"description"`

	// The status of the dataMaskRule.
	DataMaskRuleStatus DataMaskRuleStatusEnum `mandatory:"false" json:"dataMaskRuleStatus,omitempty"`

	// The current state of the DataMaskRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDataMaskRuleDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataMaskRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                           `json:"description"`
		DataMaskRuleStatus DataMaskRuleStatusEnum            `json:"dataMaskRuleStatus"`
		LifecycleState     LifecycleStateEnum                `json:"lifecycleState"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		CompartmentId      *string                           `json:"compartmentId"`
		IamGroupId         *string                           `json:"iamGroupId"`
		TargetSelected     targetselected                    `json:"targetSelected"`
		DataMaskCategories []DataMaskCategoryEnum            `json:"dataMaskCategories"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DataMaskRuleStatus = model.DataMaskRuleStatus

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

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

	m.DataMaskCategories = make([]DataMaskCategoryEnum, len(model.DataMaskCategories))
	for i, n := range model.DataMaskCategories {
		m.DataMaskCategories[i] = n
	}

	return
}
