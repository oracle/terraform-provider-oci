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

// UnifiedAuditPolicySummary Resource represents a single unified audit policy on the target database.
type UnifiedAuditPolicySummary struct {

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

	// Signifies whether the unified audit policy is seeded or not.
	IsSeeded *bool `mandatory:"false" json:"isSeeded"`

	// Indicates whether the policy has been enabled or disabled.
	Status UnifiedAuditPolicyStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Indicates the users for which the unified audit policy is enabled.
	EnabledEntities UnifiedAuditPolicyEnabledEntitiesEnum `mandatory:"false" json:"enabledEntities,omitempty"`

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

func (m UnifiedAuditPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAuditPolicySummary) ValidateEnumValue() (bool, error) {
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
