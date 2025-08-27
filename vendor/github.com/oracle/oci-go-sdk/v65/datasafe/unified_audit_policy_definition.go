// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAuditPolicyDefinition Resource represents a single unified audit policy definition.
type UnifiedAuditPolicyDefinition struct {

	// The OCID of the unified audit policy definition.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the unified audit policy definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the unified audit policy definition.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the unified audit policy definition.
	LifecycleState UnifiedAuditPolicyDefinitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the unified audit policy was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The description of the unified audit policy definition.
	Description *string `mandatory:"false" json:"description"`

	// Details about the current state of the unified audit policy definition.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The unified audit policy name in the target database.
	PolicyName *string `mandatory:"false" json:"policyName"`

	// Signifies whether the unified audit policy definition is seeded or not.
	IsSeeded *bool `mandatory:"false" json:"isSeeded"`

	// The category to which the unified audit policy belongs in the target database.
	AuditPolicyCategory UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum `mandatory:"false" json:"auditPolicyCategory,omitempty"`

	// The last date and time the unified audit policy was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The definition of the unified audit policy to be provisioned in the target database.
	PolicyDefinitionStatement *string `mandatory:"false" json:"policyDefinitionStatement"`

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

func (m UnifiedAuditPolicyDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAuditPolicyDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAuditPolicyDefinitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUnifiedAuditPolicyDefinitionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum(string(m.AuditPolicyCategory)); !ok && m.AuditPolicyCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditPolicyCategory: %s. Supported values are: %s.", m.AuditPolicyCategory, strings.Join(GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum Enum with underlying type: string
type UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum
const (
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryBasicActivity       UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "BASIC_ACTIVITY"
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryAdminUserActivity   UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "ADMIN_USER_ACTIVITY"
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryUserActivity        UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "USER_ACTIVITY"
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryOraclePredefined    UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "ORACLE_PREDEFINED"
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryComplianceStandard  UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "COMPLIANCE_STANDARD"
	UnifiedAuditPolicyDefinitionAuditPolicyCategorySqlFirewallAuditing UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "SQL_FIREWALL_AUDITING"
	UnifiedAuditPolicyDefinitionAuditPolicyCategoryCustom              UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = "CUSTOM"
)

var mappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum = map[string]UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum{
	"BASIC_ACTIVITY":        UnifiedAuditPolicyDefinitionAuditPolicyCategoryBasicActivity,
	"ADMIN_USER_ACTIVITY":   UnifiedAuditPolicyDefinitionAuditPolicyCategoryAdminUserActivity,
	"USER_ACTIVITY":         UnifiedAuditPolicyDefinitionAuditPolicyCategoryUserActivity,
	"ORACLE_PREDEFINED":     UnifiedAuditPolicyDefinitionAuditPolicyCategoryOraclePredefined,
	"COMPLIANCE_STANDARD":   UnifiedAuditPolicyDefinitionAuditPolicyCategoryComplianceStandard,
	"SQL_FIREWALL_AUDITING": UnifiedAuditPolicyDefinitionAuditPolicyCategorySqlFirewallAuditing,
	"CUSTOM":                UnifiedAuditPolicyDefinitionAuditPolicyCategoryCustom,
}

var mappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumLowerCase = map[string]UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum{
	"basic_activity":        UnifiedAuditPolicyDefinitionAuditPolicyCategoryBasicActivity,
	"admin_user_activity":   UnifiedAuditPolicyDefinitionAuditPolicyCategoryAdminUserActivity,
	"user_activity":         UnifiedAuditPolicyDefinitionAuditPolicyCategoryUserActivity,
	"oracle_predefined":     UnifiedAuditPolicyDefinitionAuditPolicyCategoryOraclePredefined,
	"compliance_standard":   UnifiedAuditPolicyDefinitionAuditPolicyCategoryComplianceStandard,
	"sql_firewall_auditing": UnifiedAuditPolicyDefinitionAuditPolicyCategorySqlFirewallAuditing,
	"custom":                UnifiedAuditPolicyDefinitionAuditPolicyCategoryCustom,
}

// GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumValues Enumerates the set of values for UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum
func GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumValues() []UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum {
	values := make([]UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum
func GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumStringValues() []string {
	return []string{
		"BASIC_ACTIVITY",
		"ADMIN_USER_ACTIVITY",
		"USER_ACTIVITY",
		"ORACLE_PREDEFINED",
		"COMPLIANCE_STANDARD",
		"SQL_FIREWALL_AUDITING",
		"CUSTOM",
	}
}

// GetMappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum(val string) (UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
