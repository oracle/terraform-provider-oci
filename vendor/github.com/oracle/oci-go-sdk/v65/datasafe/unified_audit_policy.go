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

// UnifiedAuditPolicy Resource represents a single unified audit policy on the target database.
type UnifiedAuditPolicy struct {

	// The OCID of the unified audit policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the unified audit policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the unified audit policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the unified audit policy.
	LifecycleState UnifiedAuditPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the the unified audit policy was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The description of the unified audit policy.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the security policy corresponding to the unified audit policy.
	SecurityPolicyId *string `mandatory:"false" json:"securityPolicyId"`

	// The OCID of the associated unified audit policy definition.
	UnifiedAuditPolicyDefinitionId *string `mandatory:"false" json:"unifiedAuditPolicyDefinitionId"`

	// The details of the current state of the unified audit policy in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates whether the unified audit policy is seeded or not.
	IsSeeded *bool `mandatory:"false" json:"isSeeded"`

	// Indicates whether the policy has been enabled or disabled.
	Status UnifiedAuditPolicyStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Indicates on whom the audit policy is enabled.
	EnabledEntities UnifiedAuditPolicyEnabledEntitiesEnum `mandatory:"false" json:"enabledEntities,omitempty"`

	// Lists the audit policy provisioning conditions.
	Conditions []PolicyCondition `mandatory:"false" json:"conditions"`

	// The last date and time the unified audit policy was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m UnifiedAuditPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAuditPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAuditPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUnifiedAuditPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUnifiedAuditPolicyStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUnifiedAuditPolicyStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUnifiedAuditPolicyEnabledEntitiesEnum(string(m.EnabledEntities)); !ok && m.EnabledEntities != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnabledEntities: %s. Supported values are: %s.", m.EnabledEntities, strings.Join(GetUnifiedAuditPolicyEnabledEntitiesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAuditPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                    *string                               `json:"description"`
		SecurityPolicyId               *string                               `json:"securityPolicyId"`
		UnifiedAuditPolicyDefinitionId *string                               `json:"unifiedAuditPolicyDefinitionId"`
		LifecycleDetails               *string                               `json:"lifecycleDetails"`
		IsSeeded                       *bool                                 `json:"isSeeded"`
		Status                         UnifiedAuditPolicyStatusEnum          `json:"status"`
		EnabledEntities                UnifiedAuditPolicyEnabledEntitiesEnum `json:"enabledEntities"`
		Conditions                     []policycondition                     `json:"conditions"`
		TimeUpdated                    *common.SDKTime                       `json:"timeUpdated"`
		FreeformTags                   map[string]string                     `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags                     map[string]map[string]interface{}     `json:"systemTags"`
		Id                             *string                               `json:"id"`
		CompartmentId                  *string                               `json:"compartmentId"`
		DisplayName                    *string                               `json:"displayName"`
		LifecycleState                 UnifiedAuditPolicyLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated                    *common.SDKTime                       `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.SecurityPolicyId = model.SecurityPolicyId

	m.UnifiedAuditPolicyDefinitionId = model.UnifiedAuditPolicyDefinitionId

	m.LifecycleDetails = model.LifecycleDetails

	m.IsSeeded = model.IsSeeded

	m.Status = model.Status

	m.EnabledEntities = model.EnabledEntities

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
	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	return
}

// UnifiedAuditPolicyStatusEnum Enum with underlying type: string
type UnifiedAuditPolicyStatusEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyStatusEnum
const (
	UnifiedAuditPolicyStatusEnabled  UnifiedAuditPolicyStatusEnum = "ENABLED"
	UnifiedAuditPolicyStatusDisabled UnifiedAuditPolicyStatusEnum = "DISABLED"
)

var mappingUnifiedAuditPolicyStatusEnum = map[string]UnifiedAuditPolicyStatusEnum{
	"ENABLED":  UnifiedAuditPolicyStatusEnabled,
	"DISABLED": UnifiedAuditPolicyStatusDisabled,
}

var mappingUnifiedAuditPolicyStatusEnumLowerCase = map[string]UnifiedAuditPolicyStatusEnum{
	"enabled":  UnifiedAuditPolicyStatusEnabled,
	"disabled": UnifiedAuditPolicyStatusDisabled,
}

// GetUnifiedAuditPolicyStatusEnumValues Enumerates the set of values for UnifiedAuditPolicyStatusEnum
func GetUnifiedAuditPolicyStatusEnumValues() []UnifiedAuditPolicyStatusEnum {
	values := make([]UnifiedAuditPolicyStatusEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyStatusEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyStatusEnum
func GetUnifiedAuditPolicyStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUnifiedAuditPolicyStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyStatusEnum(val string) (UnifiedAuditPolicyStatusEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UnifiedAuditPolicyEnabledEntitiesEnum Enum with underlying type: string
type UnifiedAuditPolicyEnabledEntitiesEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyEnabledEntitiesEnum
const (
	UnifiedAuditPolicyEnabledEntitiesAllUsers          UnifiedAuditPolicyEnabledEntitiesEnum = "ALL_USERS"
	UnifiedAuditPolicyEnabledEntitiesIncludeUsers      UnifiedAuditPolicyEnabledEntitiesEnum = "INCLUDE_USERS"
	UnifiedAuditPolicyEnabledEntitiesIncludeRoles      UnifiedAuditPolicyEnabledEntitiesEnum = "INCLUDE_ROLES"
	UnifiedAuditPolicyEnabledEntitiesExcludeUsers      UnifiedAuditPolicyEnabledEntitiesEnum = "EXCLUDE_USERS"
	UnifiedAuditPolicyEnabledEntitiesIncludeUsersRoles UnifiedAuditPolicyEnabledEntitiesEnum = "INCLUDE_USERS_ROLES"
	UnifiedAuditPolicyEnabledEntitiesDisabled          UnifiedAuditPolicyEnabledEntitiesEnum = "DISABLED"
)

var mappingUnifiedAuditPolicyEnabledEntitiesEnum = map[string]UnifiedAuditPolicyEnabledEntitiesEnum{
	"ALL_USERS":           UnifiedAuditPolicyEnabledEntitiesAllUsers,
	"INCLUDE_USERS":       UnifiedAuditPolicyEnabledEntitiesIncludeUsers,
	"INCLUDE_ROLES":       UnifiedAuditPolicyEnabledEntitiesIncludeRoles,
	"EXCLUDE_USERS":       UnifiedAuditPolicyEnabledEntitiesExcludeUsers,
	"INCLUDE_USERS_ROLES": UnifiedAuditPolicyEnabledEntitiesIncludeUsersRoles,
	"DISABLED":            UnifiedAuditPolicyEnabledEntitiesDisabled,
}

var mappingUnifiedAuditPolicyEnabledEntitiesEnumLowerCase = map[string]UnifiedAuditPolicyEnabledEntitiesEnum{
	"all_users":           UnifiedAuditPolicyEnabledEntitiesAllUsers,
	"include_users":       UnifiedAuditPolicyEnabledEntitiesIncludeUsers,
	"include_roles":       UnifiedAuditPolicyEnabledEntitiesIncludeRoles,
	"exclude_users":       UnifiedAuditPolicyEnabledEntitiesExcludeUsers,
	"include_users_roles": UnifiedAuditPolicyEnabledEntitiesIncludeUsersRoles,
	"disabled":            UnifiedAuditPolicyEnabledEntitiesDisabled,
}

// GetUnifiedAuditPolicyEnabledEntitiesEnumValues Enumerates the set of values for UnifiedAuditPolicyEnabledEntitiesEnum
func GetUnifiedAuditPolicyEnabledEntitiesEnumValues() []UnifiedAuditPolicyEnabledEntitiesEnum {
	values := make([]UnifiedAuditPolicyEnabledEntitiesEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyEnabledEntitiesEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyEnabledEntitiesEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyEnabledEntitiesEnum
func GetUnifiedAuditPolicyEnabledEntitiesEnumStringValues() []string {
	return []string{
		"ALL_USERS",
		"INCLUDE_USERS",
		"INCLUDE_ROLES",
		"EXCLUDE_USERS",
		"INCLUDE_USERS_ROLES",
		"DISABLED",
	}
}

// GetMappingUnifiedAuditPolicyEnabledEntitiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyEnabledEntitiesEnum(val string) (UnifiedAuditPolicyEnabledEntitiesEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyEnabledEntitiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
