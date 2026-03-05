// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateUnifiedAuditPolicyDetails The details required to create a new unified audit policy.
type UpdateUnifiedAuditPolicyDetails struct {

	// The display name of the unified audit policy in Data Safe. The name is modifiable and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the unified audit policy in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether the policy has been enabled or disabled.
	Status UnifiedAuditPolicyStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Lists the audit policy provisioning conditions.
	Conditions []PolicyCondition `mandatory:"false" json:"conditions"`

	// Indicates whether the unified audit policy overrides the enabled conditions on the target database during deployment.
	// When set to YES, This overrides both the enabled conditions and the status (enabled/disabled) on the target during security policy deployment.
	// If the same unified audit policy definition is referenced by another unified audit policy under a different security policy,
	// this value also takes precedence and overrides its enabled conditions if both are deployed on the same target database.
	OverrideTargetEnabledConditions UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum `mandatory:"false" json:"overrideTargetEnabledConditions,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateUnifiedAuditPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateUnifiedAuditPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUnifiedAuditPolicyStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUnifiedAuditPolicyStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum(string(m.OverrideTargetEnabledConditions)); !ok && m.OverrideTargetEnabledConditions != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OverrideTargetEnabledConditions: %s. Supported values are: %s.", m.OverrideTargetEnabledConditions, strings.Join(GetUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateUnifiedAuditPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                     *string                                                            `json:"displayName"`
		Description                     *string                                                            `json:"description"`
		Status                          UnifiedAuditPolicyStatusEnum                                       `json:"status"`
		Conditions                      []policycondition                                                  `json:"conditions"`
		OverrideTargetEnabledConditions UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum `json:"overrideTargetEnabledConditions"`
		FreeformTags                    map[string]string                                                  `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}                                  `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

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
	m.OverrideTargetEnabledConditions = model.OverrideTargetEnabledConditions

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum Enum with underlying type: string
type UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum string

// Set of constants representing the allowable values for UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum
const (
	UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsYes UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum = "YES"
	UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsNo  UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum = "NO"
)

var mappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum = map[string]UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum{
	"YES": UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsYes,
	"NO":  UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsNo,
}

var mappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumLowerCase = map[string]UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum{
	"yes": UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsYes,
	"no":  UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsNo,
}

// GetUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumValues Enumerates the set of values for UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum
func GetUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumValues() []UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum {
	values := make([]UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum, 0)
	for _, v := range mappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumStringValues Enumerates the set of values in String for UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum
func GetUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum(val string) (UpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnum, bool) {
	enum, ok := mappingUpdateUnifiedAuditPolicyDetailsOverrideTargetEnabledConditionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
