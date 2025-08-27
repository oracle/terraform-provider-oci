// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateUnifiedAuditPolicyDetails The details required to create a new unified audit policy.
type CreateUnifiedAuditPolicyDetails struct {

	// The OCID of the security policy corresponding to the unified audit policy.
	SecurityPolicyId *string `mandatory:"true" json:"securityPolicyId"`

	// The OCID of the associated unified audit policy definition.
	UnifiedAuditPolicyDefinitionId *string `mandatory:"true" json:"unifiedAuditPolicyDefinitionId"`

	// The OCID of the compartment in which to create the unified audit policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Indicates whether the unified audit policy has been enabled or disabled.
	Status UnifiedAuditPolicyStatusEnum `mandatory:"true" json:"status"`

	// Lists the audit policy provisioning conditions.
	Conditions []PolicyCondition `mandatory:"true" json:"conditions"`

	// The display name of the unified audit policy in Data Safe. The name is modifiable and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the unified audit policy in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateUnifiedAuditPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUnifiedAuditPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAuditPolicyStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUnifiedAuditPolicyStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateUnifiedAuditPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                    *string                           `json:"displayName"`
		Description                    *string                           `json:"description"`
		FreeformTags                   map[string]string                 `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{} `json:"definedTags"`
		SecurityPolicyId               *string                           `json:"securityPolicyId"`
		UnifiedAuditPolicyDefinitionId *string                           `json:"unifiedAuditPolicyDefinitionId"`
		CompartmentId                  *string                           `json:"compartmentId"`
		Status                         UnifiedAuditPolicyStatusEnum      `json:"status"`
		Conditions                     []policycondition                 `json:"conditions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SecurityPolicyId = model.SecurityPolicyId

	m.UnifiedAuditPolicyDefinitionId = model.UnifiedAuditPolicyDefinitionId

	m.CompartmentId = model.CompartmentId

	m.Status = model.Status

	m.Conditions = make([]PolicyCondition, len(model.Conditions))
	for i, n := range model.Conditions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Conditions[i] = nn.(PolicyCondition)
		} else {
			m.Conditions[i] = nil
		}
	}
	return
}
