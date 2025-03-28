// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDetectorRuleDetails Parameters to be updated for a detector rule within a detector recipe.
type UpdateDetectorRuleDetails struct {

	// Enablement status of the detector rule
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The risk level of the detector rule
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// List of detector rule configurations
	Configurations []DetectorConfiguration `mandatory:"false" json:"configurations"`

	Condition Condition `mandatory:"false" json:"condition"`

	// User-defined labels for a detector rule
	Labels []string `mandatory:"false" json:"labels"`

	// Description for the detector rule
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for the detector rule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// The unique identifier of the attached data source
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data source entities mapping for a detector rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`
}

func (m UpdateDetectorRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDetectorRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetRiskLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDetectorRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RiskLevel        RiskLevelEnum           `json:"riskLevel"`
		Configurations   []DetectorConfiguration `json:"configurations"`
		Condition        condition               `json:"condition"`
		Labels           []string                `json:"labels"`
		Description      *string                 `json:"description"`
		Recommendation   *string                 `json:"recommendation"`
		DataSourceId     *string                 `json:"dataSourceId"`
		EntitiesMappings []EntitiesMapping       `json:"entitiesMappings"`
		IsEnabled        *bool                   `json:"isEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.RiskLevel = model.RiskLevel

	m.Configurations = make([]DetectorConfiguration, len(model.Configurations))
	copy(m.Configurations, model.Configurations)
	nn, e = model.Condition.UnmarshalPolymorphicJSON(model.Condition.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Condition = nn.(Condition)
	} else {
		m.Condition = nil
	}

	m.Labels = make([]string, len(model.Labels))
	copy(m.Labels, model.Labels)
	m.Description = model.Description

	m.Recommendation = model.Recommendation

	m.DataSourceId = model.DataSourceId

	m.EntitiesMappings = make([]EntitiesMapping, len(model.EntitiesMappings))
	copy(m.EntitiesMappings, model.EntitiesMappings)
	m.IsEnabled = model.IsEnabled

	return
}
