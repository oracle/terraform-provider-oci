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

// CreateCompliancePolicyRuleDetails The data to create a CompliancePolicyRule.
type CreateCompliancePolicyRuleDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	ProductVersion *ProductVersionDetails `mandatory:"true" json:"productVersion"`

	// PlatformConfiguration OCID for the patch type to which this CompliancePolicyRule applies.
	PatchType []string `mandatory:"true" json:"patchType"`

	PatchSelection PatchSelectionDetails `mandatory:"true" json:"patchSelection"`

	// The OCID of the compartment the CompliancePolicyRule belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique OCID of the CompliancePolicy this CompliancePolicyRule belongs to.
	CompliancePolicyId *string `mandatory:"false" json:"compliancePolicyId"`

	// Severity to which this CompliancePolicyRule applies.
	Severity []ComplianceRuleSeverityEnum `mandatory:"false" json:"severity,omitempty"`

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

func (m CreateCompliancePolicyRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCompliancePolicyRuleDetails) ValidateEnumValue() (bool, error) {
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
func (m *CreateCompliancePolicyRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompliancePolicyId *string                           `json:"compliancePolicyId"`
		Severity           []ComplianceRuleSeverityEnum      `json:"severity"`
		GracePeriod        *string                           `json:"gracePeriod"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		ProductVersion     *ProductVersionDetails            `json:"productVersion"`
		PatchType          []string                          `json:"patchType"`
		PatchSelection     patchselectiondetails             `json:"patchSelection"`
		CompartmentId      *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompliancePolicyId = model.CompliancePolicyId

	m.Severity = make([]ComplianceRuleSeverityEnum, len(model.Severity))
	copy(m.Severity, model.Severity)
	m.GracePeriod = model.GracePeriod

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.ProductVersion = model.ProductVersion

	m.PatchType = make([]string, len(model.PatchType))
	copy(m.PatchType, model.PatchType)
	nn, e = model.PatchSelection.UnmarshalPolymorphicJSON(model.PatchSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PatchSelection = nn.(PatchSelectionDetails)
	} else {
		m.PatchSelection = nil
	}

	m.CompartmentId = model.CompartmentId

	return
}
