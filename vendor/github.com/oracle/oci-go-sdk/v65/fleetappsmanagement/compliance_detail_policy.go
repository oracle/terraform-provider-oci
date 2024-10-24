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

// ComplianceDetailPolicy Details of the Policy associated
type ComplianceDetailPolicy struct {

	// Compliance Policy Id
	CompliancePolicyId *string `mandatory:"true" json:"compliancePolicyId"`

	// Compliane Policy DisplayName
	CompliancePolicyDisplayName *string `mandatory:"false" json:"compliancePolicyDisplayName"`

	// Compliane Policy Rule Id
	CompliancePolicyRuleId *string `mandatory:"false" json:"compliancePolicyRuleId"`

	// Product Name
	CompliancePolicyRuleDisplayName *string `mandatory:"false" json:"compliancePolicyRuleDisplayName"`

	// Grace period in days,weeks,months or years the exemption is applicable for the rule.
	GracePeriod *string `mandatory:"false" json:"gracePeriod"`

	PatchSelection PatchSelectionDetails `mandatory:"false" json:"patchSelection"`
}

func (m ComplianceDetailPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceDetailPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ComplianceDetailPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompliancePolicyDisplayName     *string               `json:"compliancePolicyDisplayName"`
		CompliancePolicyRuleId          *string               `json:"compliancePolicyRuleId"`
		CompliancePolicyRuleDisplayName *string               `json:"compliancePolicyRuleDisplayName"`
		GracePeriod                     *string               `json:"gracePeriod"`
		PatchSelection                  patchselectiondetails `json:"patchSelection"`
		CompliancePolicyId              *string               `json:"compliancePolicyId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompliancePolicyDisplayName = model.CompliancePolicyDisplayName

	m.CompliancePolicyRuleId = model.CompliancePolicyRuleId

	m.CompliancePolicyRuleDisplayName = model.CompliancePolicyRuleDisplayName

	m.GracePeriod = model.GracePeriod

	nn, e = model.PatchSelection.UnmarshalPolymorphicJSON(model.PatchSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PatchSelection = nn.(PatchSelectionDetails)
	} else {
		m.PatchSelection = nil
	}

	m.CompliancePolicyId = model.CompliancePolicyId

	return
}
