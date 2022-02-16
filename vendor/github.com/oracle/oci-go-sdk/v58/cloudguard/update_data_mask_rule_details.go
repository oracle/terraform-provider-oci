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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateDataMaskRuleDetails The information to be updated.
type UpdateDataMaskRuleDetails struct {

	// Data Mask Rule Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// IAM Group id associated with the data mask rule
	IamGroupId *string `mandatory:"false" json:"iamGroupId"`

	TargetSelected TargetSelected `mandatory:"false" json:"targetSelected"`

	// Data Mask Categories
	DataMaskCategories []DataMaskCategoryEnum `mandatory:"false" json:"dataMaskCategories,omitempty"`

	// The status of the dataMaskRule.
	DataMaskRuleStatus DataMaskRuleStatusEnum `mandatory:"false" json:"dataMaskRuleStatus,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDataMaskRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDataMaskRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.DataMaskCategories {
		if _, ok := GetMappingDataMaskCategoryEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataMaskCategories: %s. Supported values are: %s.", val, strings.Join(GetDataMaskCategoryEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingDataMaskRuleStatusEnum(string(m.DataMaskRuleStatus)); !ok && m.DataMaskRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataMaskRuleStatus: %s. Supported values are: %s.", m.DataMaskRuleStatus, strings.Join(GetDataMaskRuleStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDataMaskRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		CompartmentId      *string                           `json:"compartmentId"`
		IamGroupId         *string                           `json:"iamGroupId"`
		TargetSelected     targetselected                    `json:"targetSelected"`
		DataMaskCategories []DataMaskCategoryEnum            `json:"dataMaskCategories"`
		DataMaskRuleStatus DataMaskRuleStatusEnum            `json:"dataMaskRuleStatus"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.DataMaskRuleStatus = model.DataMaskRuleStatus

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
