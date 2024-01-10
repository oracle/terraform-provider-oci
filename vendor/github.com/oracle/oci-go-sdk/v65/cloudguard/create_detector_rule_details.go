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

// CreateDetectorRuleDetails Details of a Detector Rule to be created in Detector Recipe
type CreateDetectorRuleDetails struct {

	// Name of the detector rule
	Name *string `mandatory:"true" json:"name"`

	// Id of source detector rule
	SourceDetectorRuleId *string `mandatory:"false" json:"sourceDetectorRuleId"`

	// Description of the detector rule
	Description *string `mandatory:"false" json:"description"`

	// Identifies state for detector rule
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The Risk Level
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// Configuration details
	Configurations []DetectorConfiguration `mandatory:"false" json:"configurations"`

	Condition Condition `mandatory:"false" json:"condition"`

	// user defined labels for a detector rule
	Labels []string `mandatory:"false" json:"labels"`

	// Recommendations of the detector rule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// ocid of the data source which needs to attached
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data Source entities mapping for a Detector Rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`
}

func (m CreateDetectorRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDetectorRuleDetails) ValidateEnumValue() (bool, error) {
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
func (m *CreateDetectorRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SourceDetectorRuleId *string                 `json:"sourceDetectorRuleId"`
		Description          *string                 `json:"description"`
		IsEnabled            *bool                   `json:"isEnabled"`
		RiskLevel            RiskLevelEnum           `json:"riskLevel"`
		Configurations       []DetectorConfiguration `json:"configurations"`
		Condition            condition               `json:"condition"`
		Labels               []string                `json:"labels"`
		Recommendation       *string                 `json:"recommendation"`
		DataSourceId         *string                 `json:"dataSourceId"`
		EntitiesMappings     []EntitiesMapping       `json:"entitiesMappings"`
		Name                 *string                 `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SourceDetectorRuleId = model.SourceDetectorRuleId

	m.Description = model.Description

	m.IsEnabled = model.IsEnabled

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
	m.Recommendation = model.Recommendation

	m.DataSourceId = model.DataSourceId

	m.EntitiesMappings = make([]EntitiesMapping, len(model.EntitiesMappings))
	copy(m.EntitiesMappings, model.EntitiesMappings)
	m.Name = model.Name

	return
}
