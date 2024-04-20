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

// UpdateDataMaskRuleDetails The information to be updated for a data mask rule.
type UpdateDataMaskRuleDetails struct {

	// Data mask rule display name
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment OCID where the resource is created
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// IAM group ID associated with the data mask rule
	IamGroupId *string `mandatory:"false" json:"iamGroupId"`

	TargetSelected TargetSelected `mandatory:"false" json:"targetSelected"`

	// Data mask rule categories
	DataMaskCategories []DataMaskCategoryEnum `mandatory:"false" json:"dataMaskCategories,omitempty"`

	// The status of the data mask rule
	DataMaskRuleStatus DataMaskRuleStatusEnum `mandatory:"false" json:"dataMaskRuleStatus,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
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
	copy(m.DataMaskCategories, model.DataMaskCategories)
	m.DataMaskRuleStatus = model.DataMaskRuleStatus

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
