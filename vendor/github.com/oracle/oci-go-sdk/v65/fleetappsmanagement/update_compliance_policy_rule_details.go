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

// UpdateCompliancePolicyRuleDetails The data to update a CompliancePolicyRule.
type UpdateCompliancePolicyRuleDetails struct {
	ProductVersion *ProductVersionDetails `mandatory:"false" json:"productVersion"`

	// PlatformConfiguration OCID for the patch type to which this CompliancePolicyRule applies.
	PatchType []string `mandatory:"false" json:"patchType"`

	// Severity to which this CompliancePolicyRule applies.
	Severity []ComplianceRuleSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	PatchSelection PatchSelectionDetails `mandatory:"false" json:"patchSelection"`

	// Grace period in days,weeks,months or years the exemption is applicable for the rule.
	// This enables a grace period when Fleet Application Management doesn't report the product as noncompliant when patch is not applied.
	GracePeriod *string `mandatory:"false" json:"gracePeriod"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCompliancePolicyRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCompliancePolicyRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Severity {
		if _, ok := GetMappingComplianceRuleSeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", val, strings.Join(GetComplianceRuleSeverityEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateCompliancePolicyRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ProductVersion *ProductVersionDetails            `json:"productVersion"`
		PatchType      []string                          `json:"patchType"`
		Severity       []ComplianceRuleSeverityEnum      `json:"severity"`
		PatchSelection patchselectiondetails             `json:"patchSelection"`
		GracePeriod    *string                           `json:"gracePeriod"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ProductVersion = model.ProductVersion

	m.PatchType = make([]string, len(model.PatchType))
	copy(m.PatchType, model.PatchType)
	m.Severity = make([]ComplianceRuleSeverityEnum, len(model.Severity))
	copy(m.Severity, model.Severity)
	nn, e = model.PatchSelection.UnmarshalPolymorphicJSON(model.PatchSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PatchSelection = nn.(PatchSelectionDetails)
	} else {
		m.PatchSelection = nil
	}

	m.GracePeriod = model.GracePeriod

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
